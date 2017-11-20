// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRuneListener is a complete listener for a parse tree produced by RuneParser.
type BaseRuneListener struct{}

var _ RuneListener = &BaseRuneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRuneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRuneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRuneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRuneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRuneListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRuneListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRuneListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRuneListener) ExitStatement(ctx *StatementContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseRuneListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseRuneListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseRuneListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseRuneListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRuneListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRuneListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLiteralPassthrough is called when production LiteralPassthrough is entered.
func (s *BaseRuneListener) EnterLiteralPassthrough(ctx *LiteralPassthroughContext) {}

// ExitLiteralPassthrough is called when production LiteralPassthrough is exited.
func (s *BaseRuneListener) ExitLiteralPassthrough(ctx *LiteralPassthroughContext) {}

// EnterBinaryExpression is called when production BinaryExpression is entered.
func (s *BaseRuneListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production BinaryExpression is exited.
func (s *BaseRuneListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterExpressionPassthrough is called when production ExpressionPassthrough is entered.
func (s *BaseRuneListener) EnterExpressionPassthrough(ctx *ExpressionPassthroughContext) {}

// ExitExpressionPassthrough is called when production ExpressionPassthrough is exited.
func (s *BaseRuneListener) ExitExpressionPassthrough(ctx *ExpressionPassthroughContext) {}

// EnterUnaryExpression is called when production UnaryExpression is entered.
func (s *BaseRuneListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production UnaryExpression is exited.
func (s *BaseRuneListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterRealLiteral is called when production RealLiteral is entered.
func (s *BaseRuneListener) EnterRealLiteral(ctx *RealLiteralContext) {}

// ExitRealLiteral is called when production RealLiteral is exited.
func (s *BaseRuneListener) ExitRealLiteral(ctx *RealLiteralContext) {}

// EnterIntegerLiteral is called when production IntegerLiteral is entered.
func (s *BaseRuneListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {}

// ExitIntegerLiteral is called when production IntegerLiteral is exited.
func (s *BaseRuneListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {}
