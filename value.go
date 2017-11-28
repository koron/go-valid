package valid

import (
	"errors"
)

type value struct {
	f bool
	v validators
}

func (v *value) mustSet() {
	v.v.add(func() error {
		if !v.f {
			return errors.New("required but not set")
		}
		return nil
	})
}
