package vm

import (
	"fmt"
)

type Variable struct {
	value *Object
	type_ Type
}

type Scope interface {
	Declare(name string, type_ Type) (*Variable, bool)
	Get(name string) (*Variable, bool)
}

type scope struct {
	store map[string]*Variable
	outer Scope
}

func NewScope(outer Scope) Scope {
	return &scope{
		store: make(map[string]*Variable),
		outer: outer,
	}
}

func (s *scope) Declare(name string, type_ Type) (*Variable, bool) {
	if _, ok := s.store[name]; ok {
		fmt.Printf("%s already declared\n", name)
		return nil, false // already declared in the current scope
	}
	retVal := &Variable{
		type_: type_,
	}
	s.store[name] = retVal
	fmt.Printf("%s :%s\n", name, type_)
	return retVal, true
}

func (s *scope) Get(name string) (*Variable, bool) {
	obj, ok := s.store[name]
	if !ok && s.outer != nil {
		obj, ok = s.outer.Get(name)
	}
	return obj, ok
}
