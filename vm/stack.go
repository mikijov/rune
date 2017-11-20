package vm

import ()

// type Stack struct {
// 	stack []Object
// }
//
// func NewStack() *Stack {
// 	return &Stack{stack: make([]Object, 0, 1024)}
// }
//
// func (s *Stack) PushInt(v vmInteger) {
// 	s.stack = append(s.stack, Integer{ival: v})
// }
//
// func (s *Stack) PopInt() vmInteger {
// 	top := len(s.stack) - 1
// 	retVal := Integer(s.stack[top]).Value
// 	s.stack = s.stack[:top]
// 	return retVal
// }
//
// func (s Stack) PeekInt() vmInteger {
// 	return Integer(s.stack[len(s.stack)-1]).Value
// }
//
// func (s *Stack) PushBool(v bool) {
// 	s.stack = append(s.stack, Boolean{Value: v})
// }
//
// func (s *Stack) PopBool() bool {
// 	top := len(s.stack) - 1
// 	retVal := Boolean(s.stack[top]).Value
// 	s.stack = s.stack[:top]
// 	return retVal
// }
//
// func (s Stack) PeekBool() bool {
// 	return Boolean(s.stack[len(s.stack)-1]).Value
// }
//
// func (s *Stack) PushString(v string) {
// 	s.stack = append(s.stack, String{Value: v})
// }
//
// func (s *Stack) PopString() string {
// 	top := len(s.stack) - 1
// 	retVal := String(s.stack[top]).Value
// 	s.stack = s.stack[:top]
// 	return retVal
// }
//
// func (s Stack) PeekString() string {
// 	return String(s.stack[len(s.stack)-1]).Value
// }
