package vm

type Variable struct {
	value *Object
	type_ Type
}

type Scope struct {
	store      map[string]*Variable
	statements []Statement
	outer      *Scope
}

func NewScope(outer *Scope) *Scope {
	return &Scope{
		store: make(map[string]*Variable),
		outer: outer,
	}
}

func (s *Scope) AddStatement(stmt Statement) {
	s.statements = append(s.statements, stmt)
}

func (s *Scope) Declare(name string, type_ Type) (*Variable, bool) {
	if _, ok := s.store[name]; ok {
		return nil, false // already declared in the current scope
	}
	retVal := &Variable{
		type_: type_,
	}
	s.store[name] = retVal
	return retVal, true
}

func (s *Scope) Get(name string) (*Variable, bool) {
	obj, ok := s.store[name]
	if !ok && s.outer != nil {
		obj, ok = s.outer.Get(name)
	}
	return obj, ok
}
