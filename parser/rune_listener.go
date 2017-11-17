// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuneListener is a complete listener for a parse tree produced by RuneParser.
type RuneListener interface {
	antlr.ParseTreeListener

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArithExpr is called when entering the arithExpr production.
	EnterArithExpr(c *ArithExprContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArithExpr is called when exiting the arithExpr production.
	ExitArithExpr(c *ArithExprContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}
