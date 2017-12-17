package valid

import (
	"flag"
	"fmt"
	"strconv"
)

// Uint64Value provides uint64 Value for flag pakage which validatable and
// constrainable.
type Uint64Value struct {
	value
	pv *uint64
}

func Uint64(p *uint64, val uint64) *Uint64Value {
	if p == nil {
		p = new(uint64)
	}
	*p = val
	i := &Uint64Value{pv: p}
	return i
}

func (i *Uint64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i.pv = uint64(v)
	i.f = true
	return err
}

func (i *Uint64Value) get() uint64 {
	if i.pv == nil {
		return 0
	}
	return *i.pv
}

func (i *Uint64Value) Get() interface{} { return i.get() }

func (i *Uint64Value) String() string {
	return strconv.FormatUint(i.get(), 10)
}

func (i *Uint64Value) Validate(f *flag.Flag) error { return i.v.Validate(f) }

func (i *Uint64Value) MustSet() *Uint64Value {
	i.mustSet()
	return i
}

func (i *Uint64Value) Min(min uint64) *Uint64Value {
	i.v.add(func() error {
		if n := i.get(); n < min {
			return fmt.Errorf("less than %d: %d", min, n)
		}
		return nil
	})
	return i
}

func (i *Uint64Value) Max(max uint64) *Uint64Value {
	i.v.add(func() error {
		if n := i.get(); n > max {
			return fmt.Errorf("greater than %d: %d", max, n)
		}
		return nil
	})
	return i
}

func (i *Uint64Value) OneOf(values ...uint64) *Uint64Value {
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

var _ Validatable = &Uint64Value{}
