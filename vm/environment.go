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
	"fmt"
	"reflect"
)

// Environment interface represents the current state of the running program,
// its variables and functions.
type Environment interface {
	Import(ext Externals)

	Declare(name string, value Object)
	Set(name string, value Object)
	Get(name string) Object

	SetReturnValue(value Object)
	GetReturnValue() Object
	SetReturning()
	IsReturning() bool
	SetError(msg string)
	GetError() (failed bool, msg string)

	GetLocalVariables() []string
}

type functionContext struct {
	retVal    Object
	returning bool
	errorFlag bool
	errorMsg  string
}

type environment struct {
	outer           Environment
	store           map[string]Object
	functionContext *functionContext
}

// NewEnvironment creates new environment tied to a scope. These environments
// are used for any block of code surrounded by {} like in loops, if statements
// etc. This excludes function scope. This means that each nested scope can have
// its own variables, but shares the return value with parent scopes all the way
// to the function scope.
func NewEnvironment(outer Environment) Environment {
	var fc *functionContext
	if outer != nil {
		fc = outer.(*environment).functionContext
	} else {
		fc = &functionContext{
			retVal: NewSimpleType(VOID).GetZero(),
		}
	}

	return &environment{
		outer:           outer,
		store:           make(map[string]Object),
		functionContext: fc,
	}
}

// NewFunctionEnvironment creates new environment tied to a function scope. It
// is very similar to environments created by NewEnvironment() with the exception
// that NewFunctionEnvironment creates it's own return variable.
func NewFunctionEnvironment(outer Environment, returnType Type) Environment {
	return &environment{
		outer: outer,
		store: make(map[string]Object),
		functionContext: &functionContext{
			retVal: returnType.GetZero(),
		},
	}
}

func go2runeType(typ reflect.Type) (Type, error) {
	switch typ.Kind() {
	case reflect.Int64:
		return NewSimpleType(INTEGER), nil
	case reflect.Float64:
		return NewSimpleType(REAL), nil
	case reflect.Bool:
		return NewSimpleType(BOOLEAN), nil
	case reflect.String:
		return NewSimpleType(STRING), nil
	// TODO: implement support for passing Environment to user function
	// case reflect.Interface:
	// 	return ?
	default:
		return NewSimpleType(VOID), errors.New("unsupported type")
	}
}

func (this *environment) Import(ext Externals) {
	for i := 0; i < ext.GetDeclCount(); i++ {
		name, value := ext.GetDecl(i)
		this.Declare(name, value)
	}
}

func (this *environment) Declare(name string, value Object) {
	_, ok := this.store[name]
	if ok {
		panic(fmt.Sprintf("variable redeclared: %s", name))
	}
	this.store[name] = value
}

func (this *environment) Set(name string, value Object) {
	obj, ok := this.store[name]
	if ok {
		// TODO: this check should have been done statically by the compiler
		if !obj.Type().Equal(value.Type()) {
			panic(fmt.Sprintf("type mismatch assigning: %s. %v != %v", name, obj.Type(), value.Type()))
		}
		this.store[name] = value
	} else {
		if this.outer != nil {
			this.outer.Set(name, value)
		} else {
			panic(fmt.Sprintf("undeclared variable: %s", name))
		}
	}
}

func (this *environment) Get(name string) Object {
	obj, ok := this.store[name]
	if !ok && this.outer != nil {
		return this.outer.Get(name)
	}
	return obj
}

func (this *environment) SetReturnValue(value Object) {
	// TODO: this check should have been done statically by the compiler
	if !this.functionContext.retVal.Type().Equal(value.Type()) {
		panic(fmt.Sprintf("type mismatch returning: %v != %v", this.functionContext.retVal.Type(), value.Type()))
	}
	this.functionContext.retVal = value
}

func (this *environment) GetReturnValue() Object {
	return this.functionContext.retVal
}

func (this *environment) SetReturning() {
	this.functionContext.returning = true
}

func (this *environment) IsReturning() bool {
	return this.functionContext.returning || this.functionContext.errorFlag
}

func (this *environment) SetError(msg string) {
	this.functionContext.errorFlag = true
	this.functionContext.errorMsg = msg
}

func (this *environment) GetError() (failed bool, msg string) {
	return this.functionContext.errorFlag, this.functionContext.errorMsg
}

func (this *environment) GetLocalVariables() []string {
	retVal := make([]string, len(this.store))

	i := 0
	for k := range this.store {
		retVal[i] = k
		i++
	}

	return retVal
}
