// Copyright © 2018 Milutin Jovanović jovanovic.milutin@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import (
	"fmt"
	// "runtime"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/mikijov/rune/vm"
)

// ErrorListener is the sink for all messages produced by the compiler. The
// implementation must be provided by the user of this package.
type ErrorListener interface {
	Error(line, col int, msg string)
}

// this function is here purely for early debugging. It should be removed as
// soon as the parser is relatively stable.
func trace(ctx antlr.ParserRuleContext) {
	// // return
	// pc := make([]uintptr, 10) // at least 1 entry needed
	// runtime.Callers(2, pc)
	// f := runtime.FuncForPC(pc[0])
	// // file, line := f.FileLine(pc[0])
	// // fmt.Printf("%s:%d %s\n", file, line, f.Name())
	// fmt.Printf("%s(%s)\n", f.Name(), ctx.GetText())
	// if ctx != nil {
	// 	fmt.Printf("%T\n", ctx)
	// 	return
	// 	for i, child := range ctx.GetChildren() {
	// 		switch child := child.(type) {
	// 		case antlr.ParserRuleContext:
	// 			fmt.Printf("%d:%s\n", i, child.GetText())
	// 		case *antlr.TerminalNodeImpl:
	// 			fmt.Printf("%d:%s\n", i, child.GetText())
	// 		default:
	// 			fmt.Printf("Unknown\n")
	// 		}
	// 	}
	// }
}

// MyVisitor is implementation of of BaseParseTreeVisitor that parses rune
// programs.
type MyVisitor struct {
	*antlr.BaseParseTreeVisitor
	errors          ErrorListener
	program         vm.Program
	scope           Scope
	currentFunction vm.Type
}

// NewMyVisitor creates new instance of MyVisitor and initializes it with error
// listener and externals.
func NewMyVisitor(errors ErrorListener, ext vm.Externals) *MyVisitor {
	retVal := &MyVisitor{
		errors:  errors,
		program: vm.NewProgram(),
		scope:   NewScope(nil),
	}

	if ext != nil {
		for i := 0; i < ext.GetDeclCount(); i++ {
			name, value := ext.GetDecl(i)
			retVal.scope.Declare(name, value.Type())
		}
	}

	return retVal
}

// VisitProgram is the entry point for parsing all rune programs.
func (this *MyVisitor) VisitProgram(ctx *ProgramContext) vm.Program {
	trace(ctx)

	for _, stmt := range ctx.GetStatements() {
		this.program.AddStatement(this.visitStatement(stmt))
	}

	return this.program
}

func (this *MyVisitor) visitStatement(ctx IStatementContext) vm.Statement {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *ExpressionContext:
		return vm.NewExpressionStatement(this.visitExpression(child))
	case *DeclarationContext:
		return this.visitDeclaration(child)
	case *LoopContext:
		return this.visitLoop(child)
	case *FunctionContext:
		return this.visitFunction(child)
	case *ReturnStatementContext:
		return this.visitReturnStatement(child)
	case *IfStatementContext:
		return this.visitIfStatement(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) visitExpression(ctx interface{}) vm.Expression {
	switch ctx := ctx.(type) {
	case *ExpressionContext:
		return this.visitExpression(ctx.GetChild(0))
	case *ExpressionPassthroughContext:
		return this.visitExpression(ctx.GetChild(1)) // skip parenthesis
	case *BinaryExpressionContext:
		return this.visitBinaryExpression(ctx)
	// case *UnaryExpressionContext:
	// 	return this.VisitUnaryExpression(ctx)
	case *AssignmentContext:
		return this.visitAssignment(ctx)
	case *LiteralPassthroughContext:
		return this.visitLiteralPassthrough(ctx)
	case *RealLiteralContext:
		return this.visitRealLiteral(ctx)
	case *IntegerLiteralContext:
		return this.visitIntegerLiteral(ctx)
	case *BooleanLiteralContext:
		return this.visitBooleanLiteral(ctx)
	case *VariableExpressionContext:
		return this.visitVariableExpression(ctx)
	case *FunctionCallContext:
		return this.visitFunctionCall(ctx)
	case *FieldSelectorContext:
		return this.visitFieldSelector(ctx)
	case *LambdaContext:
		return this.visitLambda(ctx)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx))
	}
}

func (this *MyVisitor) visitDeclaration(ctx *DeclarationContext) vm.Statement {
	trace(ctx)

	var typ vm.Type
	if ctx.GetType_() != nil {
		typ = this.visitTypeName(ctx.GetType_())
	}

	var value vm.Expression
	if ctx.GetValue() != nil {
		value = this.visitExpression(ctx.GetValue())
	}

	if typ != nil && value != nil {
		if !typ.Equal(value.Type()) {
			panic("type mismatch")
		}
	} else if typ != nil {
		// already have type, nothing to do
	} else if value != nil {
		typ = value.Type()
	} else {
		panic("variable type not provided") // grammar should not allow this
	}

	name := ctx.GetIdentifier().GetText()
	if !this.scope.Declare(name, typ) {
		panic("variable redeclared")
	}

	if value != nil {
		return vm.NewDeclarationStatement(name, value)
	}
	return vm.NewDeclarationStatement(name, vm.NewLiteral(typ.GetZero()))
}

func (this *MyVisitor) visitFunction(ctx *FunctionContext) vm.Statement {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()

	paramNames, paramTypes := this.visitParams(ctx.GetParams())
	// len(paramNames) == len(paramTypes)

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.visitTypeName(ctx.GetReturnType())
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	typ := vm.NewFunctionType(paramTypes, returnType)
	// declaration of this function is made in the proper, outer scope
	this.scope.Declare(name, typ)

	// now prepare the scope before parsing function body which will reference
	// parameters etc.

	oldFunction := this.currentFunction
	defer func() { this.currentFunction = oldFunction }()
	this.currentFunction = typ

	oldScope := this.scope
	defer func() { this.scope = oldScope }()
	this.scope = NewScope(this.scope)
	// add parameters as named variables
	for i := 0; i < len(paramNames); i++ {
		this.scope.Declare(paramNames[i], paramTypes[i])
	}

	return vm.NewDeclarationStatement(
		name, vm.NewLiteral(
			vm.NewFunction(typ, paramNames, this.visitScope(ctx.GetBody())),
		),
	)
}

type param struct {
	name string
	typ  vm.Type
}

func (this *MyVisitor) visitParams(ctx IParamDeclContext) (paramNames []string, paramTypes []vm.Type) {
	trace(ctx)

	groups := ctx.GetParamGroup()
	names := make([]string, 0, len(groups))
	types := make([]vm.Type, 0, len(groups))

	for _, child := range groups {
		newNames, newTypes := this.visitCombinedParam(child)
		for _, name := range newNames {
			if len(names) >= cap(names) {
				temp := make([]string, len(names), cap(names)+1)
				copy(temp, names)
				names = temp
			}
			names = append(names, name)
		}
		for _, typ := range newTypes {
			if len(types) >= cap(types) {
				temp := make([]vm.Type, len(types), cap(types)+1)
				copy(temp, types)
				types = temp
			}
			types = append(types, typ)
		}
	}

	return names, types
}

func (this *MyVisitor) visitCombinedParam(ctx ICombinedParamContext) (names []string, types []vm.Type) {
	trace(ctx)

	returnType := this.visitTypeName(ctx.GetParamType())

	names = make([]string, 0, len(ctx.GetNames()))
	types = make([]vm.Type, 0, len(ctx.GetNames()))
	for _, name := range ctx.GetNames() {
		names = append(names, name.GetText())
		types = append(types, returnType)
	}

	return names, types
}

func (this *MyVisitor) visitTypeName(ctx interface{}) vm.Type {
	switch ctx := ctx.(type) {
	case *TypeNameContext:
		return this.visitTypeName(ctx.GetChild(0))
	case *SimpleTypeContext:
		return this.visitSimpleType(ctx)
	case *FunctionTypeContext:
		return this.visitFunctionType(ctx)
	case *StructTypeContext:
		return this.visitStructType(ctx)
	default:
		panic("unknown type") // grammar should not allow this
	}
}

func (this *MyVisitor) visitSimpleType(ctx *SimpleTypeContext) vm.Type {
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

func (this *MyVisitor) visitFunctionType(ctx *FunctionTypeContext) vm.Type {
	trace(ctx)

	paramTypes := make([]vm.Type, 0, len(ctx.GetParamTypes()))
	for _, param := range ctx.GetParamTypes() {
		paramTypes = append(paramTypes, this.visitTypeName(param))
	}

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.visitTypeName(ctx.GetReturnType())
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	return vm.NewFunctionType(paramTypes, returnType)
}

func (this *MyVisitor) visitStructType(ctx *StructTypeContext) vm.Type {
	trace(ctx)

	fields := ctx.AllCombinedField()
	names := make([]string, 0, len(fields))
	types := make([]vm.Type, 0, len(fields))

	for _, child := range fields {
		newNames, newTypes := this.visitCombinedField(child)
		for _, name := range newNames {
			if len(names) >= cap(names) {
				temp := make([]string, len(names), cap(names)+1)
				copy(temp, names)
				names = temp
			}
			names = append(names, name)
		}
		for _, typ := range newTypes {
			if len(types) >= cap(types) {
				temp := make([]vm.Type, len(types), cap(types)+1)
				copy(temp, types)
				types = temp
			}
			types = append(types, typ)
		}
	}

	return vm.NewStructType(names, types)
}

func (this *MyVisitor) visitCombinedField(ctx ICombinedFieldContext) (names []string, types []vm.Type) {
	trace(ctx)

	returnType := this.visitTypeName(ctx.GetParamType())

	names = make([]string, 0, len(ctx.GetNames()))
	types = make([]vm.Type, 0, len(ctx.GetNames()))
	for _, name := range ctx.GetNames() {
		names = append(names, name.GetText())
		types = append(types, returnType)
	}

	return names, types
}

func (this *MyVisitor) visitScope(ctx IScopeContext) vm.ScopeStatement {
	trace(ctx)

	scope := vm.NewScopeStatement()
	for _, stmt := range ctx.GetStatements() {
		scope.AddStatement(this.visitStatement(stmt))
	}

	return scope
}

func (this *MyVisitor) visitReturnStatement(ctx *ReturnStatementContext) vm.Statement {
	trace(ctx)

	var retVal vm.Expression
	var retType vm.Type
	if ctx.GetRetVal() != nil {
		retVal = this.visitExpression(ctx.GetRetVal())
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

func (this *MyVisitor) visitIfStatement(ctx IIfStatementContext) vm.Statement {
	trace(ctx)

	conditions := ctx.GetConditions()
	effects := ctx.GetEffects()

	var retVal vm.IfStatement     // first if which will acutually be returned
	var previousIf vm.IfStatement // current if used to chain ifs

	for i, condCtx := range conditions {
		condition := this.visitExpression(condCtx)
		if condition.Type().GetKind() != vm.BOOLEAN {
			token := condCtx.GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				"'if' condition must be a boolean expression, but it is "+condition.Type().String())
		}

		effect := this.visitScope(effects[i])

		if retVal == nil {
			retVal = vm.NewIfStatement(condition, effect, nil)
			previousIf = retVal
		} else {
			// make new if the alternative of the previous one
			chainedIf := vm.NewIfStatement(condition, effect, nil)
			previousIf.SetAlternative(chainedIf)
			previousIf = chainedIf
		}
	}

	if ctx.GetAlternative() != nil {
		previousIf.SetAlternative(this.visitScope(ctx.GetAlternative()))
	}

	return retVal
}

func (this *MyVisitor) visitLoop(ctx *LoopContext) vm.Statement {
	trace(ctx)

	var label string
	// if ctx.GetLabel() != nil {
	// 	label = ctx.GetLabel().GetText()
	// }

	var condition vm.Expression
	// var isWhile bool
	if ctx.GetCondition() != nil {
		// isWhile = ctx.GetKind().GetText() == "while"
		//
		condition = this.visitExpression(ctx.GetCondition())
		if condition.Type().GetKind() != vm.BOOLEAN {
			token := ctx.GetCondition().GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				"if condition must be a boolean expression, but it is "+condition.Type().String())
		}
	}

	body := this.visitScope(ctx.GetBody())

	if condition == nil {
		return vm.NewLoopStatement(label, body)
	} else {
		return vm.NewWhileStatement(label, condition, body)
	}
	// } else if isWhile {
	// 	return vm.NewWhileStatement(label, condition, body)
	// } else {
	// 	return vm.NewUntilStatement(label, condition, body)
	// }
}

func (this *MyVisitor) visitLiteralPassthrough(ctx *LiteralPassthroughContext) vm.Expression {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *RealLiteralContext:
		return this.visitRealLiteral(child)
	case *IntegerLiteralContext:
		return this.visitIntegerLiteral(child)
	case *BooleanLiteralContext:
		return this.visitBooleanLiteral(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) visitAssignment(ctx *AssignmentContext) vm.Expression {
	trace(ctx)

	left := this.visitExpression(ctx.GetLeft())
	right := this.visitExpression(ctx.GetRight())

	if !left.Type().Equal(right.Type()) {
		token := ctx.GetRight().GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("type mismatch\n Left: %s\nRight: %s", left.Type(), right.Type()))
	}

	switch ctx.GetOp().GetText() {
	case "=":
		return vm.NewAssignment(left, right)
	case "+=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "+", right))
	case "-=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "-", right))
	case "*=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "*", right))
	case "/=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "/", right))
	case "%=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "%", right))
	case "&=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "&", right))
	case "|=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "|", right))
	case "^=":
		return vm.NewAssignment(left, vm.NewBinaryExpression(left, "^", right))
	default:
		panic("unknown operation")
	}
}

func (this *MyVisitor) visitBinaryExpression(ctx *BinaryExpressionContext) vm.Expression {
	trace(ctx)

	left := this.visitExpression(ctx.left)
	right := this.visitExpression(ctx.right)

	return vm.NewBinaryExpression(left, ctx.op.GetText(), right)
}

func (this *MyVisitor) visitRealLiteral(ctx *RealLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewRealLiteral(ctx.GetText())
}

func (this *MyVisitor) visitIntegerLiteral(ctx *IntegerLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewIntegerLiteral(ctx.GetText())
}

func (this *MyVisitor) visitBooleanLiteral(ctx *BooleanLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewBooleanLiteral(ctx.GetText())
}

func (this *MyVisitor) visitVariableExpression(ctx *VariableExpressionContext) vm.Expression {
	trace(ctx)

	token := ctx.GetName()
	name := token.GetText()
	typ := this.scope.Get(name)
	if typ == nil {
		this.errors.Error(token.GetLine(), token.GetColumn(), name+": undeclared variable")
		return nil
	} else {
		return vm.NewVariableReference(name, typ)
	}
}

func (this *MyVisitor) visitFunctionCall(ctx *FunctionCallContext) vm.Expression {
	trace(ctx)

	token := ctx.GetName()
	name := token.GetText()
	typ := this.scope.Get(name)
	var returnType vm.Type
	if typ == nil {
		this.errors.Error(token.GetLine(), token.GetColumn(), name+": undeclared function")
		returnType = vm.NewSimpleType(vm.VOID)
	} else {
		returnType = typ.GetResultType()
	}

	params := make([]vm.Expression, 0, len(ctx.GetParams()))
	paramTypes := make([]vm.Type, 0, len(ctx.GetParams()))
	for _, ectx := range ctx.GetParams() {
		expression := this.visitExpression(ectx)
		params = append(params, expression)
		paramTypes = append(paramTypes, expression.Type())
	}

	callType := vm.NewFunctionType(paramTypes, returnType)

	if typ != nil && !callType.Equal(typ) {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("function parameter mismatch\nExpected: %s\n     Got: %s", typ, callType))
	}

	return vm.NewFunctionCall(name, params, returnType)
}

func (this *MyVisitor) visitFieldSelector(ctx *FieldSelectorContext) vm.Expression {
	trace(ctx)

	base := this.visitExpression(ctx.GetBase())
	baseType := base.Type()

	field := ctx.GetField().GetText()
	index := baseType.GetFieldIndex(field)

	return vm.NewFieldSelector(base, index)
}

func (this *MyVisitor) visitLambda(ctx *LambdaContext) vm.Expression {
	trace(ctx)

	paramNames, paramTypes := this.visitParams(ctx.GetParams())
	// len(paramNames) == len(paramTypes)

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.visitTypeName(ctx.GetReturnType())
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	typ := vm.NewFunctionType(paramTypes, returnType)

	// now prepare the scope before parsing function body which will reference
	// parameters etc.

	oldFunction := this.currentFunction
	defer func() { this.currentFunction = oldFunction }()
	this.currentFunction = typ

	oldScope := this.scope
	defer func() { this.scope = oldScope }()
	this.scope = NewScope(this.scope)
	// add parameters as named variables
	for i := 0; i < len(paramNames); i++ {
		this.scope.Declare(paramNames[i], paramTypes[i])
	}

	return vm.NewLiteral(
		vm.NewFunction(typ, paramNames, this.visitScope(ctx.GetBody())),
	)
}
