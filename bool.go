package valid

import (
	"flag"
	"strconv"
)

// BoolValue provides bool Value for flag pakage which validatable and
// constrainable.
type BoolValue struct {
	value
	pv *bool
}

// Bool creates a validatable bool variable for flag.
func Bool(p *bool, val bool) *BoolValue {
	if p == nil {
		p = new(bool)
	}
	*p = val
	return &BoolValue{pv: p}
}

// Set sets a value by string representation.
func (b *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b.pv = v
	b.f = true
	return err
}

func (b *BoolValue) get() bool {
	if b.pv == nil {
		return false
	}
	return *b.pv
}

// Get returns value of the flag.
func (b *BoolValue) Get() interface{} { return b.get() }

// String returns string representation for value of the flag.
func (b *BoolValue) String() string { return strconv.FormatBool(b.get()) }

// IsBoolFlag represents value is bool and omittable.
func (b *BoolValue) IsBoolFlag() bool { return true }

// Validate validates value of the flag.
func (b *BoolValue) Validate(f *flag.Flag) error { return b.v.Validate(f) }

// MustSet declares "set at least once" constraint.
func (b *BoolValue) MustSet() *BoolValue {
	b.mustSet()
	return b
}

var _ Validatable = &BoolValue{}
