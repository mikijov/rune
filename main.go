package main

//go:generate java -Xmx500M org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener -o parser Rune.g4

import (
	"fmt"
	"os"

	"mikijov/rune/parser"
	"mikijov/rune/vm"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Messages interface {
	GetErrors() []string
}

func helloWorld() {
	fmt.Println("test1")
}

func add(x, y int64) int64 {
	return x + y
}

func Compile(filename string) (vm.Program, Messages) {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewRuneLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewRuneParser(stream)
	p.BuildParseTrees = true

	p.RemoveErrorListeners()
	errors := NewRuneErrorListener()
	p.AddErrorListener(errors)

	tree := p.Program()

	visitor := parser.NewMyVisitor(errors)
	context := tree.(*parser.ProgramContext)
	program := visitor.VisitProgram(context)

	return program, errors
}

func main() {
	program, messages := Compile(os.Args[1])

	// check for errors
	if len(messages.GetErrors()) > 0 {
		for _, msg := range messages.GetErrors() {
			fmt.Println(msg)
		}
	} else {
		env := vm.NewEnvironment(nil)
		env.DeclareUserFunction("hello", helloWorld)
		env.DeclareUserFunction("add", add)
		program.Execute(env)
	}
}
