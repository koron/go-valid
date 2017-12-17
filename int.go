package valid

import (
	"flag"
	"fmt"
	"strconv"
)

// IntValue provides int Value for flag pakage which validatable and
// constrainable.
type IntValue struct {
	value
	pv *int
}

// Int creates a validatable int variable for flag.
func Int(p *int, val int) *IntValue {
	if p == nil {
		p = new(int)
	}
	*p = val
	return &IntValue{pv: p}
}

// Set sets a value by string representation.
func (i *IntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	*i.pv = int(v)
	i.f = true
	return err
}

func (i *IntValue) get() int {
	if i.pv == nil {
		return 0
	}
	return *i.pv
}

// Get returns value of the flag.
func (i *IntValue) Get() interface{} { return i.get() }

// String returns string representation for value of the flag.
func (i *IntValue) String() string {
	return strconv.FormatInt(int64(i.get()), 10)
}

// Validate validates value of the flag.
func (i *IntValue) Validate(f *flag.Flag) error { return i.v.Validate(f) }

// MustSet declares "set at least once" constraint.
func (i *IntValue) MustSet() *IntValue {
	i.mustSet()
	return i
}

// Min declares lower limit constraint.
func (i *IntValue) Min(min int) *IntValue {
	i.v.add(func() error {
		if n := i.get(); n < min {
			return fmt.Errorf("less than %d: %d", min, n)
		}
		return nil
	})
	return i
}

// Max declares uppper limit constraint.
func (i *IntValue) Max(max int) *IntValue {
	i.v.add(func() error {
		if n := i.get(); n > max {
			return fmt.Errorf("greater than %d: %d", max, n)
		}
		return nil
	})
	return i
}

// OneOf declares "one of" constraint.
func (i *IntValue) OneOf(values ...int) *IntValue {
	i.v.add(func() error {
		n := i.get()
		for _, v := range values {
			if n == v {
				return nil
			}
		}
		return fmt.Errorf("not one of %v: %d", values, n)
	})
	return i
}

var _ Validatable = &IntValue{}
