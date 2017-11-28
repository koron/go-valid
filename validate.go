package valid

import (
	"flag"
	"fmt"
)

type Validatable interface {
	flag.Getter
	Validate(*flag.Flag) error
}

type ValidationErrors struct {
	Errors []error
}

func (err *ValidationErrors) Error() string {
	return fmt.Sprintf("validation errors: [%+v]", err.Errors)
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
