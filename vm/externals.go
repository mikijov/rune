package vm

import (
	"errors"
	"reflect"
	"regexp"
)

type Externals interface {
	GetDeclCount() int
	GetDecl(i int) (name string, value Object)

	DeclareUserFunction(name string, fn interface{}) error
}

type external struct {
	name  string
	value Object
}

type externals struct {
	decls []*external
}

func NewExternals() Externals {
	return &externals{
		decls: make([]*external, 0, 10),
	}
}

func (this *externals) GetDeclCount() int {
	return len(this.decls)
}

func (this *externals) GetDecl(i int) (name string, value Object) {
	decl := this.decls[i]
	return decl.name, decl.value
}

func (this *externals) DeclareUserFunction(name string, fn interface{}) error {
	matched, err := regexp.MatchString("^[a-zA-Z_][a-zA-Z0-9_]*$", name)
	if err != nil {
		return err
	} else if !matched {
		return errors.New("invalid name; must match ^[a-zA-Z_][a-zA-Z0-9_]*$")
	}

	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		return errors.New("not a function")
	}

	typ := val.Type()

	paramTypes := make([]Type, 0, typ.NumIn())
	for i := 0; i < typ.NumIn(); i += 1 {
		paramType, err := go2runeType(typ.In(i))
		if err != nil {
			return err
		}
		paramTypes = append(paramTypes, paramType)
	}

	var returnType Type
	if typ.NumOut() == 0 {
		returnType = NewSimpleType(VOID)
	} else if typ.NumOut() > 1 {
		return errors.New("function can have no or one return value")
	} else {
		returnType, err = go2runeType(typ.Out(0))
		if err != nil {
			return err
		}
	}

	functionType := NewFunctionType(paramTypes, returnType)
	functionValue := NewUserFunction(functionType, val)

	this.decls = append(this.decls, &external{
		name:  name,
		value: functionValue,
	})

	return nil
}
