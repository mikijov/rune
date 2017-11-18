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

func (this *integerLiteral) Type() vm.Type {
	return vm.INTEGER
}

func (this *integerLiteral) Value() vm.Object {
	return &this.value
}

func (this *integerLiteral) ToString() string {
	return fmt.Sprintf("%d", this.value.Value)
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

type realLiteral struct {
	value vm.Real
}

func (this *realLiteral) Type() vm.Type {
	return vm.REAL
}

func (this *realLiteral) Value() vm.Object {
	return &this.value
}

func (this *realLiteral) ToString() string {
	return fmt.Sprintf("%f", this.value.Value)
}

func NewRealLiteral(s string) Expression {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	var retVal Expression = &realLiteral{value: vm.Real{Value: vm.VmReal(val)}}
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

func (this *integerAddition) Type() vm.Type {
	return vm.INTEGER
}

func (this *integerAddition) Value() vm.Object {
	left := this.left.Value().(*vm.Integer)
	right := this.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value + right.Value}
}

func (this *integerAddition) Left() Expression {
	return this.left
}

func (this *integerAddition) Right() Expression {
	return this.right
}

func (this *integerAddition) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "+", this.right.ToString())
}

// +

type realAddition struct {
	left  Expression
	right Expression
}

func (this *realAddition) Type() vm.Type {
	return vm.REAL
}

func (this *realAddition) Value() vm.Object {
	left := this.left.Value().(*vm.Real)
	right := this.right.Value().(*vm.Real)
	return &vm.Real{left.Value + right.Value}
}

func (this *realAddition) Left() Expression {
	return this.left
}

func (this *realAddition) Right() Expression {
	return this.right
}

func (this *realAddition) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "+", this.right.ToString())
}

// -

type integerSubstraction struct {
	left  Expression
	right Expression
}

func (this *integerSubstraction) Type() vm.Type {
	return vm.INTEGER
}

func (this *integerSubstraction) Value() vm.Object {
	left := this.left.Value().(*vm.Integer)
	right := this.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value - right.Value}
}

func (this *integerSubstraction) Left() Expression {
	return this.left
}

func (this *integerSubstraction) Right() Expression {
	return this.right
}

func (this *integerSubstraction) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "-", this.right.ToString())
}

// -

type realSubstraction struct {
	left  Expression
	right Expression
}

func (this *realSubstraction) Type() vm.Type {
	return vm.REAL
}

func (this *realSubstraction) Value() vm.Object {
	left := this.left.Value().(*vm.Real)
	right := this.right.Value().(*vm.Real)
	return &vm.Real{left.Value - right.Value}
}

func (this *realSubstraction) Left() Expression {
	return this.left
}

func (this *realSubstraction) Right() Expression {
	return this.right
}

func (this *realSubstraction) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "-", this.right.ToString())
}

// *

type integerMultiplication struct {
	left  Expression
	right Expression
}

func (this *integerMultiplication) Type() vm.Type {
	return vm.INTEGER
}

func (this *integerMultiplication) Value() vm.Object {
	left := this.left.Value().(*vm.Integer)
	right := this.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value * right.Value}
}

func (this *integerMultiplication) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "*", this.right.ToString())
}

// *

type realMultiplication struct {
	left  Expression
	right Expression
}

func (this *realMultiplication) Type() vm.Type {
	return vm.REAL
}

func (this *realMultiplication) Value() vm.Object {
	left := this.left.Value().(*vm.Real)
	right := this.right.Value().(*vm.Real)
	return &vm.Real{left.Value * right.Value}
}

func (this *realMultiplication) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "*", this.right.ToString())
}

// /

type integerDivision struct {
	left  Expression
	right Expression
}

func (this *integerDivision) Type() vm.Type {
	return vm.INTEGER
}

func (this *integerDivision) Value() vm.Object {
	left := this.left.Value().(*vm.Integer)
	right := this.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value / right.Value}
}

func (this *integerDivision) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "/", this.right.ToString())
}

// /

type realDivision struct {
	left  Expression
	right Expression
}

func (this *realDivision) Type() vm.Type {
	return vm.REAL
}

func (this *realDivision) Value() vm.Object {
	left := this.left.Value().(*vm.Real)
	right := this.right.Value().(*vm.Real)
	return &vm.Real{left.Value / right.Value}
}

func (this *realDivision) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "/", this.right.ToString())
}

// %

type integerModulo struct {
	left  Expression
	right Expression
}

func (this *integerModulo) Type() vm.Type {
	return vm.INTEGER
}

func (this *integerModulo) Value() vm.Object {
	left := this.left.Value().(*vm.Integer)
	right := this.right.Value().(*vm.Integer)
	return &vm.Integer{left.Value % right.Value}
}

func (this *integerModulo) ToString() string {
	return fmt.Sprintf("%s %s %s", this.left.ToString(), "%", this.right.ToString())
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
		} else if lType == vm.REAL {
			retVal = &realAddition{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "-":
		if lType == vm.INTEGER {
			retVal = &integerSubstraction{left: left, right: right}
		} else if lType == vm.REAL {
			retVal = &realSubstraction{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "*":
		if lType == vm.INTEGER {
			retVal = &integerMultiplication{left, right}
		} else if lType == vm.REAL {
			retVal = &realMultiplication{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "/":
		if lType == vm.INTEGER {
			retVal = &integerDivision{left, right}
		} else if lType == vm.REAL {
			retVal = &realDivision{left: left, right: right}
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
