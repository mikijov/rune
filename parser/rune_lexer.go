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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 63, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 37, 10,
	8, 13, 8, 14, 8, 38, 3, 9, 5, 9, 42, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 6, 10, 49, 10, 10, 13, 10, 14, 10, 50, 3, 10, 3, 10, 3, 11, 3, 11,
	7, 11, 57, 10, 11, 12, 11, 14, 11, 60, 11, 11, 3, 11, 3, 11, 2, 2, 12,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3,
	2, 5, 3, 2, 50, 59, 4, 2, 11, 11, 34, 34, 4, 2, 12, 12, 15, 15, 2, 66,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2,
	2, 2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 33,
	3, 2, 2, 2, 15, 36, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2, 19, 48, 3, 2, 2, 2,
	21, 54, 3, 2, 2, 2, 23, 24, 7, 61, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26, 7,
	45, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 47, 2, 2, 28, 8, 3, 2, 2, 2, 29,
	30, 7, 44, 2, 2, 30, 10, 3, 2, 2, 2, 31, 32, 7, 49, 2, 2, 32, 12, 3, 2,
	2, 2, 33, 34, 7, 39, 2, 2, 34, 14, 3, 2, 2, 2, 35, 37, 9, 2, 2, 2, 36,
	35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2,
	2, 39, 16, 3, 2, 2, 2, 40, 42, 7, 15, 2, 2, 41, 40, 3, 2, 2, 2, 41, 42,
	3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 7, 12, 2, 2, 44, 45, 3, 2, 2, 2,
	45, 46, 8, 9, 2, 2, 46, 18, 3, 2, 2, 2, 47, 49, 9, 3, 2, 2, 48, 47, 3,
	2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 53, 8, 10, 2, 2, 53, 20, 3, 2, 2, 2, 54, 58, 7, 37,
	2, 2, 55, 57, 10, 4, 2, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58,
	56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2,
	2, 61, 62, 8, 11, 2, 2, 62, 22, 3, 2, 2, 2, 7, 2, 38, 41, 50, 58, 3, 8,
	2, 2,
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
	"", "", "", "", "", "", "", "NUMBER", "LINENDING", "WHITESPACE", "COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "NUMBER", "LINENDING",
	"WHITESPACE", "COMMENT",
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
	RuneLexerT__0       = 1
	RuneLexerT__1       = 2
	RuneLexerT__2       = 3
	RuneLexerT__3       = 4
	RuneLexerT__4       = 5
	RuneLexerT__5       = 6
	RuneLexerNUMBER     = 7
	RuneLexerLINENDING  = 8
	RuneLexerWHITESPACE = 9
	RuneLexerCOMMENT    = 10
)
