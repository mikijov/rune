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

func Zero(type_ Type) Object {
	switch type_ {
	case INTEGER:
		return &integer{value: VmInteger(0)}
	case REAL:
		return &real{value: VmReal(0.0)}
	case STRING:
		return &string_{value: ""}
	case BOOLEAN:
		return &boolean{value: false}
	case VOID:
		return &void{}
	default:
		if IsFunction(type_) {
			return nil
		} else {
			panic("invalid type")
		}
	}
}

func Assign(dest, src Object) {
	switch dest := dest.(type) {
	case Integer:
		dest.SetValue(src.(Integer).GetValue())
	case Real:
		dest.SetValue(src.(Real).GetValue())
	default:
		panic("unknown type")
	}
}

type void struct {
}

func (this *void) Type() Type           { return VOID }
func (this *void) Inspect() string      { return VOID }
func (this *void) Value() reflect.Value { return reflect.ValueOf(nil) }

type Integer interface {
	Object
	GetValue() VmInteger
	SetValue(value VmInteger)
}

type integer struct {
	value VmInteger
}

func (this *integer) Type() Type               { return INTEGER }
func (this *integer) Inspect() string          { return fmt.Sprintf("%d", this.value) }
func (this *integer) Value() reflect.Value     { return reflect.ValueOf(this.value) }
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

func (this *real) Type() Type            { return REAL }
func (this *real) Inspect() string       { return fmt.Sprintf("%f", this.value) }
func (this *real) Value() reflect.Value  { return reflect.ValueOf(this.value) }
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

func (this *boolean) Type() Type           { return BOOLEAN }
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

func (this *string_) Type() Type            { return STRING }
func (this *string_) Inspect() string       { return this.value }
func (this *string_) Value() reflect.Value  { return reflect.ValueOf(this.value) }
func (this *string_) GetValue() string      { return this.value }
func (this *string_) SetValue(value string) { this.value = value }

// type Callable interface {
// 	Call(env Environment, params []Expression) Object
// }

type Function interface {
	Object
	Call(env Environment, params []Expression) Object
	// GetValue() FunctionDeclaration
	// SetValue(value FunctionDeclaration)
	Equal(other Function) bool
}

type function struct {
	value FunctionDeclaration
}

func (this *function) Type() Type           { return this.value.GetType() }
func (this *function) Inspect() string      { return this.value.ToString() }
func (this *function) Value() reflect.Value { return reflect.ValueOf(this.value) }

// func (this *function) GetValue() FunctionDeclaration      { return this.value }
// func (this *function) SetValue(value FunctionDeclaration) { this.value = value }
func (this *function) Call(env Environment, params []Expression) Object {
	funcEnv := NewFunctionEnvironment(env, this.value.GetReturnType())

	paramDefs := this.value.getParameters()
	for i, param := range params {
		val := param.Execute(env)
		funcEnv.Declare(paramDefs[i].GetName(), val)
	}

	this.value.GetBody().Execute(funcEnv)

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
		return reflect.DeepEqual(this.value, o.value)
	}
}

// type UserFunction interface {
// 	Object
// 	Callable
// 	// GetValue() reflect.Value
// 	// SetValue(value reflect.Value)
// }

type userFunction struct {
	type_ Type
	value reflect.Value
}

func (this *userFunction) Type() Type           { return this.type_ }
func (this *userFunction) Inspect() string      { return "user function" }
func (this *userFunction) Value() reflect.Value { return this.value }

// func (this *userFunction) GetValue() reflect.Value      { return this.value }
// func (this *userFunction) SetValue(value reflect.Value) { this.value = value }
func (this *userFunction) Call(env Environment, params []Expression) Object {
	valueParams := make([]reflect.Value, 0, len(params))
	for _, param := range params {
		val := param.Execute(env)
		valueParams = append(valueParams, val.Value())
	}

	this.value.Call(valueParams)

	return nil
}

func (this *userFunction) Equal(other Function) bool {
	o, ok := other.(*userFunction)
	if !ok {
		return false
	} else {
		return reflect.DeepEqual(this.value, o.value)
	}
}
