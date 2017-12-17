package valid

import (
	"flag"
	"fmt"
	"strconv"
)

// Int64Value provides int64 Value for flag pakage which validatable and
// constrainable.
type Int64Value struct {
	value
	pv *int64
}

// Int64 creates a validatable int64 variable for flag.
func Int64(p *int64, val int64) *Int64Value {
	if p == nil {
		p = new(int64)
	}
	*p = val
	i := &Int64Value{pv: p}
	return i
}

// Set sets a value by string representation.
func (i *Int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i.pv = int64(v)
	i.f = true
	return err
}

func (i *Int64Value) get() int64 {
	if i.pv == nil {
		return 0
	}
	return *i.pv
}

// Get returns value of the flag.
func (i *Int64Value) Get() interface{} { return i.get() }

// String returns string representation for value of the flag.
func (i *Int64Value) String() string {
	return strconv.FormatInt(i.get(), 10)
}

// Validate validates value of the flag.
func (i *Int64Value) Validate(f *flag.Flag) error { return i.v.Validate(f) }

// MustSet declares "set at least once" constraint.
func (i *Int64Value) MustSet() *Int64Value {
	i.mustSet()
	return i
}

// Min declares lower limit constraint.
func (i *Int64Value) Min(min int64) *Int64Value {
	i.v.add(func() error {
		if n := i.get(); n < min {
			return fmt.Errorf("less than %d: %d", min, n)
		}
		return nil
	})
	return i
}

// Max declares uppper limit constraint.
func (i *Int64Value) Max(max int64) *Int64Value {
	i.v.add(func() error {
		if n := i.get(); n > max {
			return fmt.Errorf("greater than %d: %d", max, n)
		}
		return nil
	})
	return i
}

// OneOf declares "one of" constraint.
func (i *Int64Value) OneOf(values ...int64) *Int64Value {
	i.v.add(func() error {
		n := i.get()
		for _, v := range values {
			if n == v {
				return nil
			}
		}
		return fmt.Errorf("not one of %s", values)
	})
	return i
}

var _ Validatable = &Int64Value{}
