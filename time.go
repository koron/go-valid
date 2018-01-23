package valid

import (
	"flag"
	"fmt"
	"time"
)

// TimeValue provides time.Time Value for flag pakage which validatable
// and constrainable.
type TimeValue struct {
	value
	pv     *time.Time
	layout string
}

// Time creates a validatable time.Time variable for flag.
// Default layout is RFC3339, if you want change it use TimeValue.WithLayout()
func Time(p *time.Time, val time.Time) *TimeValue {
	if p == nil {
		p = new(time.Time)
	}
	*p = val
	return &TimeValue{
		pv:     p,
		layout: time.RFC3339,
	}
}

// WithLayout modifies layout string for time.Parse()
func (tv *TimeValue) WithLayout(v string) *TimeValue {
	tv.layout = v
	return tv
}

// Set sets a value by string representation.
func (tv *TimeValue) Set(s string) error {
	v, err := time.Parse(tv.layout, s)
	*tv.pv = v
	tv.f = true
	return err
}

func (tv *TimeValue) get() time.Time {
	if tv.pv == nil {
		return time.Time{}
	}
	return *tv.pv
}

// Get returns value of the flag.
func (tv *TimeValue) Get() interface{} { return tv.get() }

// String returns string representation for value of the flag.
func (tv *TimeValue) String() string {
	return tv.get().Format(tv.layout)
}

// Validate validates value of the flag.
func (tv *TimeValue) Validate(f *flag.Flag) error {
	return tv.v.Validate(f)
}

// MustSet declares "set at least once" constraint.
func (tv *TimeValue) MustSet() *TimeValue {
	tv.mustSet()
	return tv
}

// Min declares lower limit constraint.
func (tv *TimeValue) Min(min time.Time) *TimeValue {
	tv.v.add(func() error {
		if n := tv.get(); n.Before(min) {
			return fmt.Errorf("less than %s: %s", min, n)
		}
		return nil
	})
	return tv
}

// Max declares uppper limit constraint.
func (tv *TimeValue) Max(max time.Time) *TimeValue {
	tv.v.add(func() error {
		if n := tv.get(); n.After(max) {
			return fmt.Errorf("greater than %s: %s", max, n)
		}
		return nil
	})
	return tv
}

// OneOf declares "one of" constraint.
func (tv *TimeValue) OneOf(values ...time.Time) *TimeValue {
	tv.v.add(func() error {
		n := tv.get()
		for _, v := range values {
			if n.Equal(v) {
				return nil
			}
		}
		return fmt.Errorf("not one of %v: %s", values, n)
	})
	return tv
}

var _ Validatable = &TimeValue{}
