package vm

import (
	"reflect"
	"testing"
)

func valueEqual(l, r reflect.Value) bool {
	if l.Kind() != r.Kind() {
		return false
	} else if l.Kind() == reflect.Int64 {
		return l.Int() == r.Int()
	} else if l.Kind() == reflect.Float64 {
		return l.Float() == r.Float()
	} else if l.Kind() == reflect.Bool {
		return l.Bool() == r.Bool()
	} else {
		return false
	}
}

func TestObjectFromValue(t *testing.T) {
	tests := []struct {
		typ Type
		val reflect.Value
	}{
		{NewSimpleType(INTEGER), reflect.ValueOf(int64(0))},
		{NewSimpleType(INTEGER), reflect.ValueOf(int64(1))},
		{NewSimpleType(INTEGER), reflect.ValueOf(int64(-1))},
		{NewSimpleType(REAL), reflect.ValueOf(float64(0.0))},
		{NewSimpleType(REAL), reflect.ValueOf(float64(1.1))},
		{NewSimpleType(REAL), reflect.ValueOf(float64(-1.1))},
		{NewSimpleType(BOOLEAN), reflect.ValueOf(true)},
		{NewSimpleType(BOOLEAN), reflect.ValueOf(false)},
		// {NewSimpleType(VOID), reflect.ValueOf(nil)},
	}

	for i, test := range tests {
		actual := ObjectFromValue(test.typ, test.val)
		if !test.typ.Equal(actual.Type()) {
			t.Errorf("type mismatch %d: %v != %v", i, test.typ, actual.Type())
		}
		if !valueEqual(test.val, actual.Value()) {
			t.Errorf("value mismatch %d: %v != %v", i, test.val, actual.Value())
		}
	}
}
