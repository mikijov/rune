package vm

import (
	"fmt"
	"testing"
)

func TestEnvironment(t *testing.T) {
	fmt.Println("test1")

	i := &integer{value: 123}
	r := &real{value: 456.789}

	env := NewEnvironment(nil)

	if env.GetInteger("x") != nil {
		t.Error("variable should not be defined")
	}
	if env.GetReal("y") != nil {
		t.Error("variable should not be defined")
	}

	env.SetInteger("x", i)
	if env.GetInteger("x").GetValue() != i.GetValue() {
		t.Error("wrong value")
	}

	env.SetReal("y", r)
	if env.GetReal("y").GetValue() != r.GetValue() {
		t.Error("wrong value")
	}
}

func TestEnvironmentScope(t *testing.T) {
	fmt.Println("test2")

	i := &integer{value: 123}
	i2 := &integer{value: 321}
	r := &real{value: 456.789}

	outer := NewEnvironment(nil)

	// check it's empty
	if outer.GetInteger("x") != nil {
		t.Error("variable should not be defined")
	}
	if outer.GetReal("y") != nil {
		t.Error("variable should not be defined")
	}

	outer.SetInteger("x", i)
	outer.SetReal("y", r)

	// test basic get
	if outer.GetInteger("x").GetValue() != i.GetValue() {
		t.Error("wrong value")
	}
	if outer.GetReal("y").GetValue() != r.GetValue() {
		t.Error("wrong value")
	}

	inner := NewEnvironment(outer)

	// test that outer scope is visible from the inner one
	if inner.GetInteger("x").GetValue() != i.GetValue() {
		t.Error("wrong value")
	}
	if inner.GetReal("y").GetValue() != r.GetValue() {
		t.Error("wrong value")
	}

	// test that vars set in inner scope are not visible in the outer one
	inner.SetInteger("a", i)
	inner.SetReal("b", r)
	if inner.GetInteger("a").GetValue() != i.GetValue() {
		t.Error("wrong value")
	}
	if inner.GetReal("b").GetValue() != r.GetValue() {
		t.Error("wrong value")
	}
	if outer.GetInteger("a") != nil {
		t.Error("should not be visible")
	}
	if outer.GetReal("b") != nil {
		t.Error("should not be visible")
	}

	// test updating value in the outer scope form the inner one
	inner.SetInteger("x", i2)
	if outer.GetInteger("x").GetValue() != i2.GetValue() {
		t.Error("wrong value")
	}
}
