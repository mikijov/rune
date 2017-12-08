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

type parameter struct {
	name  string
	type_ Type
}

func (this *parameter) GetName() string {
	return this.name
}

func (this *parameter) GetType() Type {
	return this.type_
}

type FunctionDeclaration interface {
	Statement
	AddParameter(name string, type_ Type)
	getParameters() []*parameter
	SetBody(s ScopeStatement)
	GetBody() ScopeStatement
	GetType() Type
}

type functionDeclaration struct {
	name       string
	params     []*parameter
	returnType Type
	body       ScopeStatement
}

func NewFunctionStatement(name string, returnType Type) FunctionDeclaration {
	return &functionDeclaration{
		name:       name,
		params:     make([]*parameter, 0, 5),
		returnType: returnType,
	}
}

func (this *functionDeclaration) AddParameter(name string, type_ Type) {
	this.params = append(this.params, &parameter{name, type_})
}

func (this *functionDeclaration) getParameters() []*parameter {
	return this.params
}

func (this *functionDeclaration) SetBody(body ScopeStatement) {
	this.body = body
}

func (this *functionDeclaration) GetBody() ScopeStatement {
	return this.body
}

func (this *functionDeclaration) GetType() Type {
	paramTypes := make([]Type, 0, len(this.params))
	for _, param := range this.params {
		paramTypes = append(paramTypes, param.GetType())
	}

	return GetFunctionType(paramTypes, this.returnType)
}

func (this *functionDeclaration) Execute(env Environment) {
	env.Declare(this.name, &function{this})
}

func (this *functionDeclaration) ToString() string {
	params := ""
	for _, param := range this.params {
		if len(params) > 0 {
			params += ", "
		}
		params += param.name + " :" + string(param.type_)
	}

	if this.returnType != VOID {
		return fmt.Sprintf("func %s(%s) :%s;", this.name, params, this.returnType)
	} else {
		return fmt.Sprintf("func %s(%s);", this.name, params)
	}
}

type ScopeStatement interface {
	Statement
	AddStatement(s Statement)
}

type scopeStatement struct {
	statements []Statement
}

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

func (this *scopeStatement) ToString() string {
	if len(this.statements) == 0 {
		return "{}"
	}

	retVal := "{ "
	for _, stmt := range this.statements {
		retVal += stmt.ToString()
	}
	retVal += " }"

	return retVal
}

type returnStatement struct {
	value Expression
}

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

func (this *returnStatement) ToString() string {
	if this.value != nil {
		return "return " + this.value.ToString() + ";"
	} else {
		return "return;"
	}
}

type IfStatement interface {
	Statement
	SetAlternative(alternative Statement)
}

type ifStatement struct {
	condition   Expression
	effect      Statement
	alternative Statement
}

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

func (this *ifStatement) ToString() string {
	retVal := "if " + this.condition.ToString() + " " + this.effect.ToString()
	if this.alternative != nil {
		retVal += " else " + this.alternative.ToString()
	}
	return retVal
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
	fmt.Printf("%s :%s = %s\n", this.ToString(), value.Type(), value.Inspect())
}

func (this *expressionStatement) ToString() string {
	return this.expression.ToString() + ";"
}

// atoms
////////

type variableReference struct {
	name  string
	type_ Type
}

func (this *variableReference) Type() Type {
	return this.type_
}

func (this *variableReference) Execute(env Environment) Object {
	return env.Get(this.name)
}

func (this *variableReference) ToString() string {
	return this.name
}

func NewVariableReference(name string, type_ Type) Expression {
	return &variableReference{
		name:  name,
		type_: type_,
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
	fn := env.Get(this.name).(Function).GetValue()
	paramDefs := fn.getParameters()

	funcEnv := NewFunctionEnvironment(env, this.returnType)
	for i, param := range this.params {
		val := param.Execute(env)
		funcEnv.Declare(paramDefs[i].GetName(), val)
	}

	fn.GetBody().Execute(funcEnv)

	failed, msg := funcEnv.GetError()
	if failed {
		env.SetError(msg)
	}

	return funcEnv.GetReturnValue()
}

func (this *functionCall) ToString() string {
	retVal := this.name + "("

	first := true
	for _, param := range this.params {
		if first {
			first = false
		} else {
			retVal += ", "
		}
		retVal += param.ToString()
	}
	retVal += ")"

	return retVal
}

func NewFunctionCall(name string, params []Expression, returnType Type) Expression {
	return &functionCall{
		name:       name,
		params:     params,
		returnType: returnType,
	}
}

// literals
///////////

type integerLiteral struct {
	value Integer
}

func (this *integerLiteral) Type() Type {
	return this.value.Type()
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
	return this.value.Type()
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

type booleanLiteral struct {
	value Boolean
}

func (this *booleanLiteral) Type() Type {
	return this.value.Type()
}

func (this *booleanLiteral) Execute(env Environment) Object {
	return this.value
}

func (this *booleanLiteral) ToString() string {
	return fmt.Sprintf("%f", this.value.GetValue())
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

func NewZeroLiteral(t Type) Expression {
	switch t {
	case INTEGER:
		return &integerLiteral{value: &integer{value: VmInteger(0)}}
	case REAL:
		return &realLiteral{value: &real{value: VmReal(0.0)}}
	case BOOLEAN:
		return &booleanLiteral{value: &boolean{value: false}}
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

// |

type integerOr struct {
	left  Expression
	right Expression
}

func (this *integerOr) Type() Type {
	return INTEGER
}

func (this *integerOr) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() | right.GetValue()}
}

func (this *integerOr) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "|", this.right.ToString())
}

// ^

type integerXor struct {
	left  Expression
	right Expression
}

func (this *integerXor) Type() Type {
	return INTEGER
}

func (this *integerXor) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() ^ right.GetValue()}
}

func (this *integerXor) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "^", this.right.ToString())
}

// &

type integerAnd struct {
	left  Expression
	right Expression
}

func (this *integerAnd) Type() Type {
	return INTEGER
}

func (this *integerAnd) Execute(env Environment) Object {
	left := this.left.Execute(env).(Integer)
	right := this.right.Execute(env).(Integer)
	return &integer{left.GetValue() & right.GetValue()}
}

func (this *integerAnd) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), "&", this.right.ToString())
}

// or

type booleanOr struct {
	left  Expression
	right Expression
}

func (this *booleanOr) Type() Type {
	return BOOLEAN
}

func (this *booleanOr) Execute(env Environment) Object {
	left := this.left.Execute(env).(Boolean).GetValue()
	if left {
		return &boolean{true}
	}
	right := this.right.Execute(env).(Boolean).GetValue()
	return &boolean{right}
}

func (this *booleanOr) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), " or ", this.right.ToString())
}

// and

type booleanAnd struct {
	left  Expression
	right Expression
}

func (this *booleanAnd) Type() Type {
	return BOOLEAN
}

func (this *booleanAnd) Execute(env Environment) Object {
	left := this.left.Execute(env).(Boolean).GetValue()
	if !left {
		return &boolean{false}
	}
	right := this.right.Execute(env).(Boolean).GetValue()
	return &boolean{right}
}

func (this *booleanAnd) ToString() string {
	return fmt.Sprintf("(%s%s%s)", this.left.ToString(), " and ", this.right.ToString())
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
	default:
		panic("unsupported operand")
	}

	// fmt.Println(retVal.ToString())
	return retVal
}
