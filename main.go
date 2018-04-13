// Copyright © 2018 Milutin Jovanović jovanovic.milutin@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//go:generate java -Xmx500M org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener -o parser Rune.g4

import (
	"fmt"
	"os"

	"github.com/mikijov/rune/parser"
	"github.com/mikijov/rune/vm"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func helloWorld() {
	fmt.Println("test1")
}

func add(x, y int64) int64 {
	return x + y
}

func compile(input antlr.CharStream, ext vm.Externals) (vm.Program, parser.Messages) {
	lexer := parser.NewRuneLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewRuneParser(stream)
	p.BuildParseTrees = true

	p.RemoveErrorListeners()
	errors := parser.NewRuneErrorListener()
	p.AddErrorListener(errors)

	tree := p.Program()

	visitor := parser.NewMyVisitor(errors, ext)
	context := tree.(*parser.ProgramContext)
	program := visitor.VisitProgram(context)

	return program, errors
}

// Compile compiles the contents of the passed string and returns a program that
// can be executed. ext if not nil provides externals which is how the program
// communicates to the outside world. Any functions specified in the externals
// are recognized as defined symbols in the code. Second return value contains
// messages produced while compiling the code. If there are any error messages
// returned, this indicates that the program is not runnable.
func Compile(code string, ext vm.Externals) (vm.Program, parser.Messages) {
	input := antlr.NewInputStream(code)
	return compile(input, ext)
}

// CompileFile is identical to the Compile() except that the code is read from
// the file whose path is passed in the filename parameter.
func CompileFile(filename string, ext vm.Externals) (vm.Program, parser.Messages) {
	input := antlr.NewFileStream(filename)
	return compile(input, ext)
}

func main() {
	externals := vm.NewExternals()
	if err := externals.DeclareUserFunction("hello", helloWorld); err != nil {
		panic(err.Error())
	}
	if err := externals.DeclareUserFunction("add", add); err != nil {
		panic(err.Error())
	}

	program, messages := CompileFile(os.Args[1], externals)

	// check for errors
	for _, msg := range messages.GetWarrnings() {
		fmt.Println(msg)
	}
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
