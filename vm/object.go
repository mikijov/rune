package vm

import (
	"fmt"
)

const (
	INTEGER = "Integer"
	REAL    = "Real"
	BOOLEAN = "Boolean"
	STRING  = "String"
)

type Type string

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value VmInteger
}

func (i *Integer) Type() Type      { return INTEGER }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Real struct {
	Value VmReal
}

func (r *Real) Type() Type      { return REAL }
func (r *Real) Inspect() string { return fmt.Sprintf("%f", r.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type      { return BOOLEAN }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type String struct {
	Value string
}

func (s *String) Type() Type      { return STRING }
func (s *String) Inspect() string { return s.Value }

// type Function struct {
// 	Defn FunctionCode
// }
//
// func (f *Function) Type() string    { return FUNCTION }
// func (f *Function) Inspect() string { return f.String() }

// type Builtin struct {
// 	Fn BuiltinFunction
// }
//
// func (b *Builtin) Type() string    { return BUILTIN }
// func (b *Builtin) Inspect() string { return "builtin function" }

// type Array struct {
// 	Elements []Object
// }
//
// func (ao *Array) Type() string { return ARRAY }
// func (ao *Array) Inspect() string {
// 	var out bytes.Buffer
//
// 	elements := []string{}
// 	for _, e := range ao.Elements {
// 		elements = append(elements, e.Inspect())
// 	}
//
// 	out.WriteString("[")
// 	out.WriteString(strings.Join(elements, ", "))
// 	out.WriteString("]")
//
// 	return out.String()
// }
