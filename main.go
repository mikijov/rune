package main

//go:generate java -Xmx500M org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener -o parser Rune.g4

import (
	"fmt"
	"os"

	"mikijov/rune/parser"
	"mikijov/rune/vm"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewRuneLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewRuneParser(stream)
	p.BuildParseTrees = true

	p.RemoveErrorListeners()
	errors := NewRuneErrorListener()
	p.AddErrorListener(errors)

	tree := p.Program()
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	// check for errors
	if len(errors.messages) > 0 {
		for _, msg := range errors.messages {
			fmt.Println(msg)
		}
	}

	// convert from parser AST to VM AST
	visitor := parser.NewMyVisitor(errors)
	context := tree.(*parser.ProgramContext)
	program := visitor.VisitProgram(context)
	// program is the compiled code

	env := vm.NewEnvironment(nil)
	program.Execute(env)
}
