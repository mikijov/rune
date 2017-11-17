package parser

import (
	"fmt"
	"strconv"

	"mikijov/rune-antlr/vm"
)

type Module interface {
	AddStatement(s Statement)
	Execute()
	ToString() string
}

type Statement interface {
	Execute()
	ToString() string
}

type Expression interface {
	Value() vm.Object
	Type() vm.Type
	ToString() string
}

type UnaryExpression interface {
	Expression
	Operand() vm.Object
}

type BinaryExpression interface {
	Expression
	Left() Expression
	Right() Expression
}

// module
/////////

type module struct {
	statements []Statement
}

func NewModule() *module {
	return &module{
		statements: make([]Statement, 0, 10),
	}
}

func (m *module) AddStatement(s Statement) {
	m.statements = append(m.statements, s)
}

func (m *module) Execute() {
	for _, stmt := range m.statements {
		stmt.Execute()
	}
}

func (m *module) ToString() string {
	txt := ""
	for _, stmt := range m.statements {
		txt += stmt.ToString() + "\n"
	}
	return txt
}

// statements
/////////////

type expressionStatement struct {
	expression Expression
}

func (s *expressionStatement) Execute() {
	value := s.expression.Value()
	fmt.Printf("Expression (%s): %s\n", value.Type(), value.Inspect())
}

func (s *expressionStatement) ToString() string {
	return s.expression.ToString() + ";"
}

// literals
///////////

type integerLiteral struct {
	value vm.Integer
}

func (i *integerLiteral) Type() vm.Type {
	return vm.INTEGER
}

func (i *integerLiteral) Value() vm.Object {
	return &i.value
}

func (i *integerLiteral) ToString() string {
	return fmt.Sprintf("%d", i.value.Value)
}

func NewIntegerLiteral(s string) Expression {
	val, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	var retVal Expression = &integerLiteral{value: vm.Integer{Value: vm.VmInteger(val)}}
	fmt.Println(retVal.ToString())
	return retVal
}

// expressions
//////////////

// +

type integerAddition struct {
	left  Expression
	right Expression
}

func (i *integerAddition) Type() vm.Type {
	return vm.INTEGER
}

func (i *integerAddition) Value() vm.Object {
	left := i.left.Value().(*vm.Integer)
	right := i.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value + right.Value}
}

func (i *integerAddition) Left() Expression {
	return i.left
}

func (i *integerAddition) Right() Expression {
	return i.right
}

func (i *integerAddition) ToString() string {
	return fmt.Sprintf("%s %s %s", i.left.ToString(), "+", i.right.ToString())
}

// -

type integerSubstraction struct {
	left  Expression
	right Expression
}

func (i *integerSubstraction) Type() vm.Type {
	return vm.INTEGER
}

func (i *integerSubstraction) Value() vm.Object {
	left := i.left.Value().(*vm.Integer)
	right := i.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value - right.Value}
}

func (i *integerSubstraction) Left() Expression {
	return i.left
}

func (i *integerSubstraction) Right() Expression {
	return i.right
}

func (i *integerSubstraction) ToString() string {
	return fmt.Sprintf("%s %s %s", i.left.ToString(), "-", i.right.ToString())
}

// *

type integerMultiplication struct {
	left  Expression
	right Expression
}

func (i *integerMultiplication) Type() vm.Type {
	return vm.INTEGER
}

func (i *integerMultiplication) Value() vm.Object {
	left := i.left.Value().(*vm.Integer)
	right := i.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value * right.Value}
}

func (i *integerMultiplication) ToString() string {
	return fmt.Sprintf("%s %s %s", i.left.ToString(), "*", i.right.ToString())
}

// /

type integerDivision struct {
	left  Expression
	right Expression
}

func (i *integerDivision) Type() vm.Type {
	return vm.INTEGER
}

func (i *integerDivision) Value() vm.Object {
	left := i.left.Value().(*vm.Integer)
	right := i.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value / right.Value}
}

func (i *integerDivision) ToString() string {
	return fmt.Sprintf("%s %s %s", i.left.ToString(), "/", i.right.ToString())
}

// %

type integerModulo struct {
	left  Expression
	right Expression
}

func (i *integerModulo) Type() vm.Type {
	return vm.INTEGER
}

func (i *integerModulo) Value() vm.Object {
	left := i.left.Value().(*vm.Integer)
	right := i.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value % right.Value}
}

func (i *integerModulo) ToString() string {
	return fmt.Sprintf("%s %s %s", i.left.ToString(), "%", i.right.ToString())
}

func NewBinaryExpression(left Expression, op string, right Expression) Expression {
	lType := left.Type()
	rType := right.Type()

	if lType != rType {
		panic("mismatching types")
	}

	var retVal Expression
	switch op {
	case "+":
		if lType == vm.INTEGER {
			retVal = &integerAddition{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "-":
		if lType == vm.INTEGER {
			retVal = &integerSubstraction{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "*":
		if lType == vm.INTEGER {
			retVal = &integerMultiplication{left, right}
		} else {
			panic("unsupported type")
		}
	case "/":
		if lType == vm.INTEGER {
			retVal = &integerDivision{left, right}
		} else {
			panic("unsupported type")
		}
	case "%":
		if lType == vm.INTEGER {
			retVal = &integerModulo{left, right}
		} else {
			panic("unsupported type")
		}
	default:
		panic("unsupported operand")
	}

	fmt.Println(retVal.ToString())
	return retVal
}
