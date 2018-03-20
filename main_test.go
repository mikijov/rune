// Copyright © 2018 Milutin Jovanović jovanovic.milutin@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"

	"github.com/mikijov/rune/vm"
)

func TestDeclarations(t *testing.T) {
	tests := []struct {
		code string
		typ  vm.Type
	}{
		{"a :int", vm.NewSimpleType(vm.INTEGER)},
		{"a :real", vm.NewSimpleType(vm.REAL)},
		{"a :bool", vm.NewSimpleType(vm.BOOLEAN)},
		// TODO: test functions
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
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

func TestDeclarationsByValue(t *testing.T) {
	tests := []struct {
		code  string
		value vm.Object
	}{
		{"a := 1;", vm.NewInteger(1)},
		{"a := 1.0;", vm.NewReal(1.0)},
		{"a := true;", vm.NewBoolean(true)},
		{"a := false;", vm.NewBoolean(false)},
		// TODO: test functions
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		if len(env.GetLocalVariables()) != 1 {
			t.Errorf("test %d: unexpected variables %v", i, env.GetLocalVariables())
		}
		a := env.Get("a")
		if !test.value.Type().Equal(a.Type()) {
			t.Errorf("test %d: type mismatch %v != %v", i, a.Type(), test.value.Type())
			continue
		}
		if !test.value.Equal(a) {
			t.Errorf("test %d: unexpected value %v", i, a)
		}
	}
}

func TestDeclarationsAndInitialization(t *testing.T) {
	tests := []struct {
		code  string
		value vm.Object
	}{
		{"a :int = 1;", vm.NewInteger(1)},
		{"a :real = 1.0;", vm.NewReal(1.0)},
		{"a :bool = true;", vm.NewBoolean(true)},
		{"a :bool = false;", vm.NewBoolean(false)},
		// TODO: test functions
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		if len(env.GetLocalVariables()) != 1 {
			t.Errorf("test %d: unexpected variables %v", i, env.GetLocalVariables())
		}
		a := env.Get("a")
		if !test.value.Type().Equal(a.Type()) {
			t.Errorf("test %d: type mismatch %v != %v", i, a.Type(), test.value.Type())
			continue
		}
		if !test.value.Equal(a) {
			t.Errorf("test %d: unexpected value %v", i, a)
		}
	}
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		code  string
		value vm.Object
	}{
		{"a := 1 + 2 * 3;", vm.NewInteger(7)},
		{"a := 11 + 9 % 4;", vm.NewInteger(12)},
		{"a := (1 + 2) * 3;", vm.NewInteger(9)},
		{"a := 3 * 2 + 1;", vm.NewInteger(7)},
		{"a := 3 * (2 + 1);", vm.NewInteger(9)},
		{"a := 1 - 9 / 3;", vm.NewInteger(-2)},
		{"a := 9 / 3 - 1;", vm.NewInteger(2)},

		{"a := 1 < 2 and 2 < 3;", vm.NewBoolean(true)},
		{"a := 1 > 2 and 2 < 3;", vm.NewBoolean(false)},
		{"a := 1 > 2 or 2 < 3;", vm.NewBoolean(true)},

		// left to right of identical precedence
		{"a := 4 * 3 / 2;", vm.NewInteger(6)},
		{"a := 4 / 2 * 3;", vm.NewInteger(6)},
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		a := env.Get("a")
		if !test.value.Equal(a) {
			t.Errorf("test %d: %s unexpected value %v != %v", i, test.code, a, test.value)
		}
	}
}

func TestOperators(t *testing.T) {
	tests := []struct {
		code  string
		value vm.Object
	}{
		{"a := 1; a = a + 2;", vm.NewInteger(3)},
		{"a := 1; a = a - 3;", vm.NewInteger(-2)},
		// {"a := 1; a = a + -3;", vm.NewInteger(-2)},
		// {"a := 1; a = a + +3;", vm.NewInteger(4)},
		{"a := 2; a = a * 3;", vm.NewInteger(6)},
		{"a := 9; a = a / 2;", vm.NewInteger(4)},
		{"a := 9; a = a % 5;", vm.NewInteger(4)},
		{"a := 5; a = a & 6;", vm.NewInteger(4)},
		{"a := 6; a = a | 3;", vm.NewInteger(7)},
		{"a := 6; a = a ^ 3;", vm.NewInteger(5)},

		{"a := 1; a += 2;", vm.NewInteger(3)},
		{"a := 1; a -= 3;", vm.NewInteger(-2)},
		// {"a := 1; a += -3;", vm.NewInteger(-2)},
		// {"a := 1; a += +3;", vm.NewInteger(4)},
		{"a := 2; a *= 3;", vm.NewInteger(6)},
		{"a := 9; a /= 2;", vm.NewInteger(4)},
		{"a := 9; a %= 5;", vm.NewInteger(4)},
		{"a := 5; a &= 6;", vm.NewInteger(4)},
		{"a := 6; a |= 3;", vm.NewInteger(7)},
		{"a := 6; a ^= 3;", vm.NewInteger(5)},

		{"a := 1.1; a = a + 2.2;", vm.NewReal(3.3)},
		{"a := 1.1; a = a - 3.3;", vm.NewReal(-2.2)},
		// {"a := 1.1; a = a + -3.3;", vm.NewReal(-2.0)},
		// {"a := 1.1; a = a + +3.3;", vm.NewReal(4.0)},
		{"a := 2.1; a = a * 3.3;", vm.NewReal(6.93)},
		{"a := 7.5; a = a / 3.0;", vm.NewReal(2.5)},

		{"a := 1.1; a += 2.2;", vm.NewReal(3.3)},
		{"a := 1.1; a -= 3.3;", vm.NewReal(-2.2)},
		// {"a := 1.1; a += -3.3;", vm.NewReal(-2)},
		// {"a := 1.1; a += +3.3;", vm.NewReal(4)},
		{"a := 2.1; a *= 3.3;", vm.NewReal(6.93)},
		{"a := 7.5; a /= 3.0;", vm.NewReal(2.5)},

		{"a := 1 == 1;", vm.NewBoolean(true)},
		{"a := 1 != 1;", vm.NewBoolean(false)},
		{"a := 1 > 1;", vm.NewBoolean(false)},
		{"a := 1 < 1;", vm.NewBoolean(false)},
		{"a := 1 <= 1;", vm.NewBoolean(true)},
		{"a := 1 >= 1;", vm.NewBoolean(true)},

		{"a := 1 == 2;", vm.NewBoolean(false)},
		{"a := 1 != 2;", vm.NewBoolean(true)},
		{"a := 1 > 2;", vm.NewBoolean(false)},
		{"a := 1 < 2;", vm.NewBoolean(true)},
		{"a := 1 <= 2;", vm.NewBoolean(true)},
		{"a := 1 >= 2;", vm.NewBoolean(false)},

		{"a := 2 == 1;", vm.NewBoolean(false)},
		{"a := 2 != 1;", vm.NewBoolean(true)},
		{"a := 2 > 1;", vm.NewBoolean(true)},
		{"a := 2 < 1;", vm.NewBoolean(false)},
		{"a := 2 <= 1;", vm.NewBoolean(false)},
		{"a := 2 >= 1;", vm.NewBoolean(true)},

		{"a := 1.0 == 1.0;", vm.NewBoolean(true)},
		{"a := 1.0 != 1.0;", vm.NewBoolean(false)},
		{"a := 1.0 > 1.0;", vm.NewBoolean(false)},
		{"a := 1.0 < 1.0;", vm.NewBoolean(false)},
		{"a := 1.0 <= 1.0;", vm.NewBoolean(true)},
		{"a := 1.0 >= 1.0;", vm.NewBoolean(true)},

		{"a := 1.0 == 2.0;", vm.NewBoolean(false)},
		{"a := 1.0 != 2.0;", vm.NewBoolean(true)},
		{"a := 1.0 > 2.0;", vm.NewBoolean(false)},
		{"a := 1.0 < 2.0;", vm.NewBoolean(true)},
		{"a := 1.0 <= 2.0;", vm.NewBoolean(true)},
		{"a := 1.0 >= 2.0;", vm.NewBoolean(false)},

		{"a := 2.0 == 1.0;", vm.NewBoolean(false)},
		{"a := 2.0 != 1.0;", vm.NewBoolean(true)},
		{"a := 2.0 > 1.0;", vm.NewBoolean(true)},
		{"a := 2.0 < 1.0;", vm.NewBoolean(false)},
		{"a := 2.0 <= 1.0;", vm.NewBoolean(false)},
		{"a := 2.0 >= 1.0;", vm.NewBoolean(true)},

		{"a := true and true;", vm.NewBoolean(true)},
		{"a := true and false;", vm.NewBoolean(false)},
		{"a := false and false;", vm.NewBoolean(false)},
		{"a := true or true;", vm.NewBoolean(true)},
		{"a := true or false;", vm.NewBoolean(true)},
		{"a := false or false;", vm.NewBoolean(false)},
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		a := env.Get("a")
		if !test.value.Equal(a) {
			t.Errorf("test %d: %s unexpected value %v != %v", i, test.code, a, test.value)
		}
	}
}

func TestConditionalStatements(t *testing.T) {
	tests := []struct {
		code  string
		value vm.Object
	}{
		{"a := 0; if true { a = 1; }", vm.NewInteger(1)},
		{"a := 0; if false { a = 1; }", vm.NewInteger(0)},
		{"a := 0; if true { a = 1; } else { a = 2; }", vm.NewInteger(1)},
		{"a := 0; if false { a = 1; } else { a = 2; }", vm.NewInteger(2)},
		{"a := 10; loop while a < 10 { a += 1; }", vm.NewInteger(10)},
		{"a := 0; loop while a < 10 { a += 1; }", vm.NewInteger(10)},
		// TODO: reimplement this once we have break statement
		{"func fn() :int { r := 0; loop { r += 1; if r > 10 { return r; } } } a := fn();", vm.NewInteger(11)},
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		a := env.Get("a")
		if !test.value.Equal(a) {
			t.Errorf("test %d: %s unexpected value %v != %v", i, test.code, a, test.value)
		}
	}
}

func TestFunctions(t *testing.T) {
	tests := []struct {
		code  string
		value vm.Object
	}{
		{"a := 0; func fn() { a = 1; } fn();", vm.NewInteger(1)},
		{"a := 0; func fn() :int { return 1; } a = fn();", vm.NewInteger(1)},
		{"a := 0; func fn(x :int) { a = x; } fn(1);", vm.NewInteger(1)},
		{"func fn(x :int) :int { return x + 1; } a := fn(1);", vm.NewInteger(2)},
		{"func fn(x :int, y :real) :real { if x != 0 { return y * 2.0; } else { return y * 3.0; } } a := fn(0, 1.1);", vm.NewReal(3.3)},
		{"func fn(x :int, y :real) :real { if x != 0 { return y * 2.0; } else { return y * 3.0; } } a := fn(1, 1.1);", vm.NewReal(2.2)},
		{"func fn(x, y :int, m, n :real) :real { if x == y { return m + n; } else { return m - n; } } a := fn(1, 2, 3.2, 2.1);", vm.NewReal(1.1)},
		{"func fn(x, y :int, m, n :real) :real { if x == y { return m + n; } else { return m - n; } } a := fn(3, 3, 3.2, 2.1);", vm.NewReal(5.3)},
	}

	for i, test := range tests {
		program, errors := Compile(test.code, nil)
		if len(errors.GetErrors()) > 0 {
			t.Errorf("test %d: %s failed to compile %v", i, test.code, errors.GetErrors())
			continue
		}

		env := vm.NewEnvironment(nil)
		program.Execute(env)
		a := env.Get("a")
		if !test.value.Equal(a) {
			t.Errorf("test %d: %s unexpected value %v != %v", i, test.code, a, test.value)
		}
	}
}

func testVariable(t *testing.T, env vm.Environment, name string, expected vm.Object) {
	val := env.Get(name)
	if !val.Equal(expected) {
		t.Errorf("invalid value for variable %s; expected:%s got:%s", name, expected, val)
	}
}

func testField(t *testing.T, s vm.Object, field string, expected vm.Object) {
	typ := s.Type()
	index := typ.GetFieldIndex(field)
	if index < 0 {
		t.Errorf("unknown field %s", field)
		return
	}
	struc := s.(vm.Struct)
	val := struc.Get(index)
	if !val.Equal(expected) {
		t.Errorf("invalid value for field %s; expected:%s got:%s", field, expected, val)
	}
}

func TestStructZero(t *testing.T) {
	code := `
s1 :struct {
	b1 :bool
	i1, i2 :int
	r1, r2 :real
}
`
	program, errors := Compile(code, nil)
	if len(errors.GetErrors()) > 0 {
		t.Fatalf("failed to compile %v", errors.GetErrors())
	}

	env := vm.NewEnvironment(nil)
	program.Execute(env)
	s1 := env.Get("s1")
	if s1.Type().GetKind() != vm.STRUCT {
		t.Fatalf("expected struct but got %s", s1.Type().GetKind())
	}
	typ := s1.Type()
	if typ.GetFieldCount() != 5 {
		t.Fatalf("invalid number of fields (%d)", typ.GetFieldCount())
	}
	testField(t, s1, "b1", vm.NewSimpleType(vm.BOOLEAN).GetZero())
	testField(t, s1, "i1", vm.NewSimpleType(vm.INTEGER).GetZero())
	testField(t, s1, "i2", vm.NewSimpleType(vm.INTEGER).GetZero())
	testField(t, s1, "r1", vm.NewSimpleType(vm.REAL).GetZero())
	testField(t, s1, "r2", vm.NewSimpleType(vm.REAL).GetZero())
}

func TestStructFields(t *testing.T) {
	code := `
s1 :struct {
	b1 :bool
	i1, i2 :int
	r1, r2 :real
}
s1.b1 = true;
s1.i1 = 1;
s1.i2 = 2;
s1.r1 = 1.1;
s1.r2 = 2.2;

b := s1.b1;
i := s1.i1;
r := s1.r1;
`
	program, errors := Compile(code, nil)
	if len(errors.GetErrors()) > 0 {
		t.Fatalf("failed to compile %v", errors.GetErrors())
	}

	env := vm.NewEnvironment(nil)
	program.Execute(env)
	s1 := env.Get("s1")
	if s1.Type().GetKind() != vm.STRUCT {
		t.Fatalf("expected struct but got %s", s1.Type().GetKind())
	}
	typ := s1.Type()
	if typ.GetFieldCount() != 5 {
		t.Fatalf("invalid number of fields (%d)", typ.GetFieldCount())
	}
	testField(t, s1, "b1", vm.NewBoolean(true))
	testField(t, s1, "i1", vm.NewInteger(1))
	testField(t, s1, "i2", vm.NewInteger(2))
	testField(t, s1, "r1", vm.NewReal(1.1))
	testField(t, s1, "r2", vm.NewReal(2.2))

	testVariable(t, env, "b", vm.NewBoolean(true))
	testVariable(t, env, "i", vm.NewInteger(1))
	testVariable(t, env, "r", vm.NewReal(1.1))
}

func TestTypeDeclarationZero(t *testing.T) {
	code := `
type t1 :struct {
	b1 :bool
	i1, i2 :int
	r1, r2 :real
}
s1 :t1
`
	program, errors := Compile(code, nil)
	if len(errors.GetErrors()) > 0 {
		t.Fatalf("failed to compile %v", errors.GetErrors())
	}

	env := vm.NewEnvironment(nil)
	program.Execute(env)
	s1 := env.Get("s1")
	if s1.Type().GetKind() != vm.STRUCT {
		t.Fatalf("expected struct but got %s", s1.Type().GetKind())
	}
	typ := s1.Type()
	if typ.GetFieldCount() != 5 {
		t.Fatalf("invalid number of fields (%d)", typ.GetFieldCount())
	}
	testField(t, s1, "b1", vm.NewSimpleType(vm.BOOLEAN).GetZero())
	testField(t, s1, "i1", vm.NewSimpleType(vm.INTEGER).GetZero())
	testField(t, s1, "i2", vm.NewSimpleType(vm.INTEGER).GetZero())
	testField(t, s1, "r1", vm.NewSimpleType(vm.REAL).GetZero())
	testField(t, s1, "r2", vm.NewSimpleType(vm.REAL).GetZero())
}

func TestTypeDeclarationStructFields(t *testing.T) {
	code := `
type t1 :struct {
	b1 :bool
	i1, i2 :int
	r1, r2 :real
}
s1 :t1
s1.b1 = true;
s1.i1 = 1;
s1.i2 = 2;
s1.r1 = 1.1;
s1.r2 = 2.2;

b := s1.b1;
i := s1.i1;
r := s1.r1;
`
	program, errors := Compile(code, nil)
	if len(errors.GetErrors()) > 0 {
		t.Fatalf("failed to compile %v", errors.GetErrors())
	}

	env := vm.NewEnvironment(nil)
	program.Execute(env)
	s1 := env.Get("s1")
	if s1.Type().GetKind() != vm.STRUCT {
		t.Fatalf("expected struct but got %s", s1.Type().GetKind())
	}
	typ := s1.Type()
	if typ.GetFieldCount() != 5 {
		t.Fatalf("invalid number of fields (%d)", typ.GetFieldCount())
	}
	testField(t, s1, "b1", vm.NewBoolean(true))
	testField(t, s1, "i1", vm.NewInteger(1))
	testField(t, s1, "i2", vm.NewInteger(2))
	testField(t, s1, "r1", vm.NewReal(1.1))
	testField(t, s1, "r2", vm.NewReal(2.2))

	testVariable(t, env, "b", vm.NewBoolean(true))
	testVariable(t, env, "i", vm.NewInteger(1))
	testVariable(t, env, "r", vm.NewReal(1.1))
}

func TestFibbonaci(t *testing.T) {
	code := `
func fibbonaci(count :int) :int {
	if count <= 0 {
		return 0;
	} else if count == 1 {
		return 1;
	} else {
		n1 := 1;
		n2 := 1;
		count = count - 2;
		loop while count > 0 {
			count = count - 1;
			temp := n2 + n1;
			n2 = n1;
			n1 = temp;
		}
		return n1;
	}
}
a := fibbonaci(50);
`
	expected := vm.NewInteger(12586269025)

	program, errors := Compile(code, nil)
	if len(errors.GetErrors()) > 0 {
		t.Fatalf("failed to compile %v", errors.GetErrors())
	}

	env := vm.NewEnvironment(nil)
	program.Execute(env)
	a := env.Get("a")
	if !expected.Equal(a) {
		t.Errorf("unexpected value %s != %s", a, expected)
	}
}
