package validator

import (
	"reflect"
	"testing"
)

func TestValidator(t *testing.T) {
	v := New()
	g := reflect.Indirect(reflect.ValueOf(v)).Type().Name()
	e := reflect.ValueOf(Validator{}).Type().Name()
	if g != e {
		t.Errorf("expected struct %s, got %s", e, g)
	}
}
