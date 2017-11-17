package main

//go:generate java -Xmx500M org.antlr.v4.Tool -Dlanguage=Go -o parser Rune.g4

import (
	"fmt"
	"os"

	"mikijov/rune-antlr/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*parser.BaseRuneListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("%T:%s\n", ctx, ctx.GetText())
}

func MyAction(msg string) {
	fmt.Printf("MyAction: %s\n", msg)
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewRuneLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewRuneParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Module()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	fmt.Println("Running...")
	code := tree.GetM()
	code.Execute()
}
