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

func (v *BaseRuneVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitSimpleType(ctx *SimpleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitFunctionType(ctx *FunctionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitCustomType(ctx *CustomTypeContext) interface{} {
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

func (v *BaseRuneVisitor) VisitCombinedField(ctx *CombinedFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitLoop(ctx *LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitLiteralPassthrough(ctx *LiteralPassthroughContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
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

func (v *BaseRuneVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitFieldSelector(ctx *FieldSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitArraySelector(ctx *ArraySelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitMethodCall(ctx *MethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitUnaryOp(ctx *UnaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitAssignmentOp(ctx *AssignmentOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitRealLiteral(ctx *RealLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRuneVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
