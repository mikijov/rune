// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRuneVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRuneVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitParamDecl(ctx *ParamDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitCombinedParam(ctx *CombinedParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitLiteralPassthrough(ctx *LiteralPassthroughContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitExpressionPassthrough(ctx *ExpressionPassthroughContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitVariableExpression(ctx *VariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitRealLiteral(ctx *RealLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
