// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import (
	"fmt"
	"runtime"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"mikijov/rune-antlr/vm"
)

func trace(ctx antlr.ParserRuleContext) {
	return
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
	fmt.Printf("%s(%s)\n", f.Name(), ctx.GetText())

	if ctx != nil {
		fmt.Printf("%T\n", ctx)
		for i, child := range ctx.GetChildren() {
			switch child := child.(type) {
			case antlr.ParserRuleContext:
				fmt.Printf("%d:%s\n", i, child.GetText())
			case *antlr.TerminalNodeImpl:
				fmt.Printf("%d:%s\n", i, child.GetText())
			default:
				fmt.Printf("Unknown\n")
			}
		}
	}
}

type MyVisitor struct {
	*antlr.BaseParseTreeVisitor
	program vm.Program
	scope   vm.Scope
}

func NewMyVisitor() *MyVisitor {
	return &MyVisitor{
		program: vm.NewProgram(),
		scope:   vm.NewScope(nil),
	}
}

func (this *MyVisitor) VisitProgram(ctx *ProgramContext) vm.Program {
	trace(ctx)

	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *StatementContext:
			this.program.AddStatement(this.VisitStatement(child))
		case *antlr.TerminalNodeImpl:
			// do nothing
		default:
			fmt.Println(fmt.Sprintf("unknown type: %T\n", child))
		}
	}
	return this.program
}

func (this *MyVisitor) VisitStatement(ctx *StatementContext) vm.Statement {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *ExpressionContext:
		return vm.NewExpressionStatement(this.VisitExpression(child))
	case *DeclarationContext:
		return this.VisitDeclaration(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) VisitExpression(ctx interface{}) vm.Expression {
	switch ctx := ctx.(type) {
	case *ExpressionContext:
		return this.VisitExpression(ctx.GetChild(0))
	case *BinaryExpressionContext:
		return this.VisitBinaryExpression(ctx)
	case *LiteralPassthroughContext:
		return this.VisitLiteralPassthrough(ctx)
	case *RealLiteralContext:
		return this.VisitRealLiteral(ctx)
	case *IntegerLiteralContext:
		return this.VisitIntegerLiteral(ctx)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx))
	}
}

func (this *MyVisitor) VisitDeclaration(ctx *DeclarationContext) vm.Statement {
	trace(ctx)

	var type_ vm.Type
	if ctx.GetType_() != nil {
		type_ = this.VisitTypeName(ctx.GetType_())
	}

	var value vm.Expression
	if ctx.GetValue() != nil {
		value = this.VisitExpression(ctx.GetValue())
	}

	if type_ != "" && value != nil {
		if type_ != value.Type() {
			panic("type mismatch")
		}
	} else if type_ != "" {
		// already have type, nothing to do
	} else if value != nil {
		type_ = value.Type()
	} else {
		panic("variable type not provided") // grammar should not allow this
	}

	name := ctx.GetIdentifier().GetText()
	_, ok := this.scope.Declare(name, type_)
	if !ok {
		panic("variable redeclared")
	}

	if value != nil {
		return vm.NewDeclarationStatement(name, value)
	} else {
		return vm.NewDeclarationStatement(name, vm.NewZeroLiteral(type_))
	}
}

func (this *MyVisitor) VisitTypeName(ctx ITypeNameContext) vm.Type {
	trace(ctx)

	switch ctx.GetText() {
	case "int":
		return vm.INTEGER
	case "real":
		return vm.REAL
	case "string":
		return vm.STRING
	case "bool":
		return vm.BOOLEAN
	default:
		panic("unknown type") // grammar should not allow this
	}
}

func (this *MyVisitor) VisitLiteralPassthrough(ctx *LiteralPassthroughContext) vm.Expression {
	trace(ctx)

	switch child := ctx.GetChild(0).(type) {
	case *RealLiteralContext:
		return this.VisitRealLiteral(child)
	case *IntegerLiteralContext:
		return this.VisitIntegerLiteral(child)
	default:
		panic(fmt.Sprintf("unknown type: %T\n", ctx.GetChild(0)))
	}
}

func (this *MyVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) vm.Expression {
	trace(ctx)

	left := this.VisitExpression(ctx.left)
	right := this.VisitExpression(ctx.right)

	return vm.NewBinaryExpression(left, ctx.op.GetText(), right)
}

// func (this *MyVisitor) VisitExpressionPassthrough(ctx *ExpressionPassthroughContext) interface{} {
// 	trace()
// 	return this.VisitChildren(ctx)
// }

// func (this *MyVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
// 	trace()
// 	return this.VisitChildren(ctx)
// }

func (this *MyVisitor) VisitRealLiteral(ctx *RealLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewRealLiteral(ctx.GetText())
}

func (this *MyVisitor) VisitIntegerLiteral(ctx *IntegerLiteralContext) vm.Expression {
	trace(ctx)

	return vm.NewIntegerLiteral(ctx.GetText())
}
