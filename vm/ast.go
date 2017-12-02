package vm

import (
	"fmt"
	"strconv"
)

type Program interface {
	AddStatement(s Statement)
	Execute(env Environment)
	ToString() string
}

type Statement interface {
	Execute(env Environment)
	ToString() string
}

type Expression interface {
	Execute(env Environment) Object
	Type() Type
	ToString() string
}

type UnaryExpression interface {
	Expression
	Operand() Object
}

type BinaryExpression interface {
	Expression
	Left() Expression
	Right() Expression
}

// program
//////////

type program struct {
	statements []Statement
}

func NewProgram() Program {
	return &program{
		statements: make([]Statement, 0, 10),
	}
}

func (this *program) AddStatement(s Statement) {
	this.statements = append(this.statements, s)
}

func (this *program) Execute(env Environment) {
	for _, stmt := range this.statements {
		stmt.Execute(env)
	}
}

func (this *program) ToString() string {
	txt := ""
	for _, stmt := range this.statements {
		txt += stmt.ToString() + "\n"
	}
	return txt
}

// statements
/////////////

type declarationStatement struct {
	name       string
	expression Expression
}

func NewDeclarationStatement(name string, value Expression) Statement {
	return &declarationStatement{
		name:       name,
		expression: value,
	}
}

func (this *declarationStatement) Execute(env Environment) {
	value := this.expression.Execute(env)
	env.Declare(this.name, value)
	fmt.Printf("%s :%s = %s;\n", this.name, value.Type(), value.Inspect())
}

func (this *declarationStatement) ToString() string {
	return fmt.Sprintf("%s :%s = %s;",
		this.name,
		this.expression.Type(),
		this.expression.ToString(),
	)
}

type assignmentStatement struct {
	name       string
	expression Expression
}

func NewAssignmentStatement(name string, value Expression) Statement {
	return &assignmentStatement{
		name:       name,
		expression: value,
	}
}

func (this *assignmentStatement) Execute(env Environment) {
	value := this.expression.Execute(env)
	env.Set(this.name, value)
	fmt.Printf("%s = %s;\n", this.name, value.Type(), value.Inspect())
}

func (this *assignmentStatement) ToString() string {
	return fmt.Sprintf("%s = %s;",
		this.name,
		this.expression.Type(),
		this.expression.ToString(),
	)
}

type expressionStatement struct {
	expression Expression
}

func NewExpressionStatement(e Expression) Statement {
	return &expressionStatement{expression: e}
}

func (this *expressionStatement) Execute(env Environment) {
	value := this.expression.Execute(env)
	fmt.Printf("Expression (%s): %s\n", value.Type(), value.Inspect())
}

func (this *expressionStatement) ToString() string {
	return this.expression.ToString() + ";"
}

// literals
///////////

type integerLiteral struct {
	value Integer
}

func (this *integerLiteral) Type() Type {
	return INTEGER
}

func (this *integerLiteral) Execute(env Environment) Object {
	return this.value
}

func (this *integerLiteral) ToString() string {
	return fmt.Sprintf("%d", this.value.GetValue())
}

func NewIntegerLiteral(s string) Expression {
	val, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		panic(err)
	}
	return &integerLiteral{value: &integer{value: VmInteger(val)}}
}

type realLiteral struct {
	value Real
}

func (this *realLiteral) Type() Type {
	return REAL
}

func (this *realLiteral) Execute(env Environment) Object {
	return this.value
}

func (this *realLiteral) ToString() string {
	return fmt.Sprintf("%f", this.value.GetValue())
}

func NewRealLiteral(s string) Expression {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return &realLiteral{value: &real{value: VmReal(val)}}
}

func NewZeroLiteral(t Type) Expression {
	switch t {
	case INTEGER:
		return &integerLiteral{value: &integer{value: VmInteger(0)}}
	case REAL:
		return &realLiteral{value: &real{value: VmReal(0.0)}}
	default:
		panic("invalid type")
	}
}

// expressions
//////////////

// +

type integerAddition struct {
	left  Expression
	right Expression
}

func (this *integerAddition) Type() Type {
	return INTEGER
}

func (this *integerAddition) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() + right.GetValue()}
}

func (this *integerAddition) Left() Expression {
	return this.left
}

func (this *integerAddition) Right() Expression {
	return this.right
}

func (this *integerAddition) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "+", this.right.ToString())
}

// +

type realAddition struct {
	left  Expression
	right Expression
}

func (this *realAddition) Type() Type {
	return REAL
}

func (this *realAddition) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real)
	right := this.right.Execute(env).(Real)
	return &real{left.GetValue() + right.GetValue()}
}

func (this *realAddition) Left() Expression {
	return this.left
}

func (this *realAddition) Right() Expression {
	return this.right
}

func (this *realAddition) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "+", this.right.ToString())
}

// -

type integerSubstraction struct {
	left  Expression
	right Expression
}

func (this *integerSubstraction) Type() Type {
	return INTEGER
}

func (this *integerSubstraction) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() - right.GetValue()}
}

func (this *integerSubstraction) Left() Expression {
	return this.left
}

func (this *integerSubstraction) Right() Expression {
	return this.right
}

func (this *integerSubstraction) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "-", this.right.ToString())
}

// -

type realSubstraction struct {
	left  Expression
	right Expression
}

func (this *realSubstraction) Type() Type {
	return REAL
}

func (this *realSubstraction) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real)
	right := this.right.Execute(env).(Real)
	return &real{left.GetValue() - right.GetValue()}
}

func (this *realSubstraction) Left() Expression {
	return this.left
}

func (this *realSubstraction) Right() Expression {
	return this.right
}

func (this *realSubstraction) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "-", this.right.ToString())
}

// *

type integerMultiplication struct {
	left  Expression
	right Expression
}

func (this *integerMultiplication) Type() Type {
	return INTEGER
}

func (this *integerMultiplication) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() * right.GetValue()}
}

func (this *integerMultiplication) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "*", this.right.ToString())
}

// *

type realMultiplication struct {
	left  Expression
	right Expression
}

func (this *realMultiplication) Type() Type {
	return REAL
}

func (this *realMultiplication) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real)
	right := this.right.Execute(env).(Real)
	return &real{left.GetValue() * right.GetValue()}
}

func (this *realMultiplication) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "*", this.right.ToString())
}

// /

type integerDivision struct {
	left  Expression
	right Expression
}

func (this *integerDivision) Type() Type {
	return INTEGER
}

func (this *integerDivision) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() / right.GetValue()}
}

func (this *integerDivision) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "/", this.right.ToString())
}

// /

type realDivision struct {
	left  Expression
	right Expression
}

func (this *realDivision) Type() Type {
	return REAL
}

func (this *realDivision) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real)
	right := this.right.Execute(env).(Real)
	return &real{left.GetValue() / right.GetValue()}
}

func (this *realDivision) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "/", this.right.ToString())
}

// %

type integerModulo struct {
	left  Expression
	right Expression
}

func (this *integerModulo) Type() Type {
	return INTEGER
}

func (this *integerModulo) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() % right.GetValue()}
}

func (this *integerModulo) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "%", this.right.ToString())
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
		if lType == INTEGER {
			retVal = &integerAddition{left: left, right: right}
		} else if lType == REAL {
			retVal = &realAddition{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "-":
		if lType == INTEGER {
			retVal = &integerSubstraction{left: left, right: right}
		} else if lType == REAL {
			retVal = &realSubstraction{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "*":
		if lType == INTEGER {
			retVal = &integerMultiplication{left, right}
		} else if lType == REAL {
			retVal = &realMultiplication{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "/":
		if lType == INTEGER {
			retVal = &integerDivision{left, right}
		} else if lType == REAL {
			retVal = &realDivision{left: left, right: right}
		} else {
			panic("unsupported type")
		}
	case "%":
		if lType == INTEGER {
			retVal = &integerModulo{left, right}
		} else {
			panic("unsupported type")
		}
	default:
		panic("unsupported operand")
	}

	// fmt.Println(retVal.ToString())
	return retVal
}
