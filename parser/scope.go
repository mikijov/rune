package parser

import (
	"fmt"

	"mikijov/rune-antlr/vm"
)

type variable struct {
	value *vm.Object
	type_ vm.Type
}

type scope struct {
	store map[string]*variable
	outer *scope
}

func newScope(outer *scope) *scope {
	return &scope{
		store: make(map[string]*variable),
		outer: outer,
	}
}

func (s *scope) declare(name string, type_ vm.Type) (*variable, bool) {
	if _, ok := s.store[name]; ok {
		fmt.Printf("%s already declared\n", name)
		return nil, false // already declared in the current scope
	}
	retVal := &variable{
		type_: type_,
	}
	s.store[name] = retVal
	// fmt.Printf("%s :%s\n", name, type_)
	return retVal, true
}

func (s *scope) get(name string) (*variable, bool) {
	obj, ok := s.store[name]
	if !ok && s.outer != nil {
		obj, ok = s.outer.get(name)
	}
	return obj, ok
}
