package valid

import (
	"flag"
	"reflect"
	"testing"
)

func testParse(t *testing.T, fs *flag.FlagSet, act interface{}, ok bool, exp interface{}, args ...string) bool {
	// reset *act by zero.
	el := reflect.ValueOf(act).Elem()
	el.Set(reflect.Zero(el.Type()))
	err := Parse(fs, args)
	// when parse/validation failed
	if !ok {
		if err == nil {
			t.Errorf("should be failed: %v", args)
			return false
		}
		return true
	}
	// when parse/validation succeeded
	if err != nil {
		t.Errorf("should be succeeded: %v %v: %s", exp, args, err)
		return false
	}
	v := el.Interface()
	if !reflect.DeepEqual(v, exp) {
		t.Errorf("the option is not match: act=%v exp=%v %v", v, exp, args)
		return false
	}
	return true
}
