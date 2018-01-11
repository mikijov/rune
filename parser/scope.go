package parser

import (
	"mikijov/rune/vm"
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

func (s *scope) Declare(name string, typ vm.Type) bool {
	if _, ok := s.store[name]; ok {
		return false // already declared in the current scope
	}
	s.store[name] = typ
	return true
}

func (s *scope) Get(name string) vm.Type {
	return s.store[name]
}
