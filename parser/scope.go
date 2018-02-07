package parser

import (
	"github.com/mikijov/rune/vm"
)

type Scope interface {
	Declare(name string, typ vm.Type) bool
	Get(name string) vm.Type
}

type scope struct {
	store map[string]vm.Type
	outer Scope
}

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
