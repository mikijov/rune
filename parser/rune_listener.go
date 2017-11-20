// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuneListener is a complete listener for a parse tree produced by RuneParser.
type RuneListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLiteralPassthrough is called when entering the LiteralPassthrough production.
	EnterLiteralPassthrough(c *LiteralPassthroughContext)

	// EnterBinaryExpression is called when entering the BinaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterExpressionPassthrough is called when entering the ExpressionPassthrough production.
	EnterExpressionPassthrough(c *ExpressionPassthroughContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterRealLiteral is called when entering the RealLiteral production.
	EnterRealLiteral(c *RealLiteralContext)

	// EnterIntegerLiteral is called when entering the IntegerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLiteralPassthrough is called when exiting the LiteralPassthrough production.
	ExitLiteralPassthrough(c *LiteralPassthroughContext)

	// ExitBinaryExpression is called when exiting the BinaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitExpressionPassthrough is called when exiting the ExpressionPassthrough production.
	ExitExpressionPassthrough(c *ExpressionPassthroughContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitRealLiteral is called when exiting the RealLiteral production.
	ExitRealLiteral(c *RealLiteralContext)

	// ExitIntegerLiteral is called when exiting the IntegerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)
}
