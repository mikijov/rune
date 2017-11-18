// Generated from Rune.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 77, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 6, 8, 39, 10, 8, 13, 8, 14, 8, 40, 3, 9, 7, 9, 44, 10, 9, 12, 9, 14,
	9, 47, 11, 9, 3, 9, 3, 9, 6, 9, 51, 10, 9, 13, 9, 14, 9, 52, 3, 10, 5,
	10, 56, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 6, 11, 63, 10, 11, 13,
	11, 14, 11, 64, 3, 11, 3, 11, 3, 12, 3, 12, 7, 12, 71, 10, 12, 12, 12,
	14, 12, 74, 11, 12, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 5, 3, 2, 50, 59,
	4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 82, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3, 2, 2,
	2, 7, 29, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 33, 3, 2, 2, 2, 13, 35, 3,
	2, 2, 2, 15, 38, 3, 2, 2, 2, 17, 45, 3, 2, 2, 2, 19, 55, 3, 2, 2, 2, 21,
	62, 3, 2, 2, 2, 23, 68, 3, 2, 2, 2, 25, 26, 7, 61, 2, 2, 26, 4, 3, 2, 2,
	2, 27, 28, 7, 45, 2, 2, 28, 6, 3, 2, 2, 2, 29, 30, 7, 47, 2, 2, 30, 8,
	3, 2, 2, 2, 31, 32, 7, 44, 2, 2, 32, 10, 3, 2, 2, 2, 33, 34, 7, 49, 2,
	2, 34, 12, 3, 2, 2, 2, 35, 36, 7, 39, 2, 2, 36, 14, 3, 2, 2, 2, 37, 39,
	9, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2,
	40, 41, 3, 2, 2, 2, 41, 16, 3, 2, 2, 2, 42, 44, 9, 2, 2, 2, 43, 42, 3,
	2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46,
	48, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 50, 7, 48, 2, 2, 49, 51, 9, 2,
	2, 2, 50, 49, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53,
	3, 2, 2, 2, 53, 18, 3, 2, 2, 2, 54, 56, 7, 15, 2, 2, 55, 54, 3, 2, 2, 2,
	55, 56, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 58, 7, 12, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 60, 8, 10, 2, 2, 60, 20, 3, 2, 2, 2, 61, 63, 9, 3, 2, 2, 62,
	61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2,
	2, 65, 66, 3, 2, 2, 2, 66, 67, 8, 11, 2, 2, 67, 22, 3, 2, 2, 2, 68, 72,
	7, 37, 2, 2, 69, 71, 10, 4, 2, 2, 70, 69, 3, 2, 2, 2, 71, 74, 3, 2, 2,
	2, 72, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 72,
	3, 2, 2, 2, 75, 76, 8, 12, 2, 2, 76, 24, 3, 2, 2, 2, 9, 2, 40, 45, 52,
	55, 64, 72, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'+'", "'-'", "'*'", "'/'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "INTEGER_LITERAL", "REAL_LITERAL", "LINENDING",
	"WHITESPACE", "COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "INTEGER_LITERAL", "REAL_LITERAL",
	"LINENDING", "WHITESPACE", "COMMENT",
}

type RuneLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewRuneLexer(input antlr.CharStream) *RuneLexer {

	l := new(RuneLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Rune.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RuneLexer tokens.
const (
	RuneLexerT__0            = 1
	RuneLexerT__1            = 2
	RuneLexerT__2            = 3
	RuneLexerT__3            = 4
	RuneLexerT__4            = 5
	RuneLexerT__5            = 6
	RuneLexerINTEGER_LITERAL = 7
	RuneLexerREAL_LITERAL    = 8
	RuneLexerLINENDING       = 9
	RuneLexerWHITESPACE      = 10
	RuneLexerCOMMENT         = 11
)
