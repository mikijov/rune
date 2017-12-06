package vm

const (
	VOID    = "void"
	INTEGER = "int"
	REAL    = "real"
	BOOLEAN = "bool"
	STRING  = "string"
)

type Type string
type VmInteger int64
type VmReal float64

func GetFunctionType(paramTypes []Type, returnType Type) Type {
	retVal := "func ("

	first := true
	for _, t := range paramTypes {
		if !first {
			retVal += ", "
		} else {
			first = false
		}
		retVal += string(t)
	}

	retVal += ")"

	if returnType != VOID {
		retVal += " :" + string(returnType)
	}

	return Type(retVal)
}

func GetFunctionReturnType(def Type) Type {
	if def[:6] != "func (" {
		panic("not a function")
	}

	s := def[6:]
	bracketCount := 1
	for i, c := range s {
		if c == '(' {
			bracketCount += 1
		} else if c == ')' {
			bracketCount -= 1
			if bracketCount == 0 {
				s = s[i+1:]
				break
			}
		}
	}

	if len(s) == 0 {
		return VOID
	}

	if s[:2] != " :" {
		panic("malformed function type")
	}

	return Type(s[2:])
}
