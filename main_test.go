package main

import (
	"testing"

	"mikijov/rune/vm"
)

func TestDeclarations(t *testing.T) {
	tests := []struct {
		code string
		typ  vm.Type
	}{
		{"a :int;", vm.NewSimpleType(vm.INTEGER)},
		{"a :real;", vm.NewSimpleType(vm.REAL)},
		{"a :bool;", vm.NewSimpleType(vm.BOOLEAN)},
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: failed to compile", i)
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		if len(env.GetLocalVariables()) != 1 {
			t.Errorf("test %d: unexpected variables %v", i, env.GetLocalVariables())
		}
		a := env.Get("a")
		if !test.typ.Equal(a.Type()) {
			t.Errorf("test %d: type mismatch %v != %v", i, a.Type(), test.typ)
			continue
		}
		if !test.typ.GetZero().Equal(a) {
			t.Errorf("test %d: unexpected initial value %v", i, a)
		}
	}

}
