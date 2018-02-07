package main

//go:generate java -Xmx500M org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener -o parser Rune.g4

import (
	"fmt"
	"os"

	"github.com/mikijov/rune/parser"
	"github.com/mikijov/rune/vm"

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

func compile(input antlr.CharStream, ext vm.Externals) (vm.Program, Messages) {
	lexer := parser.NewRuneLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewRuneParser(stream)
	p.BuildParseTrees = true

	p.RemoveErrorListeners()
	errors := NewRuneErrorListener()
	p.AddErrorListener(errors)

	tree := p.Program()

	visitor := parser.NewMyVisitor(errors, ext)
	context := tree.(*parser.ProgramContext)
	program := visitor.VisitProgram(context)

	return program, errors
}

func Compile(code string, ext vm.Externals) (vm.Program, Messages) {
	input := antlr.NewInputStream(code)
	return compile(input, ext)
}

func CompileFile(filename string, ext vm.Externals) (vm.Program, Messages) {
	input := antlr.NewFileStream(filename)
	return compile(input, ext)
}

func main() {
	externals := vm.NewExternals()
	externals.DeclareUserFunction("hello", helloWorld)
	externals.DeclareUserFunction("add", add)

	program, messages := CompileFile(os.Args[1], externals)

	// check for errors
	if len(messages.GetErrors()) > 0 {
		for _, msg := range messages.GetErrors() {
			fmt.Println(msg)
		}
	} else {
		env := vm.NewEnvironment(nil)
		env.Import(externals)
		program.Execute(env)
	}
}
