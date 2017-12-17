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

func Bool(p *bool, val bool) *BoolValue {
	if p == nil {
		p = new(bool)
	}
	*p = val
	return &BoolValue{pv: p}
}

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

func (b *BoolValue) Get() interface{} { return b.get() }

func (b *BoolValue) String() string { return strconv.FormatBool(b.get()) }

func (b *BoolValue) IsBoolFlag() bool { return true }

func (b *BoolValue) Validate(f *flag.Flag) error { return b.v.Validate(f) }

func (b *BoolValue) MustSet() *BoolValue {
	b.mustSet()
	return b
}

var _ Validatable = &BoolValue{}
