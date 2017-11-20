// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import "github.com/antlr/antlr4/runtime/Go/antlr"

import (
	"mikijov/rune-antlr/vm"
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

	// Visit a parse tree produced by RuneParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#LiteralPassthrough.
	VisitLiteralPassthrough(ctx *LiteralPassthroughContext) interface{}

	// Visit a parse tree produced by RuneParser#BinaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#ExpressionPassthrough.
	VisitExpressionPassthrough(ctx *ExpressionPassthroughContext) interface{}

	// Visit a parse tree produced by RuneParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by RuneParser#RealLiteral.
	VisitRealLiteral(ctx *RealLiteralContext) interface{}

	// Visit a parse tree produced by RuneParser#IntegerLiteral.
	VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{}
}
