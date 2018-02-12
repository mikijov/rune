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

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type runeErrorListener struct {
	Messages
	errors []string
}

// NewRuneErrorListener creates instance of ErrorListener which will store all
// messages produced by the compiler.
func NewRuneErrorListener() Messages {
	return &runeErrorListener{}
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
