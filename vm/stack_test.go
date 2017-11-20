package vm

import (
// "testing"
)

// func TestStackTypes(t *testing.T) {
// 	stack := NewStack()
//
// 	iexp := 123
// 	stack.PushInt(iexp)
// 	ival := stack.PeekInt()
// 	if ival != iexp {
// 		t.Errorf("peek: expected %d got %d", iexp, ival)
// 	}
// 	ival = 0
// 	ival = stack.PopInt()
// 	if ival != iexp {
// 		t.Errorf("peek: expected %d got %d", iexp, ival)
// 	}
//
// 	bexp := true
// 	stack.PushBool(bexp)
// 	bval := stack.PeekBool()
// 	if bval != bexp {
// 		t.Errorf("peek: expected %t got %t", bexp, bval)
// 	}
// 	bval = false
// 	bval = stack.PopBool()
// 	if bval != bexp {
// 		t.Errorf("peek: expected %t got %t", bexp, bval)
// 	}
//
// 	sexp := "Hello, World!"
// 	stack.PushString(sexp)
// 	sval := stack.PeekString()
// 	if sval != sexp {
// 		t.Errorf("peek: expected %s got %s", sexp, sval)
// 	}
// 	sval = ""
// 	sval = stack.PopString()
// 	if sval != sexp {
// 		t.Errorf("peek: expected %s got %s", sexp, sval)
// 	}
// }
//
// func TestStackSize(t *testing.T) {
// 	stack := NewStack()
//
// 	for i := 0; i < 100000; i++ {
// 		stack.PushInt(i)
// 	}
//
// 	for i := 99999; i >= 0; i-- {
// 		if i != stack.PopInt() {
// 			t.Fatalf("mismatch at %d", i)
// 		}
// 	}
// }
