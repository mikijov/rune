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

// EnterModule is called when production module is entered.
func (s *BaseRuneListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseRuneListener) ExitModule(ctx *ModuleContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseRuneListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseRuneListener) ExitStmt(ctx *StmtContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseRuneListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseRuneListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRuneListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRuneListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArithExpr is called when production arithExpr is entered.
func (s *BaseRuneListener) EnterArithExpr(ctx *ArithExprContext) {}

// ExitArithExpr is called when production arithExpr is exited.
func (s *BaseRuneListener) ExitArithExpr(ctx *ArithExprContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseRuneListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseRuneListener) ExitTerm(ctx *TermContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseRuneListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseRuneListener) ExitAtom(ctx *AtomContext) {}
