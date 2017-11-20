package vm

import ()

type instruction int

const (
	NOP = iota
	CMP
	JMP
	CALL
	RET
	JZ
	JNZ
	JG
	JGE
	JL
	JLE
	PUSH
	POP
	PUSHA
	POPA
	PUSHL
	POPL
	ADD
	SUB
	MUL
	DIV
	IDIV
	MOD
	SHL
	OR
	AND
	XOR
	NOT
)

type VmInteger int64
type VmReal float64

type Assembler interface {
	DefineInteger(val VmInteger) VmInteger
	DefineReal(val VmReal) VmInteger
	DefineBool(val bool) VmInteger
	DefineString(val string) VmInteger

	PushInteger(val VmInteger) VmInteger
	PushReal(val VmReal) VmInteger
	PushBool(val bool) VmInteger
	PushString(val string) VmInteger

	PushAbsolute(adr VmInteger) VmInteger
	PopAbsolute(adr VmInteger) VmInteger
	PushLocal(offset VmInteger) VmInteger
	PopLocal(offset VmInteger) VmInteger
	UpdateAddress(instructionAdress VmInteger, newVal VmInteger)

	Op(instr instruction) VmInteger

	// Link() code
}

type Vm interface {
	// Call(params Object[], returnCount int) Object[]
}

// func NewVm(code code) *Vm {
// }
