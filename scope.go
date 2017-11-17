package main

// type Scope struct {
// 	store map[string]Object
// 	outer *Scope
// }
//
// func NewScope(outer *Scope) *Scope {
// 	return &Scope{
// 		store: make(map[string]Object),
// 		outer: outer,
// 	}
// }
//
// func (s *Scope) Get(name string) (Object, bool) {
// 	obj, ok := s.store[name]
// 	if !ok && s.outer != nil {
// 		obj, ok = s.outer.Get(name)
// 	}
// 	return obj, ok
// }
//
// func (s *Scope) Set(name string, val Object) Object {
// 	// TODO: check not to overwrite value
// 	s.store[name] = val
// 	return val
// }
