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

// Constants identifying kinds of rune data types.
const (
	VOID     Kind = "void"
	INTEGER       = "int"
	REAL          = "real"
	BOOLEAN       = "bool"
	STRING        = "string"
	FUNCTION      = "func"
)

// Kind is a higher level of grouping of data types in rune. E.g. all functions
// belong to a FUNCTION kind even though each different function signature is a
// separate type.
type Kind string

// VmInteger determines the underlying type rune uses to represent integers.
type VmInteger int64

// VmReal determines the underlying type rune uses to represent real numbers.
type VmReal float64

// Type interface describes every type in rune.
type Type interface {
	GetKind() Kind
	String() string
	Equal(other Type) bool
	GetParamCount() int
	GetParam(i int) Type
	GetResultType() Type
	GetZero() Object
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
		return &integer{value: VmInteger(0)}
	case REAL:
		return &real{value: VmReal(0.0)}
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
