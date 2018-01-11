package vm

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Environment interface {
	DeclareUserFunction(name string, fn interface{}) error

	Declare(name string, value Object)
	Set(name string, value Object)
	Get(name string) Object

	SetReturnValue(value Object)
	GetReturnValue() Object
	SetReturning()
	IsReturning() bool
	SetError(msg string)
	GetError() (failed bool, msg string)

	GetInteger(name string) Integer
	GetReal(name string) Real

	SetInteger(name string, value Integer)
	SetReal(name string, value Real)
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

func (this *environment) DeclareUserFunction(name string, fn interface{}) error {
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
	for i := 0; i < typ.NumIn(); i += 1 {
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
	this.Declare(name, NewUserFunction(functionType, val))

	return nil
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

func (this *environment) SetInteger(name string, value Integer) {
	obj := this.GetInteger(name)
	if obj != nil {
		obj.SetValue(value.GetValue())
	} else {
		this.store[name] = value
	}
}

func (this *environment) SetReal(name string, value Real) {
	obj := this.GetReal(name)
	if obj != nil {
		obj.SetValue(value.GetValue())
	} else {
		this.store[name] = value
	}
}

func (this *environment) GetInteger(name string) Integer {
	obj, ok := this.store[name].(Integer)
	if !ok && this.outer != nil {
		return this.outer.GetInteger(name)
	}
	return obj
}

func (this *environment) GetReal(name string) Real {
	obj, ok := this.store[name].(Real)
	if !ok && this.outer != nil {
		return this.outer.GetReal(name)
	}
	return obj
}
