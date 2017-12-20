package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type runeErrorListener struct {
	Messages
	errors []string
}

func NewRuneErrorListener() *runeErrorListener {
	return new(runeErrorListener)
}

func (this *runeErrorListener) Error(line, column int, msg string) {
	this.errors = append(this.errors, fmt.Sprintf("%d:%d: %s", line, column, msg))
}

func (this *runeErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	this.errors = append(this.errors, fmt.Sprintf("%d:%d: %s", line, column, msg))
}

func (this *runeErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	this.errors = append(this.errors, "AMBIGUITY")
}

func (this *runeErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	this.errors = append(this.errors, "FULL CONTEXT")
}

func (this *runeErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	this.errors = append(this.errors, "CONTEXT SENSITIVITY")
}

func (this *runeErrorListener) GetErrors() []string {
	return this.errors
}
