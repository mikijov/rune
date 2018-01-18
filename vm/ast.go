package vm

import (
	"fmt"
	"strconv"
)

// Program is a series of statements.
type Program interface {
	AddStatement(s Statement)
	Execute(env Environment)
	String() string
}

// Statement is a construct that does not have a value.
type Statement interface {
	Execute(env Environment)
	String() string
}

// Expression is a construct that has a value that can be assigned.
type Expression interface {
	Execute(env Environment) Object
	Type() Type
	String() string
}

// UnaryExpression expands on Expression and has only an operator and a single
// operand.
type UnaryExpression interface {
	Expression
	Operand() Object
}

// BinaryExpression expands on Expression and has operator and two operands.
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

// NewProgram creates new Program.
func NewProgram() Program {
	return &program{
		statements: make([]Statement, 0, 5),
	}
}

func (this *program) AddStatement(s Statement) {
	this.statements = append(this.statements, s)
}

func (this *program) Execute(env Environment) {
	for _, stmt := range this.statements {
		stmt.Execute(env)
		if env.IsReturning() {
			return
		}
	}
}

func (this *program) String() string {
	txt := ""
	for _, stmt := range this.statements {
		txt += stmt.String() + "\n"
	}
	return txt
}

// statements
/////////////

type declarationStatement struct {
	name       string
	expression Expression
}

// NewDeclarationStatement creates new declaration Statement. When executed it
// will declare new variable in the current scope and then assign initial value
// to it.
func NewDeclarationStatement(name string, value Expression) Statement {
	return &declarationStatement{
		name:       name,
		expression: value,
	}
}

func (this *declarationStatement) Execute(env Environment) {
	value := this.expression.Execute(env)
	env.Declare(this.name, value)
	// fmt.Printf("%s :%v = %s;\n", this.name, value.Type(), value.String())
}

func (this *declarationStatement) String() string {
	return fmt.Sprintf("%s :%v = %s;",
		this.name,
		this.expression.Type(),
		this.expression.String(),
	)
}

// ScopeStatement is a statement representing a list of statements all enclosed
// within a scope that determines variable lifetime.
type ScopeStatement interface {
	Statement
	AddStatement(s Statement)
}

type scopeStatement struct {
	statements []Statement
}

// NewScopeStatement creates new Statement representing a scope. Most importantly
// a scope determines lifetime of variables. Like a Program, scope consists of
// statements.
func NewScopeStatement() ScopeStatement {
	return &scopeStatement{
		statements: make([]Statement, 0, 5),
	}
}

func (this *scopeStatement) AddStatement(s Statement) {
	this.statements = append(this.statements, s)
}

func (this *scopeStatement) Execute(env Environment) {
	// TODO: Need new env for nested scopes, but this creates needless scope
	// for function calls. Can I avoid this?
	env = NewEnvironment(env)

	for _, stmt := range this.statements {
		stmt.Execute(env)
		if env.IsReturning() {
			return
		}
	}
}

func (this *scopeStatement) String() string {
	if len(this.statements) == 0 {
		return "{}"
	}

	retVal := "{ "
	for _, stmt := range this.statements {
		retVal += stmt.String()
	}
	retVal += " }"

	return retVal
}

type returnStatement struct {
	value Expression
}

// NewReturnStatement creates new return Statement. Return when executed, return
// statement indicates to the VM that the current function should return. In
// addition it handles returning result value to the caller. It is illegal to
// execute a return statement outside of a function.
func NewReturnStatement(retVal Expression) Statement {
	return &returnStatement{
		value: retVal,
	}
}

func (this *returnStatement) Execute(env Environment) {
	if this.value != nil {
		env.SetReturnValue(this.value.Execute(env))
	}

	env.SetReturning()
}

func (this *returnStatement) String() string {
	if this.value != nil {
		return "return " + this.value.String() + ";"
	}
	return "return;"
}

// IfStatement represents an if statement. It has a condition which determines
// if the effect is executed. Optionally is also has an alternative. Alternative
// is the 'else' portion of the if statement and is executed if the condition
// is not true.
type IfStatement interface {
	Statement
	SetAlternative(alternative Statement)
}

type ifStatement struct {
	condition   Expression
	effect      Statement
	alternative Statement
}

// NewIfStatement creates an if statement.
func NewIfStatement(condition Expression, effect, alternative Statement) IfStatement {
	return &ifStatement{
		condition:   condition,
		effect:      effect,
		alternative: alternative,
	}
}

func (this *ifStatement) SetAlternative(alternative Statement) {
	this.alternative = alternative
}

func (this *ifStatement) Execute(env Environment) {
	flag := this.condition.Execute(env).(Boolean).GetValue()

	if flag {
		this.effect.Execute(env)
	} else if this.alternative != nil {
		this.alternative.Execute(env)
	}
}

func (this *ifStatement) String() string {
	retVal := "if " + this.condition.String() + " " + this.effect.String()
	if this.alternative != nil {
		retVal += " else " + this.alternative.String()
	}
	return retVal
}

type loopStatement struct {
	label string
	body  Statement
}

// NewLoopStatement creates a loop statement. Simple loop statement does not
// have a condition and will be executed infinitely or until break or return
// statement is executed.
func NewLoopStatement(label string, body Statement) Statement {
	return &loopStatement{
		label: label,
		body:  body,
	}
}

func (this *loopStatement) Execute(env Environment) {
	for {
		this.body.Execute(env)
		if env.IsReturning() {
			return
		}
	}
}

func (this *loopStatement) String() string {
	retVal := "loop "
	if this.label != "" {
		retVal += this.label + " "
	}
	retVal += this.body.String()
	return retVal
}

type whileStatement struct {
	label     string
	condition Expression
	body      Statement
}

// NewWhileStatement creates a while loop statement. Condition is executed before
// each iteration through the body, and the body is executed only if the condition
// is true. This means if the first condition check fails, the body is never
// executed.
func NewWhileStatement(label string, condition Expression, body Statement) Statement {
	return &whileStatement{
		label:     label,
		condition: condition,
		body:      body,
	}
}

func (this *whileStatement) Execute(env Environment) {
	for this.condition.Execute(env).(Boolean).GetValue() {
		this.body.Execute(env)
		if env.IsReturning() {
			return
		}
	}
}

func (this *whileStatement) String() string {
	retVal := "loop "
	if this.label != "" {
		retVal += this.label + " "
	}
	retVal += "while " + this.body.String()
	return retVal
}

type untilStatement struct {
	label     string
	condition Expression
	body      Statement
}

// NewUntilStatement creates an until loop statement. Condition is executed after
// each iteration through the body. This means that the body is always executed
// at least once.
func NewUntilStatement(label string, condition Expression, body Statement) Statement {
	return &untilStatement{
		label:     label,
		condition: condition,
		body:      body,
	}
}

func (this *untilStatement) Execute(env Environment) {
	for this.condition.Execute(env).(Boolean).GetValue() {
		this.body.Execute(env)
		if env.IsReturning() {
			return
		}
	}
}

func (this *untilStatement) String() string {
	retVal := "loop "
	if this.label != "" {
		retVal += this.label + " "
	}
	retVal += "until " + this.body.String()
	return retVal
}

type expressionStatement struct {
	expression Expression
}

// NewExpressionStatement creates a statement that evaluates an expression. The
// value of the expression is ignored. Example of an expression is '5;' whose
// value is 5 but nothing is done with it. The most common use of an expression
// statement is to call a function but ignoring it's return value.
func NewExpressionStatement(e Expression) Statement {
	return &expressionStatement{expression: e}
}

func (this *expressionStatement) Execute(env Environment) {
	this.expression.Execute(env)
	// value := this.expression.Execute(env)
	// fmt.Printf("%s :%v = %s\n", this.String(), value.Type(), value.String())
}

func (this *expressionStatement) String() string {
	return this.expression.String() + ";"
}

// atoms
////////

type variableReference struct {
	name string
	typ  Type
}

func (this *variableReference) Type() Type {
	return this.typ
}

func (this *variableReference) Execute(env Environment) Object {
	return env.Get(this.name)
}

func (this *variableReference) String() string {
	return this.name
}

// NewVariableReference creates expression which evaluates to value of variable
// in scope.
func NewVariableReference(name string, typ Type) Expression {
	return &variableReference{
		name: name,
		typ:  typ,
	}
}

type functionCall struct {
	Expression
	name       string
	params     []Expression
	returnType Type
}

func (this *functionCall) Type() Type {
	return this.returnType
}

func (this *functionCall) Execute(env Environment) Object {
	return env.Get(this.name).(Function).Call(env, this.params)
}

func (this *functionCall) String() string {
	retVal := this.name + "("

	first := true
	for _, param := range this.params {
		if first {
			first = false
		} else {
			retVal += ", "
		}
		retVal += param.String()
	}
	retVal += ")"

	return retVal
}

// NewFunctionCall creates new expression which evaluates to the return value of
// the function it is calling.
func NewFunctionCall(name string, params []Expression, returnType Type) Expression {
	return &functionCall{
		name:       name,
		params:     params,
		returnType: returnType,
	}
}

// literals
///////////

type literal struct {
	Expression
	value Object
}

func (this *literal) Type() Type {
	return this.value.Type()
}

func (this *literal) Execute(env Environment) Object {
	return this.value
}

func (this *literal) String() string {
	return this.value.String()
}

// NewLiteral creates expression which evaluates to the literal value.
func NewLiteral(value Object) Expression {
	return &literal{value: value}
}

type voidLiteral struct {
}

func (this *voidLiteral) Type() Type {
	return NewSimpleType(VOID)
}

func (this *voidLiteral) Execute(env Environment) Object {
	panic("cannot execute void literal")
}

func (this *voidLiteral) String() string {
	return string(VOID)
}

func NewVoidLiteral() Expression {
	return &voidLiteral{}
}

type integerLiteral struct {
	value Integer
}

func (this *integerLiteral) Type() Type {
	return this.value.Type()
}

func (this *integerLiteral) Execute(env Environment) Object {
	return this.value
}

func (this *integerLiteral) String() string {
	return this.value.String()
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
	return this.value.Type()
}

func (this *realLiteral) Execute(env Environment) Object {
	return this.value
}

func (this *realLiteral) String() string {
	return this.value.String()
}

func NewRealLiteral(s string) Expression {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return &realLiteral{value: &real{value: VmReal(val)}}
}

type booleanLiteral struct {
	value Boolean
}

func (this *booleanLiteral) Type() Type {
	return this.value.Type()
}

func (this *booleanLiteral) Execute(env Environment) Object {
	return this.value
}

func (this *booleanLiteral) String() string {
	return this.value.String()
}

func NewBooleanLiteral(s string) Expression {
	if s == "true" {
		return &booleanLiteral{value: &boolean{value: true}}
	} else if s == "false" {
		return &booleanLiteral{value: &boolean{value: false}}
	} else {
		panic("unexpected boolean value")
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
	return NewSimpleType(INTEGER)
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

func (this *integerAddition) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "+", this.right.String())
}

// +

type realAddition struct {
	left  Expression
	right Expression
}

func (this *realAddition) Type() Type {
	return NewSimpleType(REAL)
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

func (this *realAddition) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "+", this.right.String())
}

// -

type integerSubstraction struct {
	left  Expression
	right Expression
}

func (this *integerSubstraction) Type() Type {
	return NewSimpleType(INTEGER)
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

func (this *integerSubstraction) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "-", this.right.String())
}

// -

type realSubstraction struct {
	left  Expression
	right Expression
}

func (this *realSubstraction) Type() Type {
	return NewSimpleType(REAL)
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

func (this *realSubstraction) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "-", this.right.String())
}

// *

type integerMultiplication struct {
	left  Expression
	right Expression
}

func (this *integerMultiplication) Type() Type {
	return NewSimpleType(INTEGER)
}

func (this *integerMultiplication) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() * right.GetValue()}
}

func (this *integerMultiplication) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "*", this.right.String())
}

// *

type realMultiplication struct {
	left  Expression
	right Expression
}

func (this *realMultiplication) Type() Type {
	return NewSimpleType(REAL)
}

func (this *realMultiplication) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real)
	right := this.right.Execute(env).(Real)
	return &real{left.GetValue() * right.GetValue()}
}

func (this *realMultiplication) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "*", this.right.String())
}

// /

type integerDivision struct {
	left  Expression
	right Expression
}

func (this *integerDivision) Type() Type {
	return NewSimpleType(INTEGER)
}

func (this *integerDivision) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() / right.GetValue()}
}

func (this *integerDivision) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "/", this.right.String())
}

// /

type realDivision struct {
	left  Expression
	right Expression
}

func (this *realDivision) Type() Type {
	return NewSimpleType(REAL)
}

func (this *realDivision) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real)
	right := this.right.Execute(env).(Real)
	return &real{left.GetValue() / right.GetValue()}
}

func (this *realDivision) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "/", this.right.String())
}

// %

type integerModulo struct {
	left  Expression
	right Expression
}

func (this *integerModulo) Type() Type {
	return NewSimpleType(INTEGER)
}

func (this *integerModulo) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() % right.GetValue()}
}

func (this *integerModulo) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "%", this.right.String())
}

// |

type integerOr struct {
	left  Expression
	right Expression
}

func (this *integerOr) Type() Type {
	return NewSimpleType(INTEGER)
}

func (this *integerOr) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() | right.GetValue()}
}

func (this *integerOr) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "|", this.right.String())
}

// ^

type integerXor struct {
	left  Expression
	right Expression
}

func (this *integerXor) Type() Type {
	return NewSimpleType(INTEGER)
}

func (this *integerXor) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() ^ right.GetValue()}
}

func (this *integerXor) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "^", this.right.String())
}

// &

type integerAnd struct {
	left  Expression
	right Expression
}

func (this *integerAnd) Type() Type {
	return NewSimpleType(INTEGER)
}

func (this *integerAnd) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() & right.GetValue()}
}

func (this *integerAnd) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), "&", this.right.String())
}

// or

type booleanOr struct {
	left  Expression
	right Expression
}

func (this *booleanOr) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *booleanOr) Execute(env Environment) Object {
	left := this.left.Execute(env).(Boolean).GetValue()
	if left {
		return &boolean{true}
	}
	right := this.right.Execute(env).(Boolean).GetValue()
	return &boolean{right}
}

func (this *booleanOr) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " or ", this.right.String())
}

// and

type booleanAnd struct {
	left  Expression
	right Expression
}

func (this *booleanAnd) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *booleanAnd) Execute(env Environment) Object {
	left := this.left.Execute(env).(Boolean).GetValue()
	if !left {
		return &boolean{false}
	}
	right := this.right.Execute(env).(Boolean).GetValue()
	return &boolean{right}
}

func (this *booleanAnd) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " and ", this.right.String())
}

// equal

type integerEqual struct {
	left  Expression
	right Expression
}

func (this *integerEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *integerEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer).GetValue()
	right := this.right.Execute(env).(Integer).GetValue()
	return &boolean{left == right}
}

func (this *integerEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " == ", this.right.String())
}

type realEqual struct {
	left  Expression
	right Expression
}

func (this *realEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *realEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real).GetValue()
	right := this.right.Execute(env).(Real).GetValue()
	return &boolean{left == right}
}

func (this *realEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " == ", this.right.String())
}

type booleanEqual struct {
	left  Expression
	right Expression
}

func (this *booleanEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *booleanEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Boolean).GetValue()
	right := this.right.Execute(env).(Boolean).GetValue()
	return &boolean{left == right}
}

func (this *booleanEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " == ", this.right.String())
}

type funcEqual struct {
	left  Expression
	right Expression
}

func (this *funcEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *funcEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Function)
	right := this.right.Execute(env).(Function)
	return &boolean{left.Equal(right)}
}

func (this *funcEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " == ", this.right.String())
}

// notEqual

type integerNotEqual struct {
	left  Expression
	right Expression
}

func (this *integerNotEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *integerNotEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer).GetValue()
	right := this.right.Execute(env).(Integer).GetValue()
	return &boolean{left != right}
}

func (this *integerNotEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " != ", this.right.String())
}

type realNotEqual struct {
	left  Expression
	right Expression
}

func (this *realNotEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *realNotEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real).GetValue()
	right := this.right.Execute(env).(Real).GetValue()
	return &boolean{left != right}
}

func (this *realNotEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " != ", this.right.String())
}

type booleanNotEqual struct {
	left  Expression
	right Expression
}

func (this *booleanNotEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *booleanNotEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Boolean).GetValue()
	right := this.right.Execute(env).(Boolean).GetValue()
	return &boolean{left != right}
}

func (this *booleanNotEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " != ", this.right.String())
}

type funcNotEqual struct {
	left  Expression
	right Expression
}

func (this *funcNotEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *funcNotEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Function)
	right := this.right.Execute(env).(Function)
	return &boolean{!left.Equal(right)}
}

func (this *funcNotEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " != ", this.right.String())
}

// lessThan

type integerLessThan struct {
	left  Expression
	right Expression
}

func (this *integerLessThan) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *integerLessThan) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer).GetValue()
	right := this.right.Execute(env).(Integer).GetValue()
	return &boolean{left < right}
}

func (this *integerLessThan) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " < ", this.right.String())
}

type realLessThan struct {
	left  Expression
	right Expression
}

func (this *realLessThan) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *realLessThan) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real).GetValue()
	right := this.right.Execute(env).(Real).GetValue()
	return &boolean{left < right}
}

func (this *realLessThan) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " < ", this.right.String())
}

// greaterThan

type integerGreaterThan struct {
	left  Expression
	right Expression
}

func (this *integerGreaterThan) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *integerGreaterThan) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer).GetValue()
	right := this.right.Execute(env).(Integer).GetValue()
	return &boolean{left > right}
}

func (this *integerGreaterThan) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " > ", this.right.String())
}

type realGreaterThan struct {
	left  Expression
	right Expression
}

func (this *realGreaterThan) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *realGreaterThan) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real).GetValue()
	right := this.right.Execute(env).(Real).GetValue()
	return &boolean{left > right}
}

func (this *realGreaterThan) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " > ", this.right.String())
}

// lessOrEqual

type integerLessOrEqual struct {
	left  Expression
	right Expression
}

func (this *integerLessOrEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *integerLessOrEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer).GetValue()
	right := this.right.Execute(env).(Integer).GetValue()
	return &boolean{left <= right}
}

func (this *integerLessOrEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " <= ", this.right.String())
}

type realLessOrEqual struct {
	left  Expression
	right Expression
}

func (this *realLessOrEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *realLessOrEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real).GetValue()
	right := this.right.Execute(env).(Real).GetValue()
	return &boolean{left <= right}
}

func (this *realLessOrEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " <= ", this.right.String())
}

// greaterOrEqual

type integerGreaterOrEqual struct {
	left  Expression
	right Expression
}

func (this *integerGreaterOrEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *integerGreaterOrEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer).GetValue()
	right := this.right.Execute(env).(Integer).GetValue()
	return &boolean{left >= right}
}

func (this *integerGreaterOrEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " >= ", this.right.String())
}

type realGreaterOrEqual struct {
	left  Expression
	right Expression
}

func (this *realGreaterOrEqual) Type() Type {
	return NewSimpleType(BOOLEAN)
}

func (this *realGreaterOrEqual) Execute(env Environment) Object {
	left := this.left.Execute(env).(Real).GetValue()
	right := this.right.Execute(env).(Real).GetValue()
	return &boolean{left >= right}
}

func (this *realGreaterOrEqual) String() string {
	return fmt.Sprintf("(%s%s%s)", this.left.String(), " >= ", this.right.String())
}

// NewBinaryExpression

func NewBinaryExpression(left Expression, op string, right Expression) Expression {
	lType := left.Type().GetKind()
	rType := right.Type().GetKind()

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
	case "|":
		if lType == INTEGER {
			retVal = &integerOr{left, right}
		} else {
			panic("unsupported type")
		}
	case "^":
		if lType == INTEGER {
			retVal = &integerXor{left, right}
		} else {
			panic("unsupported type")
		}
	case "&":
		if lType == INTEGER {
			retVal = &integerAnd{left, right}
		} else {
			panic("unsupported type")
		}
	case "or":
		if lType == BOOLEAN {
			retVal = &booleanOr{left, right}
		} else {
			panic("unsupported type")
		}
	case "and":
		if lType == BOOLEAN {
			retVal = &booleanAnd{left, right}
		} else {
			panic("unsupported type")
		}
	case "==":
		if lType == BOOLEAN {
			retVal = &booleanEqual{left, right}
		} else if lType == INTEGER {
			retVal = &integerEqual{left, right}
		} else if lType == REAL {
			retVal = &realEqual{left, right}
		} else if lType == FUNCTION {
			retVal = &funcEqual{left, right}
		} else {
			panic("unsupported type")
		}
	case "!=":
		if lType == BOOLEAN {
			retVal = &booleanNotEqual{left, right}
		} else if lType == INTEGER {
			retVal = &integerNotEqual{left, right}
		} else if lType == REAL {
			retVal = &realNotEqual{left, right}
		} else if lType == FUNCTION {
			retVal = &funcNotEqual{left, right}
		} else {
			panic("unsupported type")
		}
	case "<":
		if lType == INTEGER {
			retVal = &integerLessThan{left, right}
		} else if lType == REAL {
			retVal = &realLessThan{left, right}
		} else {
			panic("unsupported type")
		}
	case ">":
		if lType == INTEGER {
			retVal = &integerGreaterThan{left, right}
		} else if lType == REAL {
			retVal = &realGreaterThan{left, right}
		} else {
			panic("unsupported type")
		}
	case "<=":
		if lType == INTEGER {
			retVal = &integerLessOrEqual{left, right}
		} else if lType == REAL {
			retVal = &realLessOrEqual{left, right}
		} else {
			panic("unsupported type")
		}
	case ">=":
		if lType == INTEGER {
			retVal = &integerGreaterOrEqual{left, right}
		} else if lType == REAL {
			retVal = &realGreaterOrEqual{left, right}
		} else {
			panic("unsupported type")
		}
	default:
		panic("unsupported operand")
	}

	// fmt.Println(retVal.String())
	return retVal
}

type assignment struct {
	name       string
	expression Expression
}

func NewAssignment(name string, value Expression) Expression {
	return &assignment{
		name:       name,
		expression: value,
	}
}

func (this *assignment) Type() Type {
	return this.expression.Type()
}

func (this *assignment) Execute(env Environment) Object {
	value := this.expression.Execute(env)
	env.Set(this.name, value)
	return value
}

func (this *assignment) String() string {
	return fmt.Sprintf("%s = %s",
		this.name,
		this.expression.String(),
	)
}
