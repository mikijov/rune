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

package parser

import (
	"github.com/mikijov/rune/vm"
)

// Scope interface represents a scope which in turn determines lifetime and
// visibility of variables.
type Scope interface {
	Declare(name string, typ vm.Type) bool
	Get(name string) vm.Type
}

type scope struct {
	store map[string]vm.Type
	outer Scope
}

// NewScope creates new scope in which variable can be declared. If outer scope
// is specified, any variable lookup will recursively search outer scopes if
// the variable is not found in the current scope.
func NewScope(outer Scope) Scope {
	return &scope{
		store: make(map[string]vm.Type),
		outer: outer,
	}
}

func (this *scope) Declare(name string, typ vm.Type) bool {
	if _, ok := this.store[name]; ok {
		return false // already declared in the current scope
	}
	this.store[name] = typ
	return true
}

func (this *scope) Get(name string) vm.Type {
	retVal := this.store[name]
	if retVal != nil {
		return retVal
	}
	if this.outer != nil {
		return this.outer.Get(name)
	}
	return nil
}
