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

func Uint(p *uint, val uint) *UintValue {
	if p == nil {
		p = new(uint)
	}
	*p = val
	i := &UintValue{pv: p}
	return i
}

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

func (i *UintValue) Get() interface{} { return i.get() }

func (i *UintValue) String() string {
	return strconv.FormatUint(uint64(i.get()), 10)
}

func (i *UintValue) Validate(f *flag.Flag) error { return i.v.Validate(f) }

func (i *UintValue) MustSet() *UintValue {
	i.mustSet()
	return i
}

func (i *UintValue) Min(min uint) *UintValue {
	i.v.add(func() error {
		if n := i.get(); n < min {
			return fmt.Errorf("less than %d: %d", min, n)
		}
		return nil
	})
	return i
}

func (i *UintValue) Max(max uint) *UintValue {
	i.v.add(func() error {
		if n := i.get(); n > max {
			return fmt.Errorf("greater than %d: %d", max, n)
		}
		return nil
	})
	return i
}

func (i *UintValue) OneOf(values []uint) *UintValue {
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
