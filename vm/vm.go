package vm

import ()

// type Vm interface {
// 	GetEnvironment() Environment
// 	Load(filename string)
// }
//
// type vm struct {
// 	environment Environment
// 	errors      runeErrorListener
// }
//
// func NewVm() Vm {
// 	return &vm{
// 		environment: NewEnvironment(nil),
// 		errors:      NewRuneErrorListener(),
// 	}
// }
//
// func (this *vm) GetEnvironment() Environment {
// 	return this.environment
// }
//
// func (this *vm) Load(filename string) {
// 	input := antlr.NewFileStream(os.Args[1])
// 	lexer := parser.NewRuneLexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, 0)
//
// 	p := parser.NewRuneParser(stream)
// 	p.BuildParseTrees = true
//
// 	p.RemoveErrorListeners()
// 	p.AddErrorListener(this.errors)
//
// 	tree := p.Program()
// 	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
//
// 	// check for errors
// 	if len(errors.messages) > 0 {
// 		for _, msg := range errors.messages {
// 			fmt.Println(msg)
// 		}
// 	}
//
// 	// convert from parser AST to VM AST
// 	visitor := parser.NewMyVisitor(errors)
// 	context := tree.(*parser.ProgramContext)
// 	program := visitor.VisitProgram(context)
// 	// program is the compiled code
//
// 	env := vm.NewEnvironment(nil)
// 	program.Execute(env)
// }
