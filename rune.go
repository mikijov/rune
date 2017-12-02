package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type runeErrorListener struct {
	messages []string
}

func NewRuneErrorListener() *runeErrorListener {
	return new(runeErrorListener)
}

func (this *runeErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	this.messages = append(this.messages, fmt.Sprintf("%d:%d: %s", line, column, msg))
}

func (this *runeErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	this.messages = append(this.messages, "AMBIGUITY")
}

func (this *runeErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	this.messages = append(this.messages, "FULL CONTEXT")
}

func (this *runeErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	this.messages = append(this.messages, "CONTEXT SENSITIVITY")
}

type Rune interface {
}
