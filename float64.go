package valid

import (
	"flag"
	"fmt"
	"strconv"
)

// Float64Value provides float64 Value for flag pakage which validatable and
// constrainable.
type Float64Value struct {
	value
	pv *float64
}

// Float64 creates a validatable float64 variable for flag.
func Float64(p *float64, val float64) *Float64Value {
	if p == nil {
		p = new(float64)
	}
	*p = val
	return &Float64Value{pv: p}
}

// Set sets a value by string representation.
func (fv *Float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*fv.pv = v
	fv.f = true
	return err
}

func (fv *Float64Value) get() float64 {
	if fv.pv == nil {
		return 0
	}
	return *fv.pv
}

// Get returns value of the flag.
func (fv *Float64Value) Get() interface{} { return fv.get() }

// String returns string representation for value of the flag.
func (fv *Float64Value) String() string {
	return strconv.FormatFloat(fv.get(), 'g', -1, 64)
}

// Validate validates value of the flag.
func (fv *Float64Value) Validate(f *flag.Flag) error { return fv.v.Validate(f) }

// MustSet declares "set at least once" constraint.
func (fv *Float64Value) MustSet() *Float64Value {
	fv.mustSet()
	return fv
}

// Min declares lower limit constraint.
func (fv *Float64Value) Min(min float64) *Float64Value {
	fv.v.add(func() error {
		if n := fv.get(); n < min {
			return fmt.Errorf("less than %e: %e", min, n)
		}
		return nil
	})
	return fv
}

// Max declares uppper limit constraint.
func (fv *Float64Value) Max(max float64) *Float64Value {
	fv.v.add(func() error {
		if n := fv.get(); n > max {
			return fmt.Errorf("greater than %e: %e", max, n)
		}
		return nil
	})
	return fv
}

// OneOf declares "one of" constraint.
func (fv *Float64Value) OneOf(values ...float64) *Float64Value {
	fv.v.add(func() error {
		n := fv.get()
		for _, v := range values {
			if n == v {
				return nil
			}
		}
		return fmt.Errorf("not one of %v: %e", values, n)
	})
	return fv
}

var _ Validatable = &Float64Value{}
