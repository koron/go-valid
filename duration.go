package valid

import (
	"flag"
	"fmt"
	"time"
)

// DurationValue provides time.Duration Value for flag pakage which validatable
// and constrainable.
type DurationValue struct {
	value
	pv *time.Duration
}

func Duration(p *time.Duration, val time.Duration) *DurationValue {
	if p == nil {
		p = new(time.Duration)
	}
	*p = val
	return &DurationValue{pv: p}
}

func (dv *DurationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*dv.pv = v
	dv.f = true
	return err
}

func (dv *DurationValue) get() time.Duration {
	if dv.pv == nil {
		return 0
	}
	return *dv.pv
}

func (dv *DurationValue) Get() interface{} { return dv.get() }

func (dv *DurationValue) String() string {
	return dv.String()
}

func (dv *DurationValue) Validate(f *flag.Flag) error {
	return dv.v.Validate(f)
}

func (dv *DurationValue) MustSet() *DurationValue {
	dv.mustSet()
	return dv
}

func (dv *DurationValue) Min(min time.Duration) *DurationValue {
	dv.v.add(func() error {
		if n := dv.get(); n < min {
			return fmt.Errorf("less than %s: %s", min, n)
		}
		return nil
	})
	return dv
}

func (dv *DurationValue) Max(max time.Duration) *DurationValue {
	dv.v.add(func() error {
		if n := dv.get(); n > max {
			return fmt.Errorf("greater than %s: %s", max, n)
		}
		return nil
	})
	return dv
}

func (dv *DurationValue) OneOf(values []time.Duration) *DurationValue {
	dv.v.add(func() error {
		n := dv.get()
		for _, v := range values {
			if n == v {
				return nil
			}
		}
		return fmt.Errorf("not one of %v: %s", values, n)
	})
	return dv
}

var _ Validatable = &DurationValue{}
