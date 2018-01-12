package vm

import (
	"fmt"
	"reflect"
)

type Object interface {
	Type() Type
	Inspect() string
	Value() reflect.Value
}

func ObjectFromValue(typ Type, val reflect.Value) Object {
	switch typ.GetKind() {
	case INTEGER:
		return &integer{value: VmInteger(val.Int())}
	case REAL:
		return &real{value: VmReal(val.Float())}
	case STRING:
		return &string_{value: val.String()}
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

func (this *void) Type() Type           { return NewSimpleType(VOID) }
func (this *void) Inspect() string      { return string(VOID) }
func (this *void) Value() reflect.Value { return reflect.ValueOf(VOID) }

type Integer interface {
	Object
	GetValue() VmInteger
	SetValue(value VmInteger)
}

type integer struct {
	value VmInteger
}

func (this *integer) Type() Type               { return NewSimpleType(INTEGER) }
func (this *integer) Inspect() string          { return fmt.Sprintf("%d", this.value) }
func (this *integer) Value() reflect.Value     { return reflect.ValueOf(int64(this.value)) }
func (this *integer) GetValue() VmInteger      { return this.value }
func (this *integer) SetValue(value VmInteger) { this.value = value }

type Real interface {
	Object
	GetValue() VmReal
	SetValue(value VmReal)
}

type real struct {
	value VmReal
}

func (this *real) Type() Type            { return NewSimpleType(REAL) }
func (this *real) Inspect() string       { return fmt.Sprintf("%f", this.value) }
func (this *real) Value() reflect.Value  { return reflect.ValueOf(float64(this.value)) }
func (this *real) GetValue() VmReal      { return this.value }
func (this *real) SetValue(value VmReal) { this.value = value }

type Boolean interface {
	Object
	GetValue() bool
	SetValue(value bool)
}

type boolean struct {
	value bool
}

func (this *boolean) Type() Type           { return NewSimpleType(BOOLEAN) }
func (this *boolean) Inspect() string      { return fmt.Sprintf("%t", this.value) }
func (this *boolean) Value() reflect.Value { return reflect.ValueOf(this.value) }
func (this *boolean) GetValue() bool       { return this.value }
func (this *boolean) SetValue(value bool)  { this.value = value }

type String interface {
	Object
	GetValue() string
	SetValue(value string)
}

type string_ struct {
	value string
}

func (this *string_) Type() Type            { return NewSimpleType(STRING) }
func (this *string_) Inspect() string       { return this.value }
func (this *string_) Value() reflect.Value  { return reflect.ValueOf(this.value) }
func (this *string_) GetValue() string      { return this.value }
func (this *string_) SetValue(value string) { this.value = value }

type Function interface {
	Object
	Call(env Environment, params []Expression) Object
	Equal(other Function) bool
}

type zeroFunction struct {
	typ Type
}

func NewZeroFunction(typ Type) Function {
	return &zeroFunction{
		typ: typ,
	}
}

func (this *zeroFunction) Type() Type           { return this.typ }
func (this *zeroFunction) Inspect() string      { return "nil" }
func (this *zeroFunction) Value() reflect.Value { return reflect.ValueOf(nil) }

func (this *zeroFunction) Call(env Environment, params []Expression) Object {
	env.SetError("call to nil function")
	return nil
}

func (this *zeroFunction) Equal(other Function) bool {
	return this.typ.Equal(other.Type())
}

type function struct {
	typ        Type
	paramNames []string
	body       Statement
}

func NewFunction(typ Type, paramNames []string, body Statement) Function {
	return &function{
		typ:        typ,
		paramNames: paramNames,
		body:       body,
	}
}

func (this *function) Type() Type           { return this.typ }
func (this *function) Inspect() string      { return this.typ.String() }
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

func (this *function) Equal(other Function) bool {
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

func NewUserFunction(typ Type, value reflect.Value) Function {
	return &userFunction{
		typ:   typ,
		value: value,
	}
}

func (this *userFunction) Type() Type           { return this.typ }
func (this *userFunction) Inspect() string      { return "user function" }
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

func (this *userFunction) Equal(other Function) bool {
	// TODO: return simply DeepEqual; no need to cast
	o, ok := other.(*userFunction)
	if !ok {
		return false
	} else {
		return reflect.DeepEqual(this.value, o.value)
	}
}
