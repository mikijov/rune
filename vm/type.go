package vm

const (
	VOID     Kind = "void"
	INTEGER       = "int"
	REAL          = "real"
	BOOLEAN       = "bool"
	STRING        = "string"
	FUNCTION      = "func"
)

type Kind string
type VmInteger int64
type VmReal float64

type Type interface {
	GetName() string
	GetKind() Kind
	String() string
	Equal(other Type) bool
	GetParamCount() int
	GetParam(i int) Type
	GetResultType() Type
	GetZero() Object
}

type simpleType struct {
	name string
	kind Kind
}

func NewSimpleType(kind Kind) Type {
	return &simpleType{
		name: string(kind),
		kind: kind,
	}
}

func (this *simpleType) GetName() string {
	return this.name
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
		return &string_{value: ""}
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
	// TODO: is name really necessary?
	name       string
	paramTypes []Type
	returnType Type
}

func NewFunctionType(name string, paramTypes []Type, returnType Type) Type {
	return &functionType{
		name:       name,
		paramTypes: paramTypes,
		returnType: returnType,
	}
}

func (this *functionType) GetName() string {
	return this.name
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
