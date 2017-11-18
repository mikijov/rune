// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 70, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 45, 10, 6, 12, 6, 14,
	6, 48, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7,
	7, 59, 10, 7, 12, 7, 14, 7, 62, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 68,
	10, 8, 3, 8, 2, 4, 10, 12, 9, 2, 4, 6, 8, 10, 12, 14, 2, 4, 3, 2, 4, 5,
	3, 2, 6, 8, 2, 66, 2, 19, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 27, 3, 2, 2,
	2, 8, 32, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 49, 3, 2, 2, 2, 14, 67, 3,
	2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19,
	17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 22, 3, 2, 2, 2, 21, 19, 3, 2, 2,
	2, 22, 23, 7, 2, 2, 3, 23, 3, 3, 2, 2, 2, 24, 25, 5, 6, 4, 2, 25, 26, 8,
	3, 1, 2, 26, 5, 3, 2, 2, 2, 27, 28, 5, 8, 5, 2, 28, 29, 8, 4, 1, 2, 29,
	30, 3, 2, 2, 2, 30, 31, 7, 3, 2, 2, 31, 7, 3, 2, 2, 2, 32, 33, 5, 10, 6,
	2, 33, 34, 8, 5, 1, 2, 34, 9, 3, 2, 2, 2, 35, 36, 8, 6, 1, 2, 36, 37, 5,
	12, 7, 2, 37, 38, 8, 6, 1, 2, 38, 46, 3, 2, 2, 2, 39, 40, 12, 3, 2, 2,
	40, 41, 9, 2, 2, 2, 41, 42, 5, 12, 7, 2, 42, 43, 8, 6, 1, 2, 43, 45, 3,
	2, 2, 2, 44, 39, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46,
	47, 3, 2, 2, 2, 47, 11, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 8, 7, 1,
	2, 50, 51, 5, 14, 8, 2, 51, 52, 8, 7, 1, 2, 52, 60, 3, 2, 2, 2, 53, 54,
	12, 3, 2, 2, 54, 55, 9, 3, 2, 2, 55, 56, 5, 14, 8, 2, 56, 57, 8, 7, 1,
	2, 57, 59, 3, 2, 2, 2, 58, 53, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58,
	3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 13, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2,
	63, 64, 7, 9, 2, 2, 64, 68, 8, 8, 1, 2, 65, 66, 7, 10, 2, 2, 66, 68, 8,
	8, 1, 2, 67, 63, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 15, 3, 2, 2, 2, 6,
	19, 46, 60, 67,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'+'", "'-'", "'*'", "'/'", "'%'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "INTEGER_LITERAL", "REAL_LITERAL", "LINENDING",
	"WHITESPACE", "COMMENT",
}

var ruleNames = []string{
	"module", "stmt", "simpleStmt", "expression", "arithExpr", "term", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RuneParser struct {
	*antlr.BaseParser
}

func NewRuneParser(input antlr.TokenStream) *RuneParser {
	this := new(RuneParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rune.g4"

	return this
}

// RuneParser tokens.
const (
	RuneParserEOF             = antlr.TokenEOF
	RuneParserT__0            = 1
	RuneParserT__1            = 2
	RuneParserT__2            = 3
	RuneParserT__3            = 4
	RuneParserT__4            = 5
	RuneParserT__5            = 6
	RuneParserINTEGER_LITERAL = 7
	RuneParserREAL_LITERAL    = 8
	RuneParserLINENDING       = 9
	RuneParserWHITESPACE      = 10
	RuneParserCOMMENT         = 11
)

// RuneParser rules.
const (
	RuneParserRULE_module     = 0
	RuneParserRULE_stmt       = 1
	RuneParserRULE_simpleStmt = 2
	RuneParserRULE_expression = 3
	RuneParserRULE_arithExpr  = 4
	RuneParserRULE_term       = 5
	RuneParserRULE_atom       = 6
)

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m attribute.
	GetM() Module

	// SetM sets the m attribute.
	SetM(Module)

	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	m      Module
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) GetM() Module { return s.m }

func (s *ModuleContext) SetM(v Module) { s.m = v }

func (s *ModuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuneParserEOF, 0)
}

func (s *ModuleContext) AllStmt() []IStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStmtContext)(nil)).Elem())
	var tst = make([]IStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStmtContext)
		}
	}

	return tst
}

func (s *ModuleContext) Stmt(i int) IStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitModule(s)
	}
}

func (p *RuneParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuneParserRULE_module)

	localctx.(*ModuleContext).m = NewModule()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuneParserINTEGER_LITERAL || _la == RuneParserREAL_LITERAL {
		{
			p.SetState(14)
			p.Stmt()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(20)
		p.Match(RuneParserEOF)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))

	for _, stmt := range localctx.(*ModuleContext).AllStmt() {
		localctx.(*ModuleContext).m.AddStatement(stmt.GetS())
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_simpleStmt returns the _simpleStmt rule contexts.
	Get_simpleStmt() ISimpleStmtContext

	// Set_simpleStmt sets the _simpleStmt rule contexts.
	Set_simpleStmt(ISimpleStmtContext)

	// GetS returns the s attribute.
	GetS() Statement

	// SetS sets the s attribute.
	SetS(Statement)

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	s           Statement
	_simpleStmt ISimpleStmtContext
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Get_simpleStmt() ISimpleStmtContext { return s._simpleStmt }

func (s *StmtContext) Set_simpleStmt(v ISimpleStmtContext) { s._simpleStmt = v }

func (s *StmtContext) GetS() Statement { return s.s }

func (s *StmtContext) SetS(v Statement) { s.s = v }

func (s *StmtContext) SimpleStmt() ISimpleStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *RuneParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuneParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)

		var _x = p.SimpleStmt()

		localctx.(*StmtContext)._simpleStmt = _x
	}
	localctx.(*StmtContext).s = localctx.(*StmtContext).Get_simpleStmt().GetS()

	return localctx
}

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetS returns the s attribute.
	GetS() Statement

	// SetS sets the s attribute.
	SetS(Statement)

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	s           Statement
	_expression IExpressionContext
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_simpleStmt
	return p
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Get_expression() IExpressionContext { return s._expression }

func (s *SimpleStmtContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *SimpleStmtContext) GetS() Statement { return s.s }

func (s *SimpleStmtContext) SetS(v Statement) { s.s = v }

func (s *SimpleStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (p *RuneParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuneParserRULE_simpleStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)

		var _x = p.Expression()

		localctx.(*SimpleStmtContext)._expression = _x
	}
	localctx.(*SimpleStmtContext).s = &expressionStatement{localctx.(*SimpleStmtContext).Get_expression().GetExpr()}

	{
		p.SetState(28)
		p.Match(RuneParserT__0)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IArithExprContext

	// SetE sets the e rule contexts.
	SetE(IArithExprContext)

	// GetExpr returns the expr attribute.
	GetExpr() Expression

	// SetExpr sets the expr attribute.
	SetExpr(Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   Expression
	e      IArithExprContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetE() IArithExprContext { return s.e }

func (s *ExpressionContext) SetE(v IArithExprContext) { s.e = v }

func (s *ExpressionContext) GetExpr() Expression { return s.expr }

func (s *ExpressionContext) SetExpr(v Expression) { s.expr = v }

func (s *ExpressionContext) ArithExpr() IArithExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithExprContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RuneParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RuneParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)

		var _x = p.arithExpr(0)

		localctx.(*ExpressionContext).e = _x
	}
	localctx.(*ExpressionContext).expr = localctx.(*ExpressionContext).GetE().GetExpr()

	return localctx
}

// IArithExprContext is an interface to support dynamic dispatch.
type IArithExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IArithExprContext

	// GetE returns the e rule contexts.
	GetE() ITermContext

	// GetRight returns the right rule contexts.
	GetRight() ITermContext

	// SetLeft sets the left rule contexts.
	SetLeft(IArithExprContext)

	// SetE sets the e rule contexts.
	SetE(ITermContext)

	// SetRight sets the right rule contexts.
	SetRight(ITermContext)

	// GetExpr returns the expr attribute.
	GetExpr() Expression

	// SetExpr sets the expr attribute.
	SetExpr(Expression)

	// IsArithExprContext differentiates from other interfaces.
	IsArithExprContext()
}

type ArithExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   Expression
	left   IArithExprContext
	e      ITermContext
	op     antlr.Token
	right  ITermContext
}

func NewEmptyArithExprContext() *ArithExprContext {
	var p = new(ArithExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_arithExpr
	return p
}

func (*ArithExprContext) IsArithExprContext() {}

func NewArithExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithExprContext {
	var p = new(ArithExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_arithExpr

	return p
}

func (s *ArithExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithExprContext) GetOp() antlr.Token { return s.op }

func (s *ArithExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ArithExprContext) GetLeft() IArithExprContext { return s.left }

func (s *ArithExprContext) GetE() ITermContext { return s.e }

func (s *ArithExprContext) GetRight() ITermContext { return s.right }

func (s *ArithExprContext) SetLeft(v IArithExprContext) { s.left = v }

func (s *ArithExprContext) SetE(v ITermContext) { s.e = v }

func (s *ArithExprContext) SetRight(v ITermContext) { s.right = v }

func (s *ArithExprContext) GetExpr() Expression { return s.expr }

func (s *ArithExprContext) SetExpr(v Expression) { s.expr = v }

func (s *ArithExprContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ArithExprContext) ArithExpr() IArithExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithExprContext)
}

func (s *ArithExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterArithExpr(s)
	}
}

func (s *ArithExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitArithExpr(s)
	}
}

func (p *RuneParser) ArithExpr() (localctx IArithExprContext) {
	return p.arithExpr(0)
}

func (p *RuneParser) arithExpr(_p int) (localctx IArithExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArithExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, RuneParserRULE_arithExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)

		var _x = p.term(0)

		localctx.(*ArithExprContext).e = _x
	}
	localctx.(*ArithExprContext).expr = localctx.(*ArithExprContext).GetE().GetExpr()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArithExprContext(p, _parentctx, _parentState)
			localctx.(*ArithExprContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_arithExpr)
			p.SetState(37)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(38)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ArithExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == RuneParserT__1 || _la == RuneParserT__2) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ArithExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(39)

				var _x = p.term(0)

				localctx.(*ArithExprContext).right = _x
			}
			localctx.(*ArithExprContext).expr = NewBinaryExpression(localctx.(*ArithExprContext).GetLeft().GetExpr(), (func() string {
				if localctx.(*ArithExprContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*ArithExprContext).GetOp().GetText()
				}
			}()), localctx.(*ArithExprContext).GetRight().GetExpr())

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() ITermContext

	// GetE returns the e rule contexts.
	GetE() IAtomContext

	// GetRight returns the right rule contexts.
	GetRight() IAtomContext

	// SetLeft sets the left rule contexts.
	SetLeft(ITermContext)

	// SetE sets the e rule contexts.
	SetE(IAtomContext)

	// SetRight sets the right rule contexts.
	SetRight(IAtomContext)

	// GetExpr returns the expr attribute.
	GetExpr() Expression

	// SetExpr sets the expr attribute.
	SetExpr(Expression)

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   Expression
	left   ITermContext
	e      IAtomContext
	op     antlr.Token
	right  IAtomContext
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) GetOp() antlr.Token { return s.op }

func (s *TermContext) SetOp(v antlr.Token) { s.op = v }

func (s *TermContext) GetLeft() ITermContext { return s.left }

func (s *TermContext) GetE() IAtomContext { return s.e }

func (s *TermContext) GetRight() IAtomContext { return s.right }

func (s *TermContext) SetLeft(v ITermContext) { s.left = v }

func (s *TermContext) SetE(v IAtomContext) { s.e = v }

func (s *TermContext) SetRight(v IAtomContext) { s.right = v }

func (s *TermContext) GetExpr() Expression { return s.expr }

func (s *TermContext) SetExpr(v Expression) { s.expr = v }

func (s *TermContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *RuneParser) Term() (localctx ITermContext) {
	return p.term(0)
}

func (p *RuneParser) term(_p int) (localctx ITermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, RuneParserRULE_term, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)

		var _x = p.Atom()

		localctx.(*TermContext).e = _x
	}
	localctx.(*TermContext).expr = localctx.(*TermContext).GetE().GetExpr()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTermContext(p, _parentctx, _parentState)
			localctx.(*TermContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_term)
			p.SetState(51)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(52)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*TermContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__3)|(1<<RuneParserT__4)|(1<<RuneParserT__5))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*TermContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(53)

				var _x = p.Atom()

				localctx.(*TermContext).right = _x
			}
			localctx.(*TermContext).expr = NewBinaryExpression(localctx.(*TermContext).GetLeft().GetExpr(), (func() string {
				if localctx.(*TermContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*TermContext).GetOp().GetText()
				}
			}()), localctx.(*TermContext).GetRight().GetExpr())

		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// GetExpr returns the expr attribute.
	GetExpr() Expression

	// SetExpr sets the expr attribute.
	SetExpr(Expression)

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   Expression
	val    antlr.Token
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) GetVal() antlr.Token { return s.val }

func (s *AtomContext) SetVal(v antlr.Token) { s.val = v }

func (s *AtomContext) GetExpr() Expression { return s.expr }

func (s *AtomContext) SetExpr(v Expression) { s.expr = v }

func (s *AtomContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(RuneParserINTEGER_LITERAL, 0)
}

func (s *AtomContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(RuneParserREAL_LITERAL, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuneListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *RuneParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuneParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuneParserINTEGER_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)

			var _m = p.Match(RuneParserINTEGER_LITERAL)

			localctx.(*AtomContext).val = _m
		}
		localctx.(*AtomContext).expr = NewIntegerLiteral((func() string {
			if localctx.(*AtomContext).GetVal() == nil {
				return ""
			} else {
				return localctx.(*AtomContext).GetVal().GetText()
			}
		}()))

	case RuneParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)

			var _m = p.Match(RuneParserREAL_LITERAL)

			localctx.(*AtomContext).val = _m
		}
		localctx.(*AtomContext).expr = NewRealLiteral((func() string {
			if localctx.(*AtomContext).GetVal() == nil {
				return ""
			} else {
				return localctx.(*AtomContext).GetVal().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *RuneParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ArithExprContext = nil
		if localctx != nil {
			t = localctx.(*ArithExprContext)
		}
		return p.ArithExpr_Sempred(t, predIndex)

	case 5:
		var t *TermContext = nil
		if localctx != nil {
			t = localctx.(*TermContext)
		}
		return p.Term_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuneParser) ArithExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RuneParser) Term_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
