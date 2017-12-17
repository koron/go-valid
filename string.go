package valid

import (
	"flag"
	"fmt"
)

// StringValue provides string Value for flag pakage which validatable and
// constrainable.
type StringValue struct {
	value
	pv *string
}

// String creates a validatable string variable for flag.
func String(p *string, val string) *StringValue {
	if p == nil {
		p = new(string)
	}
	*p = val
	return &StringValue{pv: p}
}

// Set sets a value by string representation.
func (s *StringValue) Set(v string) error {
	*s.pv = v
	s.f = true
	return nil
}

func (s *StringValue) get() string {
	if s.pv == nil {
		return ""
	}
	return *s.pv
}

// Get returns value of the flag.
func (s *StringValue) Get() interface{} { return s.get() }

// String returns string representation for value of the flag.
func (s *StringValue) String() string { return s.get() }

// Validate validates value of the flag.
func (s *StringValue) Validate(f *flag.Flag) error { return s.v.Validate(f) }

// MustSet declares "set at least once" constraint.
func (s *StringValue) MustSet() *StringValue {
	s.mustSet()
	return s
}

// Min declares length lower limit constraint.
func (s *StringValue) Min(min int) *StringValue {
	s.v.add(func() error {
		if l := len(s.get()); l < min {
			return fmt.Errorf("shorter than %d: %d", min, l)
		}
		return nil
	})
	return s
}

// Max declares length uppper limit constraint.
func (s *StringValue) Max(max int) *StringValue {
	s.v.add(func() error {
		if l := len(s.get()); l > max {
			return fmt.Errorf("longer than %d: %d", max, l)
		}
		return nil
	})
	return s
}

// OneOf declares "one of" constraint.
func (s *StringValue) OneOf(values ...string) *StringValue {
	s.v.add(func() error {
		t := s.get()
		for _, v := range values {
			if t == v {
				return nil
			}
		}
		return fmt.Errorf("not one of %v: %q", values, t)
	})
	return s
}

var _ Validatable = &StringValue{}
