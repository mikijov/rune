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

package vm

import (
	"testing"
)

func TestEnvironment(t *testing.T) {
	i := &integer{value: 123}
	i1 := &integer{value: 321}
	r := &real{value: 456.789}
	r1 := &real{value: 654.987}

	env := NewEnvironment(nil)

	if env.Get("x") != nil {
		t.Error("variable should not be defined")
	}
	if env.Get("y") != nil {
		t.Error("variable should not be defined")
	}

	env.Declare("x", i)
	if env.Get("x").(Integer).GetValue() != i.GetValue() {
		t.Error("wrong value")
	}

	env.Declare("y", r)
	if env.Get("y").(Real).GetValue() != r.GetValue() {
		t.Error("wrong value")
	}

	env.Set("x", i1)
	if env.Get("x").(Integer).GetValue() != i1.GetValue() {
		t.Error("wrong value")
	}

	env.Set("y", r1)
	if env.Get("y").(Real).GetValue() != r1.GetValue() {
		t.Error("wrong value")
	}

	expected := map[string]bool{
		"x": false,
		"y": false,
	}
	for _, variable := range env.GetLocalVariables() {
		encoutered, ok := expected[variable]
		if !ok {
			t.Errorf("unexpected variable '%s'", variable)
			continue
		}
		if encoutered {
			t.Errorf("repeated variable '%s'", variable)
			continue
		}
		expected[variable] = true
	}
	for name, encoutered := range expected {
		if !encoutered {
			t.Errorf("expected but not declared '%s'", name)
		}
	}
}

func TestEnvironmentScope(t *testing.T) {
	i := &integer{value: 123}
	i2 := &integer{value: 321}
	r := &real{value: 456.789}

	outer := NewEnvironment(nil)
	inner := NewEnvironment(outer)

	// test that outer scope is visible from the inner one
	outer.Declare("x", i)
	outer.Declare("y", r)
	if outer.Get("x").(Integer).GetValue() != i.GetValue() {
		t.Error("wrong value")
	}
	if outer.Get("y").(Real).GetValue() != r.GetValue() {
		t.Error("wrong value")
	}
	if inner.Get("x").(Integer).GetValue() != i.GetValue() {
		t.Error("wrong value")
	}
	if inner.Get("y").(Real).GetValue() != r.GetValue() {
		t.Error("wrong value")
	}

	// test that vars set in inner scope are not visible in the outer one
	inner.Declare("a", i)
	inner.Declare("b", r)
	if inner.Get("a").(Integer).GetValue() != i.GetValue() {
		t.Error("wrong value")
	}
	if inner.Get("b").(Real).GetValue() != r.GetValue() {
		t.Error("wrong value")
	}
	if outer.Get("a") != nil {
		t.Error("should not be visible")
	}
	if outer.Get("b") != nil {
		t.Error("should not be visible")
	}

	// test updating value in the outer scope form the inner one
	inner.Set("x", i2)
	if outer.Get("x").(Integer).GetValue() != i2.GetValue() {
		t.Error("wrong value")
	}
	if inner.Get("x").(Integer).GetValue() != i2.GetValue() {
		t.Error("wrong value")
	}
}
