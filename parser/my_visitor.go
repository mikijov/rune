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
	return
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
	fmt.Printf("%s(%s)\n", f.Name(), ctx.GetText())

	if ctx != nil {
		fmt.Printf("%T\n", ctx)
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
	currentFunction vm.FunctionDeclaration
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
	case *FunctionContext:
		return this.VisitFunction(child)
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
	case *LiteralPassthroughContext:
		return this.VisitLiteralPassthrough(ctx)
	case *RealLiteralContext:
		return this.VisitRealLiteral(ctx)
	case *IntegerLiteralContext:
		return this.VisitIntegerLiteral(ctx)
	case *VariableExpressionContext:
		return this.VisitVariableExpression(ctx)
	case *FunctionCallContext:
		return this.VisitFunctionCall(ctx)
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

	if type_ != "" && value != nil {
		if type_ != value.Type() {
			panic("type mismatch")
		}
	} else if type_ != "" {
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
		return vm.NewDeclarationStatement(name, vm.NewZeroLiteral(type_))
	}
}

func (this *MyVisitor) VisitFunction(ctx *FunctionContext) vm.Statement {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()

	var returnType vm.Type = vm.VOID
	if ctx.GetReturnType() != nil {
		returnType = this.VisitTypeName(ctx.GetReturnType())
	}

	oldFunction := this.currentFunction
	defer func() { this.currentFunction = oldFunction }()
	this.currentFunction = vm.NewFunctionStatement(name, returnType)

	oldScope := this.scope
	defer func() { this.scope = oldScope }()
	this.scope = newScope(this.scope)

	this.VisitParams(ctx.GetParams())

	this.currentFunction.SetBody(this.VisitScope(ctx.GetBody()))

	this.scope = oldScope // need to manually set it so that the declaration is made in proper scope
	this.scope.declare(name, this.currentFunction.GetType())

	return this.currentFunction
}

func (this *MyVisitor) VisitParams(ctx IParamDeclContext) {
	trace(ctx)

	for _, child := range ctx.GetParamGroup() {
		this.VisitCombinedParam(child)
	}
}

func (this *MyVisitor) VisitCombinedParam(ctx ICombinedParamContext) {
	trace(ctx)

	returnType := this.VisitTypeName(ctx.GetParamType())

	for _, name := range ctx.GetNames() {
		this.currentFunction.AddParameter(name.GetText(), returnType)
	}
}

func (this *MyVisitor) VisitTypeName(ctx ITypeNameContext) vm.Type {
	trace(ctx)

	switch ctx.GetText() {
	case "int":
		return vm.INTEGER
	case "real":
		return vm.REAL
	case "string":
		return vm.STRING
	case "bool":
		return vm.BOOLEAN
	default:
		panic("unknown type") // grammar should not allow this
	}
}

func (this *MyVisitor) VisitScope(ctx IScopeContext) vm.ScopeStatement {
	trace(ctx)

	scope := vm.NewScopeStatement()
	for _, stmt := range ctx.GetStatements() {
		scope.AddStatement(this.VisitStatement(stmt))
	}

	return scope
}

func (this *MyVisitor) VisitLiteralPassthrough(ctx *LiteralPassthroughContext) vm.Expression {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *RealLiteralContext:
		return this.VisitRealLiteral(child)
	case *IntegerLiteralContext:
		return this.VisitIntegerLiteral(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) vm.Expression {
	trace(ctx)

	left := this.VisitExpression(ctx.left)
	right := this.VisitExpression(ctx.right)

	return vm.NewBinaryExpression(left, ctx.op.GetText(), right)
}

// func (this *MyVisitor) VisitExpressionPassthrough(ctx *ExpressionPassthroughContext) interface{} {
// 	trace()
// 	return this.VisitChildren(ctx)
// }

// func (this *MyVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
// 	trace()
// 	return this.VisitChildren(ctx)
// }

func (this *MyVisitor) VisitRealLiteral(ctx *RealLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewRealLiteral(ctx.GetText())
}

func (this *MyVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewIntegerLiteral(ctx.GetText())
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
		returnType = vm.VOID
	} else {
		returnType = vm.GetFunctionReturnType(variable.type_)
	}

	params := make([]vm.Expression, 0, len(ctx.GetParams()))
	paramTypes := make([]vm.Type, 0, len(ctx.GetParams()))
	for _, ectx := range ctx.GetParams() {
		expression := this.VisitExpression(ectx)
		params = append(params, expression)
		paramTypes = append(paramTypes, expression.Type())
	}

	callType := vm.GetFunctionType(paramTypes, returnType)

	if declared && callType != variable.type_ {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("function parameter mismatch\nExpected: %s\n     Got: %s", variable.type_, callType))
	}

	return vm.NewFunctionCall(name, params, returnType)
}
