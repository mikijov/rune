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
	"errors"
	"reflect"
	"regexp"
)

// Externals interface allows user program to specify functions with which rune
// programs can interface with outside system. Examples of external functions
// are print, console, file etc. functions.
type Externals interface {
	GetDeclCount() int
	GetDecl(i int) (name string, value Object)

	DeclareUserFunction(name string, fn interface{}) error
}

type external struct {
	name  string
	value Object
}

type externals struct {
	decls []*external
}

// NewExternals creates new Externals instance.
func NewExternals() Externals {
	return &externals{
		decls: make([]*external, 0, 10),
	}
}

func (this *externals) GetDeclCount() int {
	return len(this.decls)
}

func (this *externals) GetDecl(i int) (name string, value Object) {
	decl := this.decls[i]
	return decl.name, decl.value
}

func (this *externals) DeclareUserFunction(name string, fn interface{}) error {
	matched, err := regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9_]*$", name)
	if err != nil {
		return err
	} else if !matched {
		return errors.New("invalid name; must match ^[a-zA-Z_][a-zA-Z0-9_]*$")
	}

	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		return errors.New("not a function")
	}

	typ := val.Type()

	paramTypes := make([]Type, 0, typ.NumIn())
	for i := 0; i < typ.NumIn(); i++ {
		paramType, err := go2runeType(typ.In(i))
		if err != nil {
			return err
		}
		paramTypes = append(paramTypes, paramType)
	}

	var returnType Type
	if typ.NumOut() == 0 {
		returnType = NewSimpleType(VOID)
	} else if typ.NumOut() > 1 {
		return errors.New("function can have no or one return value")
	} else {
		returnType, err = go2runeType(typ.Out(0))
		if err != nil {
			return err
		}
	}

	functionType := NewFunctionType(paramTypes, returnType)
	functionValue := NewUserFunction(functionType, val)

	this.decls = append(this.decls, &external{
		name:  name,
		value: functionValue,
	})

	return nil
}
