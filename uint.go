package valid

import (
	"flag"
	"fmt"
	"strconv"
)

// UintValue provides uint Value for flag pakage which validatable and
// constrainable.
type UintValue struct {
	value
	pv *uint
}

// Uint creates a validatable uint variable for flag.
func Uint(p *uint, val uint) *UintValue {
	if p == nil {
		p = new(uint)
	}
	*p = val
	i := &UintValue{pv: p}
	return i
}

// Set sets a value by string representation.
func (i *UintValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i.pv = uint(v)
	i.f = true
	return err
}

func (i *UintValue) get() uint {
	if i.pv == nil {
		return 0
	}
	return *i.pv
}

// Get returns value of the flag.
func (i *UintValue) Get() interface{} { return i.get() }

// String returns string representation for value of the flag.
func (i *UintValue) String() string {
	return strconv.FormatUint(uint64(i.get()), 10)
}

// Validate validates value of the flag.
func (i *UintValue) Validate(f *flag.Flag) error { return i.v.Validate(f) }

// MustSet declares "set at least once" constraint.
func (i *UintValue) MustSet() *UintValue {
	i.mustSet()
	return i
}

// Min declares lower limit constraint.
func (i *UintValue) Min(min uint) *UintValue {
	i.v.add(func() error {
		if n := i.get(); n < min {
			return fmt.Errorf("less than %d: %d", min, n)
		}
		return nil
	})
	return i
}

// Max declares uppper limit constraint.
func (i *UintValue) Max(max uint) *UintValue {
	i.v.add(func() error {
		if n := i.get(); n > max {
			return fmt.Errorf("greater than %d: %d", max, n)
		}
		return nil
	})
	return i
}

// OneOf declares "one of" constraint.
func (i *UintValue) OneOf(values ...uint) *UintValue {
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

var _ Validatable = &UintValue{}
