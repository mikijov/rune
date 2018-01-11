// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import "github.com/antlr/antlr4/runtime/Go/antlr"

import (
	"mikijov/rune/vm"
)

var _ vm.Type // inhibit unused import error

// A complete Visitor for a parse tree produced by RuneParser.
type RuneVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RuneParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RuneParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RuneParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by RuneParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by RuneParser#SimpleType.
	VisitSimpleType(ctx *SimpleTypeContext) interface{}

	// Visit a parse tree produced by RuneParser#FunctionType.
	VisitFunctionType(ctx *FunctionTypeContext) interface{}

	// Visit a parse tree produced by RuneParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by RuneParser#paramDecl.
	VisitParamDecl(ctx *ParamDeclContext) interface{}

	// Visit a parse tree produced by RuneParser#combinedParam.
	VisitCombinedParam(ctx *CombinedParamContext) interface{}

	// Visit a parse tree produced by RuneParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by RuneParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by RuneParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by RuneParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by RuneParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#Assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by RuneParser#LiteralPassthrough.
	VisitLiteralPassthrough(ctx *LiteralPassthroughContext) interface{}

	// Visit a parse tree produced by RuneParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#ExpressionPassthrough.
	VisitExpressionPassthrough(ctx *ExpressionPassthroughContext) interface{}

	// Visit a parse tree produced by RuneParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#VariableExpression.
	VisitVariableExpression(ctx *VariableExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by RuneParser#Lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by RuneParser#assignmentOp.
	VisitAssignmentOp(ctx *AssignmentOpContext) interface{}

	// Visit a parse tree produced by RuneParser#RealLiteral.
	VisitRealLiteral(ctx *RealLiteralContext) interface{}

	// Visit a parse tree produced by RuneParser#IntegerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}

	// Visit a parse tree produced by RuneParser#BooleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}
}
