package vm

import ()

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
