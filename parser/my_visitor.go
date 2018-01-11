// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import (
	"fmt"
	"runtime"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"mikijov/rune/vm"
)

type ErrorListener interface {
	Error(line, col int, msg string)
}

func trace(ctx antlr.ParserRuleContext) {
	// return
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
	fmt.Printf("%s(%s)\n", f.Name(), ctx.GetText())
	if ctx != nil {
		fmt.Printf("%T\n", ctx)
		return
		for i, child := range ctx.GetChildren() {
			switch child := child.(type) {
			case antlr.ParserRuleContext:
				fmt.Printf("%d:%s\n", i, child.GetText())
			case *antlr.TerminalNodeImpl:
				fmt.Printf("%d:%s\n", i, child.GetText())
			default:
				fmt.Printf("Unknown\n")
			}
		}
	}
}

type MyVisitor struct {
	*antlr.BaseParseTreeVisitor
	errors          ErrorListener
	program         vm.Program
	scope           *scope
	currentFunction vm.Type
}

func NewMyVisitor(errors ErrorListener) *MyVisitor {
	return &MyVisitor{
		errors:  errors,
		program: vm.NewProgram(),
		scope:   newScope(nil),
	}
}

func (this *MyVisitor) VisitProgram(ctx *ProgramContext) vm.Program {
	trace(ctx)

	for _, stmt := range ctx.GetStatements() {
		this.program.AddStatement(this.VisitStatement(stmt))
	}

	return this.program
}

func (this *MyVisitor) VisitStatement(ctx IStatementContext) vm.Statement {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *ExpressionContext:
		return vm.NewExpressionStatement(this.VisitExpression(child))
	case *DeclarationContext:
		return this.VisitDeclaration(child)
	case *LoopContext:
		return this.VisitLoop(child)
	case *FunctionContext:
		return this.VisitFunction(child)
	case *ReturnStatementContext:
		return this.VisitReturnStatement(child)
	case *IfStatementContext:
		return this.VisitIfStatement(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) VisitExpression(ctx interface{}) vm.Expression {
	switch ctx := ctx.(type) {
	case *ExpressionContext:
		return this.VisitExpression(ctx.GetChild(0))
	case *BinaryExpressionContext:
		return this.VisitBinaryExpression(ctx)
	case *AssignmentContext:
		return this.VisitAssignment(ctx)
	case *LiteralPassthroughContext:
		return this.VisitLiteralPassthrough(ctx)
	case *RealLiteralContext:
		return this.VisitRealLiteral(ctx)
	case *IntegerLiteralContext:
		return this.VisitIntegerLiteral(ctx)
	case *BooleanLiteralContext:
		return this.VisitBooleanLiteral(ctx)
	case *VariableExpressionContext:
		return this.VisitVariableExpression(ctx)
	case *FunctionCallContext:
		return this.VisitFunctionCall(ctx)
	case *LambdaContext:
		return this.VisitLambda(ctx)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx))
	}
}

func (this *MyVisitor) VisitDeclaration(ctx *DeclarationContext) vm.Statement {
	trace(ctx)

	var type_ vm.Type
	if ctx.GetType_() != nil {
		type_ = this.VisitTypeName(ctx.GetType_())
	}

	var value vm.Expression
	if ctx.GetValue() != nil {
		value = this.VisitExpression(ctx.GetValue())
	}

	if type_ != nil && value != nil {
		if !type_.Equal(value.Type()) {
			panic("type mismatch")
		}
	} else if type_ != nil {
		// already have type, nothing to do
	} else if value != nil {
		type_ = value.Type()
	} else {
		panic("variable type not provided") // grammar should not allow this
	}

	name := ctx.GetIdentifier().GetText()
	_, ok := this.scope.declare(name, type_)
	if !ok {
		panic("variable redeclared")
	}

	if value != nil {
		return vm.NewDeclarationStatement(name, value)
	} else {
		return vm.NewDeclarationStatement(name, vm.NewLiteral(type_.GetZero()))
	}
}

func (this *MyVisitor) VisitFunction(ctx *FunctionContext) vm.Statement {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()

	paramNames, paramTypes := this.VisitParams(ctx.GetParams())
	// len(paramNames) == len(paramTypes)

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.VisitTypeName(ctx.GetReturnType())
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	typ := vm.NewFunctionType(name, paramTypes, returnType)
	// declaration of this function is made in the proper, outer scope
	this.scope.declare(name, typ)

	// now prepare the scope before parsing function body which will reference
	// parameters etc.

	oldFunction := this.currentFunction
	defer func() { this.currentFunction = oldFunction }()
	this.currentFunction = typ

	oldScope := this.scope
	defer func() { this.scope = oldScope }()
	this.scope = newScope(this.scope)
	// add parameters as named variables
	for i := 0; i < len(paramNames); i++ {
		this.scope.declare(paramNames[i], paramTypes[i])
	}

	return vm.NewDeclarationStatement(
		name, vm.NewLiteral(
			vm.NewFunction(typ, paramNames, this.VisitScope(ctx.GetBody())),
		),
	)
}

type param struct {
	name string
	typ  vm.Type
}

func (this *MyVisitor) VisitParams(ctx IParamDeclContext) (paramNames []string, paramTypes []vm.Type) {
	trace(ctx)

	names := make([]string, 0, 10)
	types := make([]vm.Type, 0, 10)

	for _, child := range ctx.GetParamGroup() {
		newNames, newTypes := this.VisitCombinedParam(child)
		for _, name := range newNames {
			names = append(names, name)
		}
		for _, typ := range newTypes {
			types = append(types, typ)
		}
	}

	return names, types
}

func (this *MyVisitor) VisitCombinedParam(ctx ICombinedParamContext) (names []string, types []vm.Type) {
	trace(ctx)

	returnType := this.VisitTypeName(ctx.GetParamType())

	names = make([]string, 0, 3)
	types = make([]vm.Type, 0, 3)
	for _, name := range ctx.GetNames() {
		names = append(names, name.GetText())
		types = append(types, returnType)
	}

	return names, types
}

func (this *MyVisitor) VisitTypeName(ctx interface{}) vm.Type {
	switch ctx := ctx.(type) {
	case *TypeNameContext:
		return this.VisitTypeName(ctx.GetChild(0))
	case *SimpleTypeContext:
		return this.VisitSimpleType(ctx)
	case *FunctionTypeContext:
		return this.VisitFunctionType(ctx)
	default:
		panic("unknown type") // grammar should not allow this
	}
}

func (this *MyVisitor) VisitSimpleType(ctx *SimpleTypeContext) vm.Type {
	trace(ctx)

	switch ctx.GetText() {
	case "int":
		return vm.NewSimpleType(vm.INTEGER)
	case "real":
		return vm.NewSimpleType(vm.REAL)
	case "string":
		return vm.NewSimpleType(vm.STRING)
	case "bool":
		return vm.NewSimpleType(vm.BOOLEAN)
	default:
		panic("unknown type") // grammar should not allow this
	}
}

func (this *MyVisitor) VisitFunctionType(ctx *FunctionTypeContext) vm.Type {
	trace(ctx)

	paramTypes := make([]vm.Type, 0, 10)
	for _, param := range ctx.GetParamTypes() {
		paramTypes = append(paramTypes, this.VisitTypeName(param))
	}

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.VisitTypeName(ctx.GetReturnType())
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	return vm.NewFunctionType("", paramTypes, returnType)
}

func (this *MyVisitor) VisitScope(ctx IScopeContext) vm.ScopeStatement {
	trace(ctx)

	scope := vm.NewScopeStatement()
	for _, stmt := range ctx.GetStatements() {
		scope.AddStatement(this.VisitStatement(stmt))
	}

	return scope
}

func (this *MyVisitor) VisitReturnStatement(ctx *ReturnStatementContext) vm.Statement {
	trace(ctx)

	var retVal vm.Expression
	var retType vm.Type
	if ctx.GetRetVal() != nil {
		retVal = this.VisitExpression(ctx.GetRetVal())
		retType = retVal.Type()
	} else {
		retType = vm.NewSimpleType(vm.VOID)
	}

	if this.currentFunction == nil {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			"return statement outside function declaration")
	} else {
		declType := this.currentFunction.GetResultType()

		if declType.GetKind() == vm.VOID {
			if retType.GetKind() != vm.VOID {
				token := ctx.GetRetVal().GetStart()
				this.errors.Error(token.GetLine(), token.GetColumn(),
					fmt.Sprintf("return value provided when none expected"))
			}
		} else if retType.GetKind() == vm.VOID {
			token := ctx.GetRetVal().GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				fmt.Sprintf("return value required"))
		} else if !declType.Equal(retType) {
			token := ctx.GetRetVal().GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				fmt.Sprintf("return value mismatch\nExpected: %v\n     Got: %v",
					declType, retVal.Type()))
		}
	}

	return vm.NewReturnStatement(retVal)
}

func (this *MyVisitor) VisitIfStatement(ctx IIfStatementContext) vm.Statement {
	trace(ctx)

	conditions := ctx.GetConditions()
	effects := ctx.GetEffects()

	var retVal vm.IfStatement
	var current vm.IfStatement

	for i, condCtx := range conditions {
		condition := this.VisitExpression(condCtx)
		if condition.Type().GetKind() != vm.BOOLEAN {
			token := condCtx.GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				"'if' condition must be a boolean expression, but it is "+condition.Type().String())
		}

		effect := this.VisitScope(effects[i])

		if retVal == nil {
			retVal = vm.NewIfStatement(condition, effect, nil)
			current = retVal
		} else {
			nestedIf := vm.NewIfStatement(condition, effect, nil)
			current.SetAlternative(nestedIf)
			current = nestedIf
		}
	}

	if ctx.GetAlternative() != nil {
		current.SetAlternative(this.VisitScope(ctx.GetAlternative()))
	}

	return retVal
}

func (this *MyVisitor) VisitLoop(ctx *LoopContext) vm.Statement {
	trace(ctx)

	var label string
	// if ctx.GetLabel() != nil {
	// 	label = ctx.GetLabel().GetText()
	// }

	var condition vm.Expression
	var isWhile bool
	if ctx.GetCondition() != nil {
		isWhile = ctx.GetKind().GetText() == "while"

		condition = this.VisitExpression(ctx.GetCondition())
		if condition.Type().GetKind() != vm.BOOLEAN {
			token := ctx.GetCondition().GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				"if condition must be a boolean expression, but it is "+condition.Type().String())
		}
	}

	body := this.VisitScope(ctx.GetBody())

	if condition == nil {
		return vm.NewLoopStatement(label, body)
	} else if isWhile {
		return vm.NewWhileStatement(label, condition, body)
	} else {
		return vm.NewUntilStatement(label, condition, body)
	}
}

func (this *MyVisitor) VisitLiteralPassthrough(ctx *LiteralPassthroughContext) vm.Expression {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *RealLiteralContext:
		return this.VisitRealLiteral(child)
	case *IntegerLiteralContext:
		return this.VisitIntegerLiteral(child)
	case *BooleanLiteralContext:
		return this.VisitBooleanLiteral(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) VisitAssignment(ctx *AssignmentContext) vm.Expression {
	trace(ctx)

	name := ctx.GetLeft().GetText()
	variable, declared := this.scope.get(name)
	if !declared {
		token := ctx.GetLeft()
		this.errors.Error(token.GetLine(), token.GetColumn(), name+": undeclared variable")
	}

	right := this.VisitExpression(ctx.right)

	if !variable.type_.Equal(right.Type()) {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("type mismatch\n Left: %s\nRight: %s", variable.type_, right.Type()))
	}

	switch ctx.GetOp().GetText() {
	case "=":
		return vm.NewAssignment(name, right)
	case "+=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "+", right))
	case "-=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "-", right))
	case "*=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "*", right))
	case "/=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "/", right))
	case "%=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "%", right))
	case "&=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "&", right))
	case "|=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "|", right))
	case "^=":
		return vm.NewAssignment(name,
			vm.NewBinaryExpression(vm.NewVariableReference(name, variable.type_), "^", right))
	default:
		panic("unknown operation")
	}
}

func (this *MyVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) vm.Expression {
	trace(ctx)

	left := this.VisitExpression(ctx.left)
	right := this.VisitExpression(ctx.right)

	return vm.NewBinaryExpression(left, ctx.op.GetText(), right)
}

func (this *MyVisitor) VisitRealLiteral(ctx *RealLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewRealLiteral(ctx.GetText())
}

func (this *MyVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewIntegerLiteral(ctx.GetText())
}

func (this *MyVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewBooleanLiteral(ctx.GetText())
}

func (this *MyVisitor) VisitVariableExpression(ctx *VariableExpressionContext) vm.Expression {
	trace(ctx)

	token := ctx.GetName()
	name := token.GetText()
	variable, declared := this.scope.get(name)
	if !declared {
		this.errors.Error(token.GetLine(), token.GetColumn(), name+": undeclared variable")
	}

	return vm.NewVariableReference(name, variable.type_)
}

func (this *MyVisitor) VisitFunctionCall(ctx *FunctionCallContext) vm.Expression {
	trace(ctx)

	token := ctx.GetName()
	name := token.GetText()
	variable, declared := this.scope.get(name)
	var returnType vm.Type
	if !declared {
		this.errors.Error(token.GetLine(), token.GetColumn(), name+": undeclared function")
		returnType = vm.NewSimpleType(vm.VOID)
	} else {
		returnType = variable.type_.GetResultType()
	}

	params := make([]vm.Expression, 0, len(ctx.GetParams()))
	paramTypes := make([]vm.Type, 0, len(ctx.GetParams()))
	for _, ectx := range ctx.GetParams() {
		expression := this.VisitExpression(ectx)
		params = append(params, expression)
		paramTypes = append(paramTypes, expression.Type())
	}

	callType := vm.NewFunctionType("", paramTypes, returnType)

	if declared && !callType.Equal(variable.type_) {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("function parameter mismatch\nExpected: %s\n     Got: %s", variable.type_, callType))
	}

	return vm.NewFunctionCall(name, params, returnType)
}

func (this *MyVisitor) VisitLambda(ctx *LambdaContext) vm.Expression {
	trace(ctx)

	paramNames, paramTypes := this.VisitParams(ctx.GetParams())
	// len(paramNames) == len(paramTypes)

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.VisitTypeName(ctx.GetReturnType())
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	typ := vm.NewFunctionType("", paramTypes, returnType)

	// now prepare the scope before parsing function body which will reference
	// parameters etc.

	oldFunction := this.currentFunction
	defer func() { this.currentFunction = oldFunction }()
	this.currentFunction = typ

	oldScope := this.scope
	defer func() { this.scope = oldScope }()
	this.scope = newScope(this.scope)
	// add parameters as named variables
	for i := 0; i < len(paramNames); i++ {
		this.scope.declare(paramNames[i], paramTypes[i])
	}

	return vm.NewLiteral(
		vm.NewFunction(typ, paramNames, this.VisitScope(ctx.GetBody())),
	)
}
