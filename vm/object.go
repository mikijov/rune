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
	"fmt"
	"math"
	"reflect"
)

// Object interface encapsulates all values in rune.
type Object interface {
	Type() Type
	String() string
	Value() reflect.Value
	Equal(other Object) bool
}

// ObjectFromValue creates new instance of Object of proper type according to
// the parameters passed to it.
func ObjectFromValue(typ Type, val reflect.Value) Object {
	switch typ.GetKind() {
	case INTEGER:
		return &integer{value: VmInteger(val.Int())}
	case REAL:
		return &real{value: VmReal(val.Float())}
	case STRING:
		return &strng{value: val.String()}
	case BOOLEAN:
		return &boolean{value: val.Bool()}
	case VOID:
		return &void{}
	case FUNCTION:
		return val.Interface().(Function)
	default:
		panic("invalid type")
	}
}

type void struct {
}

// NewVoid creates new Object of type void.
func NewVoid() Object                   { return &void{} }
func (this *void) Type() Type           { return NewSimpleType(VOID) }
func (this *void) String() string       { return string(VOID) }
func (this *void) Value() reflect.Value { return reflect.ValueOf(nil) }

func (this *void) Equal(other Object) bool {
	_, ok := other.(*void)
	return ok
}

// Integer is specialization of Object for integer values.
type Integer interface {
	Object
	GetValue() VmInteger
}

type integer struct {
	value VmInteger
}

// NewInteger creates Integer object with specified value.
func NewInteger(val VmInteger) Integer     { return &integer{val} }
func (this *integer) Type() Type           { return NewSimpleType(INTEGER) }
func (this *integer) String() string       { return fmt.Sprintf("%d", this.value) }
func (this *integer) Value() reflect.Value { return reflect.ValueOf(int64(this.value)) }
func (this *integer) GetValue() VmInteger  { return this.value }

func (this *integer) Equal(other Object) bool {
	o, ok := other.(*integer)
	if !ok {
		return false
	}
	return this.value == o.value
}

// Real is specialization of Object for real values.
type Real interface {
	Object
	GetValue() VmReal
}

type real struct {
	value VmReal
}

// NewReal creates Real object with specified value.
func NewReal(val VmReal) Real           { return &real{val} }
func (this *real) Type() Type           { return NewSimpleType(REAL) }
func (this *real) String() string       { return fmt.Sprintf("%f", this.value) }
func (this *real) Value() reflect.Value { return reflect.ValueOf(float64(this.value)) }
func (this *real) GetValue() VmReal     { return this.value }

func (this *real) Equal(other Object) bool {
	o, ok := other.(*real)
	if !ok {
		return false
	}
	// compares floats while ignoring last bit
	// fmt.Printf("%f: %d . %d\n", this.value, math.Float64bits(float64(this.value)), math.Float64bits(float64(o.value)))
	return math.Nextafter(float64(this.value), float64(o.value)) == float64(o.value)
}

// Boolean is specialization of Object for boolean values.
type Boolean interface {
	Object
	GetValue() bool
}

type boolean struct {
	value bool
}

// NewBoolean creates Boolean object with specified value.
func NewBoolean(val bool) Boolean          { return &boolean{val} }
func (this *boolean) Type() Type           { return NewSimpleType(BOOLEAN) }
func (this *boolean) String() string       { return fmt.Sprintf("%t", this.value) }
func (this *boolean) Value() reflect.Value { return reflect.ValueOf(this.value) }
func (this *boolean) GetValue() bool       { return this.value }

func (this *boolean) Equal(other Object) bool {
	o, ok := other.(*boolean)
	if !ok {
		return false
	}
	return this.value == o.value
}

// String is specialization of Object for string values.
type String interface {
	Object
	GetValue() string
}

type strng struct {
	value string
}

// NewString creates String object with specified value.
func NewString(val string) String        { return &strng{val} }
func (this *strng) Type() Type           { return NewSimpleType(STRING) }
func (this *strng) String() string       { return this.value }
func (this *strng) Value() reflect.Value { return reflect.ValueOf(this.value) }
func (this *strng) GetValue() string     { return this.value }

func (this *strng) Equal(other Object) bool {
	o, ok := other.(*strng)
	if !ok {
		return false
	}
	return this.value == o.value
}

// Function is specialization of Object for function values.
type Function interface {
	Object
	Call(env Environment, params []Expression) Object
	// Equal(other Function) bool
}

type zeroFunction struct {
	typ Type
}

// NewZeroFunction creates Function object of specific type but no implementation.
func NewZeroFunction(typ Type) Function {
	return &zeroFunction{
		typ: typ,
	}
}

func (this *zeroFunction) Type() Type           { return this.typ }
func (this *zeroFunction) String() string       { return "nil" }
func (this *zeroFunction) Value() reflect.Value { return reflect.ValueOf(nil) }

func (this *zeroFunction) Call(env Environment, params []Expression) Object {
	env.SetError("call to nil function")
	return nil
}

func (this *zeroFunction) Equal(other Object) bool {
	// TODO: return simply DeepEqual; no need to cast
	o, ok := other.(*zeroFunction)
	if !ok {
		return false
	} else {
		return reflect.DeepEqual(this, o)
	}
}

type function struct {
	typ        Type
	paramNames []string
	body       Statement
}

// NewFunction creates Function object of specific type and implementation.
func NewFunction(typ Type, paramNames []string, body Statement) Function {
	return &function{
		typ:        typ,
		paramNames: paramNames,
		body:       body,
	}
}

func (this *function) Type() Type           { return this.typ }
func (this *function) String() string       { return this.typ.String() }
func (this *function) Value() reflect.Value { return reflect.ValueOf(this) }

func (this *function) Call(env Environment, params []Expression) Object {
	funcEnv := NewFunctionEnvironment(env, this.typ.GetResultType())

	for i, param := range params {
		val := param.Execute(env)
		funcEnv.Declare(this.paramNames[i], val)
	}

	this.body.Execute(funcEnv)

	failed, msg := funcEnv.GetError()
	if failed {
		env.SetError(msg)
	}

	return funcEnv.GetReturnValue()
}

func (this *function) Equal(other Object) bool {
	o, ok := other.(*function)
	if !ok {
		return false
	} else {
		return reflect.DeepEqual(this, o)
	}
}

type userFunction struct {
	typ   Type
	value reflect.Value
}

// NewUserFunction creates Function object of specific type and external
// implementation.
func NewUserFunction(typ Type, value reflect.Value) Function {
	return &userFunction{
		typ:   typ,
		value: value,
	}
}

func (this *userFunction) Type() Type           { return this.typ }
func (this *userFunction) String() string       { return "user function" }
func (this *userFunction) Value() reflect.Value { return this.value }

func (this *userFunction) Call(env Environment, params []Expression) Object {
	valueParams := make([]reflect.Value, 0, len(params))
	for _, param := range params {
		val := param.Execute(env)
		valueParams = append(valueParams, val.Value())
	}

	returnValues := this.value.Call(valueParams)
	returnType := this.typ.GetResultType()
	if returnType.GetKind() == VOID {
		return returnType.GetZero()
	} else {
		return ObjectFromValue(returnType, returnValues[0])
	}
}

func (this *userFunction) Equal(other Object) bool {
	// TODO: return simply DeepEqual; no need to cast
	o, ok := other.(*userFunction)
	if !ok {
		return false
	} else {
		return reflect.DeepEqual(this.value, o.value)
	}
}
