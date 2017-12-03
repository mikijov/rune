package vm

import (
	"fmt"
)

const (
	VOID    = "void"
	INTEGER = "int"
	REAL    = "real"
	BOOLEAN = "bool"
	STRING  = "string"
)

type VmInteger int64
type VmReal float64

type Type string

type Object interface {
	Type() Type
	Inspect() string
}

func Zero(type_ Type) Object {
	switch type_ {
	case INTEGER:
		return &integer{value: VmInteger(0)}
	case REAL:
		return &real{value: VmReal(0.0)}
	case STRING:
		return &string_{value: ""}
	case BOOLEAN:
		return &boolean{value: false}
	case VOID:
		return &void{}
	default:
		panic("invalid type")
	}
}

func Assign(dest, src Object) {
	switch dest := dest.(type) {
	case Integer:
		dest.SetValue(src.(Integer).GetValue())
	case Real:
		dest.SetValue(src.(Real).GetValue())
	default:
		panic("unknown type")
	}
}

type void struct {
}

func (this *void) Type() Type      { return VOID }
func (this *void) Inspect() string { return VOID }

type Integer interface {
	Object
	GetValue() VmInteger
	SetValue(value VmInteger)
}

type integer struct {
	value VmInteger
}

func (this *integer) Type() Type               { return INTEGER }
func (this *integer) Inspect() string          { return fmt.Sprintf("%d", this.value) }
func (this *integer) GetValue() VmInteger      { return this.value }
func (this *integer) SetValue(value VmInteger) { this.value = value }

type Real interface {
	Object
	GetValue() VmReal
	SetValue(value VmReal)
}

type real struct {
	value VmReal
}

func (this *real) Type() Type            { return REAL }
func (this *real) Inspect() string       { return fmt.Sprintf("%f", this.value) }
func (this *real) GetValue() VmReal      { return this.value }
func (this *real) SetValue(value VmReal) { this.value = value }

type Boolean interface {
	Object
	GetValue() bool
	SetValue(value bool)
}

type boolean struct {
	value bool
}

func (this *boolean) Type() Type          { return BOOLEAN }
func (this *boolean) Inspect() string     { return fmt.Sprintf("%t", this.value) }
func (this *boolean) GetValue() bool      { return this.value }
func (this *boolean) SetValue(value bool) { this.value = value }

type String interface {
	Object
	GetValue() string
	SetValue(value string)
}

type string_ struct {
	value string
}

func (this *string_) Type() Type            { return STRING }
func (this *string_) Inspect() string       { return this.value }
func (this *string_) GetValue() string      { return this.value }
func (this *string_) SetValue(value string) { this.value = value }

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
