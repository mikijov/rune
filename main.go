package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*BaserunelangListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := NewrunelangLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewrunelangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Module()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
