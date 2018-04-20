// Copyright © 2018 Milutin Jovanović jovanovic.milutin@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vm

import (
	"fmt"
)

// Constants identifying kinds of rune data types.
const (
	VOID      Kind = "void"
	INTEGER        = "int"
	REAL           = "real"
	BOOLEAN        = "bool"
	STRING         = "string"
	FUNCTION       = "func"
	ARRAY          = "array"
	STRUCT         = "struct"
	INTERFACE      = "interface"
)

// Kind is a higher level of grouping of data types in rune. E.g. all functions
// belong to a FUNCTION kind even though each different function signature is a
// separate type.
type Kind string

// Type interface describes every type in rune.
type Type interface {
	GetKind() Kind
	String() string
	Equal(other Type) bool
	GetElementType() Type
	GetParamCount() int
	GetParam(i int) Type
	GetResultType() Type
	GetZero() Object
	GetFieldCount() int
	GetFieldName(i int) string
	GetFieldType(i int) Type
	GetFieldIndex(name string) int
	GetFunctionCount() int
	GetFunction(i int) Function
	GetFunctionName(i int) string
	GetFunctionIndex(name string) int
	AddFunction(name string, fn Function) bool
	RemoveFunction(name string) bool
}

type simpleType struct {
	kind Kind
}

// NewSimpleType creates Type that represents simple rune type like integer or
// real. Currently all types other then functions are simple types.
func NewSimpleType(kind Kind) Type {
	return &simpleType{
		kind: kind,
	}
}

func (this *simpleType) GetKind() Kind {
	return this.kind
}

func (this *simpleType) String() string {
	return string(this.kind)
}

func (this *simpleType) Equal(other Type) bool {
	if o, ok := other.(*simpleType); !ok {
		return false
	} else {
		return this.kind == o.kind
	}
}

func (this *simpleType) GetElementType() Type {
	panic("not an array")
}

func (this *simpleType) GetParamCount() int {
	panic("not a function")
}

func (this *simpleType) GetParam(i int) Type {
	panic("not a function")
}

func (this *simpleType) GetResultType() Type {
	return this
}

func (this *simpleType) GetZero() Object {
	switch this.kind {
	case INTEGER:
		return &integer{value: 0}
	case REAL:
		return &real{value: 0.0}
	case STRING:
		return &strng{value: ""}
	case BOOLEAN:
		return &boolean{value: false}
	case VOID:
		return &void{}
	case FUNCTION:
		return NewZeroFunction(this)
	default:
		panic("invalid type")
	}
}

func (this *simpleType) GetFieldCount() int {
	panic("not a struct")
}

func (this *simpleType) GetFieldName(i int) string {
	panic("not a struct")
}

func (this *simpleType) GetFieldType(i int) Type {
	panic("not a struct")
}

func (this *simpleType) GetFieldIndex(name string) int {
	panic("not a struct")
}

func (this *simpleType) GetFunctionCount() int {
	panic("not an interface")
}

func (this *simpleType) GetFunction(i int) Function {
	panic("not an interface")
}

func (this *simpleType) GetFunctionName(i int) name string {
	panic("not an interface")
}

func (this *simpleType) GetFunctionIndex(name string) int {
	panic("not an interface")
}

func (this *simpleType) AddFunction(name string, fn Function) bool {
	panic("not an interface")
}

func (this *simpleType) RemoveFunction(name string) bool {
	panic("not an interface")
}

type arrayType struct {
	element Type
}

// NewArrayType creates Type that represents an array of elements of specified
// type.
func NewArrayType(element Type) Type {
	return &arrayType{
		element: element,
	}
}

func (this *arrayType) GetKind() Kind {
	return ARRAY
}

func (this *arrayType) String() string {
	return fmt.Sprintf("%v[]", this.element)
}

func (this *arrayType) Equal(other Type) bool {
	if o, ok := other.(*arrayType); !ok {
		return false
	} else {
		return this.element == o.element
	}
}

func (this *arrayType) GetElementType() Type {
	return this.element
}

func (this *arrayType) GetParamCount() int {
	panic("not a function")
}

func (this *arrayType) GetParam(i int) Type {
	panic("not a function")
}

func (this *arrayType) GetResultType() Type {
	return this
}

func (this *arrayType) GetZero() Object {
	return NewArray(this)
}

func (this *arrayType) GetFieldCount() int {
	panic("not a struct")
}

func (this *arrayType) GetFieldName(i int) string {
	panic("not a struct")
}

func (this *arrayType) GetFieldType(i int) Type {
	panic("not a struct")
}

func (this *arrayType) GetFieldIndex(name string) int {
	panic("not a struct")
}

func (this *arrayType) GetFunctionCount() int {
	panic("not an interface")
}

func (this *arrayType) GetFunction(i int) Function {
	panic("not an interface")
}

func (this *arrayType) GetFunctionName(i int) string {
	panic("not an interface")
}

func (this *arrayType) GetFunctionIndex(name string) int {
	panic("not an interface")
}

func (this *arrayType) AddFunction(name string, fn Function) bool {
	panic("not an interface")
}

func (this *arrayType) RemoveFunction(name string) bool {
	panic("not an interface")
}

type functionType struct {
	paramTypes []Type
	returnType Type
}

// NewFunctionType creates Type that represents functions.
func NewFunctionType(paramTypes []Type, returnType Type) Type {
	return &functionType{
		paramTypes: paramTypes,
		returnType: returnType,
	}
}

func (this *functionType) GetKind() Kind {
	return FUNCTION
}

func (this *functionType) String() string {
	retVal := "func ("

	first := true
	for _, t := range this.paramTypes {
		if !first {
			retVal += ", "
		} else {
			first = false
		}
		retVal += t.String()
	}

	retVal += ")"

	if this.returnType.GetKind() != VOID {
		retVal += " :" + this.returnType.String()
	}

	return retVal
}

func (this *functionType) Equal(other Type) bool {
	if o, ok := other.(*functionType); !ok {
		return false
	} else {
		if len(this.paramTypes) != len(o.paramTypes) {
			return false
		}
		for i := 0; i < len(this.paramTypes); i++ {
			if !this.paramTypes[i].Equal(o.paramTypes[i]) {
				return false
			}
		}
		return this.returnType.Equal(o.returnType)
	}
}

func (this *functionType) GetElementType() Type {
	panic("not an array")
}

func (this *functionType) GetParamCount() int {
	return len(this.paramTypes)
}

func (this *functionType) GetParam(i int) Type {
	return this.paramTypes[i]
}

func (this *functionType) GetResultType() Type {
	return this.returnType
}

func (this *functionType) GetZero() Object {
	return NewZeroFunction(this)
}

func (this *functionType) GetFieldCount() int {
	panic("not a struct")
}

func (this *functionType) GetFieldName(i int) string {
	panic("not a struct")
}

func (this *functionType) GetFieldType(i int) Type {
	panic("not a struct")
}

func (this *functionType) GetFieldIndex(name string) int {
	panic("not a struct")
}

func (this *functionType) GetFunctionCount() int {
	panic("not an interface")
}

func (this *functionType) GetFunction(i int) Function {
	panic("not an interface")
}

func (this *functionType) GetFunctionName(i int) string {
	panic("not an interface")
}

func (this *functionType) GetFunctionIndex(name string) int {
	panic("not an interface")
}

func (this *functionType) AddFunction(name string, fn Function) bool {
	panic("not an interface")
}

func (this *functionType) RemoveFunction(name string) bool {
	panic("not an interface")
}

type structType struct {
	fieldNames    []string
	fieldTypes    []Type
	functionNames []string
	functions     []Type
}

// NewStructType creates Type that represents structs.
func NewStructType(fieldNames []string, fieldTypes []Type) Type {
	return &structType{
		fieldNames: fieldNames,
		fieldTypes: fieldTypes,
	}
}

func (this *structType) GetKind() Kind {
	return STRUCT
}

func (this *structType) String() string {
	retVal := "struct {"

	for i := 0; i < len(this.fieldNames); i++ {
		retVal += this.fieldNames[i] + ":" + this.fieldTypes[i].String() + ";"
	}

	retVal += "}"

	return retVal
}

func (this *structType) Equal(other Type) bool {
	if o, ok := other.(*structType); !ok {
		return false
	} else {
		if len(this.fieldNames) != len(o.fieldNames) {
			return false
		}
		for i := 0; i < len(this.fieldNames); i++ {
			if this.fieldNames[i] != o.fieldNames[i] {
				return false
			}
			if !this.fieldTypes[i].Equal(o.fieldTypes[i]) {
				return false
			}
		}
		return true
	}
}

func (this *structType) GetElementType() Type {
	panic("not an array")
}

func (this *structType) GetParamCount() int {
	panic("not a struct")
}

func (this *structType) GetParam(i int) Type {
	panic("not a struct")
}

func (this *structType) GetResultType() Type {
	return this
}

func (this *structType) GetZero() Object {
	return NewStruct(this)
}

func (this *structType) GetFieldCount() int {
	return len(this.fieldNames)
}

func (this *structType) GetFieldName(i int) string {
	return this.fieldNames[i]
}

func (this *structType) GetFieldType(i int) Type {
	return this.fieldTypes[i]
}

func (this *structType) GetFieldIndex(name string) int {
	for i, field := range this.fieldNames {
		if field == name {
			return i
		}
	}
	return -1
}

func (this *structType) GetFunctionCount() int {
	return len(this.functions)
}

func (this *structType) GetFunction(i int) Function {
	return this.functions[i]
}

func (this *structType) GetFunctionName(i int) string {
	return this.functionNames[i]
}

func (this *structType) GetFunctionIndex(name string) int {
	for i, fnName := range this.functionNames {
		if fnName == name {
			return i
		}
	}
	return -1
}

func (this *structType) AddFunction(name string, fn Function) bool {
	if index := this.GetFunctionIndex(name); index != -1 {
		return false
	}
	this.functionNames = append(this.functionNames, name)
	this.functions = append(this.functions, fn)
	return true
}

func (this *structType) RemoveFunction(name string) bool {
	index := this.GetFunctionIndex(name)
	if index != len(this.functionNames) {
		panic("can remove only last added method to avoid problem of renumbered indexes")
		return false
	}
	this.functionNames = append(this.functionNames[:index], this.functionNames[index+1:]...)
	this.functions = append(this.functions[:index], this.functions[index+1:]...)
	return true
}

type interfaceType struct {
	functionNames []string
	functions     []Type
}

// NewInterfaceType creates Type that represents interface.
func NewInterfaceType(names []string, functions []Type) Type {
	return &interfaceType{
		functionNames: names,
		functions:     functions,
	}
}

func (this *interfaceType) GetKind() Kind {
	return INTERFACE
}

func (this *interfaceType) String() string {
	retVal := "interface {"

	for i := 0; i < len(this.functions); i++ {
		retVal += fmt.Sprintf("%v\n", this.functions[i])
	}

	retVal += "}"

	return retVal
}

func (this *interfaceType) Equal(other Type) bool {
	return this == other
}

func (this *interfaceType) GetElementType() Type {
	panic("not an array")
}

func (this *interfaceType) GetParamCount() int {
	panic("not a struct")
}

func (this *interfaceType) GetParam(i int) Type {
	panic("not a struct")
}

func (this *interfaceType) GetResultType() Type {
	return this
}

func (this *interfaceType) GetZero() Object {
	panic("not impelemted")
	// return NewReference(this)
}

func (this *interfaceType) GetFieldCount() int {
	panic("not a struct")
}

func (this *interfaceType) GetFieldName(i int) string {
	panic("not a struct")
}

func (this *interfaceType) GetFieldType(i int) Type {
	panic("not a struct")
}

func (this *interfaceType) GetFieldIndex(name string) int {
	panic("not a struct")
}

func (this *interfaceType) GetFunctionCount() int {
	return len(this.functions)
}

func (this *interfaceType) GetFunction(i int) Function {
	return this.functions[i]
}

func (this *interfaceType) GetFunctionName(i int) string {
	return this.functionNames[i]
}

func (this *interfaceType) GetFunctionIndex(name string) int {
	for i, fnName := range this.functionNames {
		if fnName == name {
			return i
		}
	}
	return -1
}

func (this *interfaceType) AddFunction(name string, fn Function) bool {
	panic("not a struct")
}

func (this *interfaceType) RemoveFunction(name string) bool {
	panic("not a struct")
}
