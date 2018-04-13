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
	// if ctx == nil {
	// 	fmt.Printf("nil\n")
	// }
	// fmt.Printf("%s(%s)\n", f.Name(), ctx.GetText())
	// if ctx != nil {
	// 	fmt.Printf("%T\n", ctx)
	// 	return
	// 	// for i, child := range ctx.GetChildren() {
	// 	// 	switch child := child.(type) {
	// 	// 	case antlr.ParserRuleContext:
	// 	// 		fmt.Printf("%d:%s\n", i, child.GetText())
	// 	// 	case *antlr.TerminalNodeImpl:
	// 	// 		fmt.Printf("%d:%s\n", i, child.GetText())
	// 	// 	default:
	// 	// 		fmt.Printf("Unknown\n")
	// 	// 	}
	// 	// }
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

	savedErrors := this.errors // preserve current error state

	statements := make([]IStatementContext, len(ctx.GetStatements()))
	copy(statements, ctx.GetStatements())
	for {
		this.errors = NewRuneErrorListener()

		foundAnyStatements := false
		for i, stmtCtx := range statements {
			if statements[i] != nil {
				if done := this.visitTopLevelDeclaration(stmtCtx); done {
					statements[i] = nil
					foundAnyStatements = true
				}
			}
		}

		if !foundAnyStatements {
			break
		}
	}

	// repeat once more to record all errors
	this.errors = savedErrors
	hasErrors := false
	for i, stmtCtx := range statements {
		if statements[i] != nil {
			hasErrors = true // any statement in this state must be an error
			if done := this.visitTopLevelDeclaration(stmtCtx); done {
				panic("this should never happen")
			}
		}
	}

	if !hasErrors {
		for _, stmtCtx := range ctx.GetStatements() {
			if stmt := this.visitTopLevelImplementation(stmtCtx); stmt != nil {
				this.program.AddStatement(stmt)
			}
		}
	}

	return this.program
}

func (this *MyVisitor) visitTopLevelDeclaration(ctx IStatementContext) bool {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *DeclarationContext:
		return this.visitDeclaration1(child)
	case *TypeDeclarationContext:
		return this.visitTypeDeclaration1(child)
	case *FunctionContext:
		return this.visitFunction1(child)
	default:
		return true
	}
}

func (this *MyVisitor) visitTopLevelImplementation(ctx IStatementContext) vm.Statement {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *DeclarationContext:
		return this.visitDeclaration2(child)
	case *TypeDeclarationContext:
		return this.visitTypeDeclaration2(child)
	case *FunctionContext:
		return this.visitFunction2(child)
	case *ExpressionContext:
		return vm.NewExpressionStatement(this.visitExpression(child))
	case *LoopContext:
		return this.visitLoop(child)
	case *ReturnStatementContext:
		return this.visitReturnStatement(child)
	case *IfStatementContext:
		return this.visitIfStatement(child)
	default:
		panic("unexpected top level construct")
	}
}

func (this *MyVisitor) visitStatement(ctx IStatementContext) vm.Statement {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *ExpressionContext:
		return vm.NewExpressionStatement(this.visitExpression(child))
	case *DeclarationContext:
		return this.visitDeclaration(child)
	case *TypeDeclarationContext:
		return this.visitTypeDeclaration(child)
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
	case *ArraySelectorContext:
		return this.visitArraySelector(ctx)
	case *LambdaContext:
		return this.visitLambda(ctx)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx))
	}
}

func (this *MyVisitor) visitTypeDeclaration1(ctx *TypeDeclarationContext) bool {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()

	typ := this.visitTypeName(ctx.GetType_())
	if typ == nil {
		return false
	}

	if !this.scope.Declare(name, typ) {
		token := ctx.GetIdentifier()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("type redeclared '%s'", name))
		return false
	}
	return true
}

func (this *MyVisitor) visitTypeDeclaration2(ctx *TypeDeclarationContext) vm.Statement {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()
	typ := this.scope.Get(name)
	// typ != nil

	return vm.NewTypeDeclarationStatement(name, typ)
}

func (this *MyVisitor) visitTypeDeclaration(ctx *TypeDeclarationContext) vm.Statement {
	if !this.visitTypeDeclaration1(ctx) {
		return nil
	}
	return this.visitTypeDeclaration2(ctx)
}

func (this *MyVisitor) visitDeclaration1(ctx *DeclarationContext) bool {
	trace(ctx)

	var typ vm.Type
	if ctx.GetType_() != nil {
		typ = this.visitTypeName(ctx.GetType_())
		if typ == nil {
			return false
		}
	}

	var value vm.Expression
	if ctx.GetValue() != nil {
		value = this.visitExpression(ctx.GetValue())
	}

	if typ != nil && value != nil {
		if !typ.Equal(value.Type()) {
			token := ctx.GetIdentifier()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				fmt.Sprintf("type mismatch %s != %s", typ, value.Type()))
			return false
		}
	} else if typ != nil {
		// already have type, nothing to do
	} else if value != nil {
		typ = value.Type()
	} else {
		// normally grammar should not allow this, probably caused by incomplete top level declaration
		// but at top level also possible due to incomplete types, declare error just in case
		token := ctx.GetIdentifier()
		this.errors.Error(token.GetLine(), token.GetColumn(), "type not provided")
		return false
	}

	name := ctx.GetIdentifier().GetText()
	if !this.scope.Declare(name, typ) {
		token := ctx.GetIdentifier()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("'%s' redeclared", name))
		return false
	}

	return true
}

func (this *MyVisitor) visitDeclaration2(ctx *DeclarationContext) vm.Statement {
	trace(ctx)

	var typ vm.Type
	if ctx.GetType_() != nil {
		typ = this.visitTypeName(ctx.GetType_())
		if typ == nil {
			return nil
		}
	}

	var value vm.Expression
	if ctx.GetValue() != nil {
		value = this.visitExpression(ctx.GetValue())
		if value == nil {
			return nil
		}
	}

	if typ != nil && value != nil {
		if !typ.Equal(value.Type()) {
			token := ctx.GetIdentifier()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				fmt.Sprintf("type mismatch %s != %s", typ, value.Type()))
			return nil
		}
	} else if typ != nil {
		// already have type, nothing to do
	} else if value != nil {
		typ = value.Type()
	} else {
		panic("declaration should have taken care of this")
	}

	name := ctx.GetIdentifier().GetText()
	if value != nil {
		return vm.NewDeclarationStatement(name, value)
	}
	return vm.NewDeclarationStatement(name, vm.NewLiteral(typ.GetZero()))
}

func (this *MyVisitor) visitDeclaration(ctx *DeclarationContext) vm.Statement {
	if !this.visitDeclaration1(ctx) {
		return nil
	}
	return this.visitDeclaration2(ctx)
}

func (this *MyVisitor) visitFunction1(ctx *FunctionContext) bool {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()

	paramNames, paramTypes := this.visitParams(ctx.GetParams())
	if paramNames == nil {
		return false
	}
	// len(paramNames) == len(paramTypes)

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.visitTypeName(ctx.GetReturnType())
		if returnType == nil {
			return false
		}
	} else {
		returnType = vm.NewSimpleType(vm.VOID)
	}

	typ := vm.NewFunctionType(paramTypes, returnType)
	if !this.scope.Declare(name, typ) {
		token := ctx.GetIdentifier()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("'%s' redeclared", name))
	}
	return true
}

func (this *MyVisitor) visitFunction2(ctx *FunctionContext) vm.Statement {
	trace(ctx)

	name := ctx.GetIdentifier().GetText()
	paramNames, paramTypes := this.visitParams(ctx.GetParams())
	typ := this.scope.Get(name)
	// assert typ != nil
	// typ.params == paramTypes

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

	scope := this.visitScope(ctx.GetBody())
	if scope == nil {
		return nil
	}

	return vm.NewDeclarationStatement(
		name, vm.NewLiteral(
			vm.NewFunction(typ, paramNames, scope),
		),
	)
}

func (this *MyVisitor) visitFunction(ctx *FunctionContext) vm.Statement {
	if !this.visitFunction1(ctx) {
		return nil
	}
	return this.visitFunction2(ctx)
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
		if newNames == nil {
			return nil, nil
		}
		for _, name := range newNames {
			names = append(names, name)
		}
		for _, typ := range newTypes {
			types = append(types, typ)
		}
	}

	return names, types
}

func (this *MyVisitor) visitCombinedParam(ctx ICombinedParamContext) (names []string, types []vm.Type) {
	trace(ctx)

	returnType := this.visitTypeName(ctx.GetParamType())
	if returnType == nil {
		return nil, nil
	}

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
		typ := this.visitTypeName(ctx.GetChild(0))
		if typ != nil && ctx.GetIsarray() != nil {
			typ = vm.NewArrayType(typ)
		}
		return typ
	case *SimpleTypeContext:
		return this.visitSimpleType(ctx)
	case *FunctionTypeContext:
		return this.visitFunctionType(ctx)
	case *StructTypeContext:
		return this.visitStructType(ctx)
	case *CustomTypeContext:
		return this.visitCustomType(ctx)
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

func (this *MyVisitor) visitCustomType(ctx *CustomTypeContext) vm.Type {
	trace(ctx)

	name := ctx.GetText()
	typ := this.scope.Get(name)
	if typ == nil {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("unknown type '%s'", name))
		return nil
	}
	return typ
}

func (this *MyVisitor) visitFunctionType(ctx *FunctionTypeContext) vm.Type {
	trace(ctx)

	paramTypes := make([]vm.Type, 0, len(ctx.GetParamTypes()))
	for _, param := range ctx.GetParamTypes() {
		typ := this.visitTypeName(param)
		if typ == nil {
			return nil
		}
		paramTypes = append(paramTypes, typ)
	}

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.visitTypeName(ctx.GetReturnType())
		if returnType == nil {
			return nil
		}
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
		if newNames == nil {
			return nil
		}
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
	if returnType == nil {
		return nil, nil
	}

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
	for _, stmtCtx := range ctx.GetStatements() {
		scope.AddStatement(this.visitStatement(stmtCtx))
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
	if left == nil {
		return nil
	}
	right := this.visitExpression(ctx.GetRight())
	if right == nil {
		return nil
	}

	if !left.Type().Equal(right.Type()) {
		token := ctx.GetRight().GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("type mismatch\n Left: %s\nRight: %s", left.Type(), right.Type()))
		return nil
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
	if left == nil {
		return nil
	}
	right := this.visitExpression(ctx.right)
	if right == nil {
		return nil
	}

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
	}
	return vm.NewVariableReference(name, typ)
}

func (this *MyVisitor) visitFunctionCall(ctx *FunctionCallContext) vm.Expression {
	trace(ctx)

	token := ctx.GetName()
	name := token.GetText()
	typ := this.scope.Get(name)
	var returnType vm.Type
	if typ == nil {
		this.errors.Error(token.GetLine(), token.GetColumn(), name+": undeclared function")
		// returnType = vm.NewSimpleType(vm.VOID)
		return nil
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
		return nil
	}

	return vm.NewFunctionCall(name, params, returnType)
}

func (this *MyVisitor) visitFieldSelector(ctx *FieldSelectorContext) vm.Expression {
	trace(ctx)

	base := this.visitExpression(ctx.GetBase())
	if base == nil {
		return nil
	}
	baseType := base.Type()

	field := ctx.GetField().GetText()
	index := baseType.GetFieldIndex(field)
	if index < 0 {
		token := ctx.GetField()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("unknown field: %s", field))
		return nil
	}

	return vm.NewFieldSelector(base, index)
}

func (this *MyVisitor) visitArraySelector(ctx *ArraySelectorContext) vm.Expression {
	trace(ctx)

	array := this.visitExpression(ctx.GetArray())
	if array == nil {
		return nil
	}
	if array.Type().GetKind() != vm.ARRAY {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(), "not an array")
		return nil
	}
	index := this.visitExpression(ctx.GetIndex())
	if index == nil {
		return nil
	}
	if !index.Type().GetResultType().Equal(vm.NewSimpleType(vm.INTEGER)) {
		token := ctx.GetStart()
		this.errors.Error(token.GetLine(), token.GetColumn(),
			fmt.Sprintf("index is not an integer but %v", index.Type().GetResultType()))
		return nil
	}

	return vm.NewArraySelector(array, index)
}

func (this *MyVisitor) visitLambda(ctx *LambdaContext) vm.Expression {
	trace(ctx)

	paramNames, paramTypes := this.visitParams(ctx.GetParams())
	if paramNames == nil {
		return nil
	}
	// len(paramNames) == len(paramTypes)

	var returnType vm.Type
	if ctx.GetReturnType() != nil {
		returnType = this.visitTypeName(ctx.GetReturnType())
		if returnType == nil {
			return nil
		}
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
		if ok := this.scope.Declare(paramNames[i], paramTypes[i]); !ok {
			token := ctx.GetStart()
			this.errors.Error(token.GetLine(), token.GetColumn(),
				fmt.Sprintf("duplicate parameter name %v", paramNames[i]))
			return nil
		}
	}

	return vm.NewLiteral(
		vm.NewFunction(typ, paramNames, this.visitScope(ctx.GetBody())),
	)
}
