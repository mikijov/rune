package vm

import (
	"fmt"
)

type Environment interface {
	Declare(name string, value Object)
	Set(name string, value Object)
	Get(name string) Object

	GetInteger(name string) Integer
	GetReal(name string) Real

	SetInteger(name string, value Integer)
	SetReal(name string, value Real)
}

type environment struct {
	store map[string]Object
	outer Environment
}

func NewEnvironment(outer Environment) Environment {
	return &environment{
		store: make(map[string]Object),
		outer: outer,
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
	obj := this.Get(name)
	if obj == nil {
		panic(fmt.Sprintf("undeclared variable: %s", name))
	}
	switch obj := obj.(type) {
	case Integer:
		obj.SetValue(value.(Integer).GetValue())
	case Real:
		obj.SetValue(value.(Real).GetValue())
	default:
		panic("unknown type")
	}
}

func (this *environment) Get(name string) Object {
	obj, ok := this.store[name]
	if !ok && this.outer != nil {
		return this.outer.Get(name)
	}
	return obj
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
