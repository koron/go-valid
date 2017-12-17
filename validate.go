package valid

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

type Validatable interface {
	flag.Getter
	Validate(*flag.Flag) error
}

// Validate validates all flags in FlagSet.  It will returns ValidationErrors if met some errors.
func Validate(fs *flag.FlagSet) error {
	errs := make([]error, 0, fs.NFlag())
	fs.VisitAll(func(f *flag.Flag) {
		v, ok := f.Value.(Validatable)
		if !ok {
			return
		}
		err := v.Validate(f)
		if err == nil {
			return
		}
		errs = append(errs, err)
	})
	if len(errs) == 0 {
		return nil
	}
	return &ValidationErrors{Errors: errs}
}

// Parse parses flag definitions from the argument list and validates
// constraint.
func Parse(fs *flag.FlagSet, args []string) error {
	err := fs.Parse(args)
	if err != nil {
		return err
	}
	err = Validate(fs)
	if err != nil {
		switch errorHandling(fs) {
		case flag.ContinueOnError:
			return err
		case flag.ExitOnError:
			os.Exit(2)
		case flag.PanicOnError:
			panic(err)
		}
	}
	return nil
}

type validator func() error

type validators []validator

func (vs *validators) add(v validator) {
	*vs = append(*vs, v)
}

func (vs validators) Validate(f *flag.Flag) error {
	for _, v := range vs {
		err := v()
		if err != nil {
			return fmt.Errorf("option -%s: %s", f.Name, err.Error())
		}
	}
	return nil
}

func errorHandling(fs *flag.FlagSet) flag.ErrorHandling {
	v := reflect.ValueOf(fs).Elem().FieldByName("errorHandling")
	switch v.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return flag.ErrorHandling(v.Int())
	default:
		return 0
	}
}
