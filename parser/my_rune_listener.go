// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

// import (
// 	"fmt"
// 	"runtime"
//
// 	"github.com/antlr/antlr4/runtime/Go/antlr"
//
// 	"mikijov/rune-antlr/vm"
// )
//
// func trace() {
// 	pc := make([]uintptr, 10) // at least 1 entry needed
// 	runtime.Callers(2, pc)
// 	f := runtime.FuncForPC(pc[0])
// 	file, line := f.FileLine(pc[0])
// 	fmt.Printf("%s:%d %s\n", file, line, f.Name())
// }
//
// // MyRuneListener is a complete listener for a parse tree produced by RuneParser.
// type MyRuneListener struct {
// 	scope *vm.Scope
// }
//
// // var _ RuneListener = &MyRuneListener{}
//
// func NewMyRuneListener() RuneListener {
// 	return &MyRuneListener{
// 		scope: vm.NewScope(nil),
// 	}
// }
//
// func (s *MyRuneListener) VisitTerminal(node antlr.TerminalNode) {
// 	trace()
// }
//
// func (s *MyRuneListener) VisitErrorNode(node antlr.ErrorNode) {
// 	panic(node.GetText())
// }
//
// func (s *MyRuneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) EnterProgram(ctx *ProgramContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitProgram(ctx *ProgramContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) EnterStatement(ctx *StatementContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitStatement(ctx *StatementContext) {
// 	s.scope.AddStatement(ctx.s)
// }
//
// func (s *MyRuneListener) EnterDeclaration(ctx *DeclarationContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitDeclaration(ctx *DeclarationContext) {
// 	var type_ vm.Type
// 	if ctx.GetType_() != nil {
// 		type_ = ctx.GetType_().GetType_()
// 		if ctx.GetValue() != nil {
// 			if type_ != ctx.GetValue().GetExpr().Type() {
// 				panic("type mismatch")
// 			}
// 		}
// 	} else {
// 		if ctx.GetValue() == nil {
// 			panic("variable type not provided") // grammar should not allow this
// 		}
// 		type_ = ctx.GetValue().GetExpr().Type()
// 	}
//
// 	_, ok := s.scope.Declare(ctx.GetIdentifier().GetText(), type_)
// 	if !ok {
// 		panic("variable redeclared")
// 	}
// }
//
// func (s *MyRuneListener) EnterTypeName(ctx *TypeNameContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitTypeName(ctx *TypeNameContext) {
// 	switch ctx.GetText() {
// 	case "int":
// 		ctx.SetType_(vm.INTEGER)
// 	case "real":
// 		ctx.SetType_(vm.REAL)
// 	case "bool":
// 		ctx.SetType_(vm.BOOLEAN)
// 	default:
// 		panic("unknown type")
// 	}
// }
//
// func (s *MyRuneListener) EnterRealLiteral(ctx *RealLiteralContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitRealLiteral(ctx *RealLiteralContext) {
// 	ctx.SetExpr(vm.NewRealLiteral(ctx.GetText()))
// }
//
// func (s *MyRuneListener) EnterIntegerLiteral(ctx *IntegerLiteralContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitIntegerLiteral(ctx *IntegerLiteralContext) {
// 	ctx.SetExpr(vm.NewIntegerLiteral(ctx.GetText()))
// }
//
// func (s *MyRuneListener) EnterLiteralPassthrough(ctx *LiteralPassthroughContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitLiteralPassthrough(ctx *LiteralPassthroughContext) {
// 	ctx.SetExpr(ctx.GetValue().GetExpr())
// }
//
// func (s *MyRuneListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {
// 	ctx.SetExpr(vm.NewBinaryExpression(ctx.GetLeft().GetExpr(), ctx.op.GetText(), ctx.GetRight().GetExpr()))
// }
//
// func (s *MyRuneListener) EnterExpressionPassthrough(ctx *ExpressionPassthroughContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitExpressionPassthrough(ctx *ExpressionPassthroughContext) {
// 	trace()
// 	ctx.SetExpr(ctx.GetValue().GetExpr())
// }
//
// func (s *MyRuneListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {
// 	trace()
// }
//
// func (s *MyRuneListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {
// 	trace()
// }
