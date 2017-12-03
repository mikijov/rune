// Generated from Rune.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	"mikijov/rune/vm"
)

var _ vm.Type // inhibit unused import error

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 168,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 6, 21, 112,
	10, 21, 13, 21, 14, 21, 113, 3, 22, 7, 22, 117, 10, 22, 12, 22, 14, 22,
	120, 11, 22, 3, 22, 3, 22, 6, 22, 124, 10, 22, 13, 22, 14, 22, 125, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 137,
	10, 23, 3, 24, 3, 24, 7, 24, 141, 10, 24, 12, 24, 14, 24, 144, 11, 24,
	3, 25, 5, 25, 147, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 6, 26, 154,
	10, 26, 13, 26, 14, 26, 155, 3, 26, 3, 26, 3, 27, 3, 27, 7, 27, 162, 10,
	27, 12, 27, 14, 27, 165, 11, 27, 3, 27, 3, 27, 2, 2, 28, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 3, 2, 7, 3, 2, 50, 59, 5, 2, 67, 92,
	97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 11, 11, 34,
	34, 4, 2, 12, 12, 15, 15, 2, 175, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2,
	53, 3, 2, 2, 2, 3, 55, 3, 2, 2, 2, 5, 57, 3, 2, 2, 2, 7, 60, 3, 2, 2, 2,
	9, 62, 3, 2, 2, 2, 11, 64, 3, 2, 2, 2, 13, 68, 3, 2, 2, 2, 15, 73, 3, 2,
	2, 2, 17, 80, 3, 2, 2, 2, 19, 85, 3, 2, 2, 2, 21, 90, 3, 2, 2, 2, 23, 92,
	3, 2, 2, 2, 25, 94, 3, 2, 2, 2, 27, 96, 3, 2, 2, 2, 29, 98, 3, 2, 2, 2,
	31, 100, 3, 2, 2, 2, 33, 102, 3, 2, 2, 2, 35, 104, 3, 2, 2, 2, 37, 106,
	3, 2, 2, 2, 39, 108, 3, 2, 2, 2, 41, 111, 3, 2, 2, 2, 43, 118, 3, 2, 2,
	2, 45, 136, 3, 2, 2, 2, 47, 138, 3, 2, 2, 2, 49, 146, 3, 2, 2, 2, 51, 153,
	3, 2, 2, 2, 53, 159, 3, 2, 2, 2, 55, 56, 7, 61, 2, 2, 56, 4, 3, 2, 2, 2,
	57, 58, 7, 60, 2, 2, 58, 59, 7, 63, 2, 2, 59, 6, 3, 2, 2, 2, 60, 61, 7,
	60, 2, 2, 61, 8, 3, 2, 2, 2, 62, 63, 7, 63, 2, 2, 63, 10, 3, 2, 2, 2, 64,
	65, 7, 107, 2, 2, 65, 66, 7, 112, 2, 2, 66, 67, 7, 118, 2, 2, 67, 12, 3,
	2, 2, 2, 68, 69, 7, 116, 2, 2, 69, 70, 7, 103, 2, 2, 70, 71, 7, 99, 2,
	2, 71, 72, 7, 110, 2, 2, 72, 14, 3, 2, 2, 2, 73, 74, 7, 117, 2, 2, 74,
	75, 7, 118, 2, 2, 75, 76, 7, 116, 2, 2, 76, 77, 7, 107, 2, 2, 77, 78, 7,
	112, 2, 2, 78, 79, 7, 105, 2, 2, 79, 16, 3, 2, 2, 2, 80, 81, 7, 100, 2,
	2, 81, 82, 7, 113, 2, 2, 82, 83, 7, 113, 2, 2, 83, 84, 7, 110, 2, 2, 84,
	18, 3, 2, 2, 2, 85, 86, 7, 104, 2, 2, 86, 87, 7, 119, 2, 2, 87, 88, 7,
	112, 2, 2, 88, 89, 7, 101, 2, 2, 89, 20, 3, 2, 2, 2, 90, 91, 7, 42, 2,
	2, 91, 22, 3, 2, 2, 2, 92, 93, 7, 46, 2, 2, 93, 24, 3, 2, 2, 2, 94, 95,
	7, 43, 2, 2, 95, 26, 3, 2, 2, 2, 96, 97, 7, 125, 2, 2, 97, 28, 3, 2, 2,
	2, 98, 99, 7, 127, 2, 2, 99, 30, 3, 2, 2, 2, 100, 101, 7, 47, 2, 2, 101,
	32, 3, 2, 2, 2, 102, 103, 7, 44, 2, 2, 103, 34, 3, 2, 2, 2, 104, 105, 7,
	49, 2, 2, 105, 36, 3, 2, 2, 2, 106, 107, 7, 39, 2, 2, 107, 38, 3, 2, 2,
	2, 108, 109, 7, 45, 2, 2, 109, 40, 3, 2, 2, 2, 110, 112, 9, 2, 2, 2, 111,
	110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114,
	3, 2, 2, 2, 114, 42, 3, 2, 2, 2, 115, 117, 9, 2, 2, 2, 116, 115, 3, 2,
	2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2,
	119, 121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 123, 7, 48, 2, 2, 122,
	124, 9, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 123,
	3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 44, 3, 2, 2, 2, 127, 128, 7, 118,
	2, 2, 128, 129, 7, 116, 2, 2, 129, 130, 7, 119, 2, 2, 130, 137, 7, 103,
	2, 2, 131, 132, 7, 104, 2, 2, 132, 133, 7, 99, 2, 2, 133, 134, 7, 110,
	2, 2, 134, 135, 7, 117, 2, 2, 135, 137, 7, 103, 2, 2, 136, 127, 3, 2, 2,
	2, 136, 131, 3, 2, 2, 2, 137, 46, 3, 2, 2, 2, 138, 142, 9, 3, 2, 2, 139,
	141, 9, 4, 2, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140,
	3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 48, 3, 2, 2, 2, 144, 142, 3, 2,
	2, 2, 145, 147, 7, 15, 2, 2, 146, 145, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 149, 7, 12, 2, 2, 149, 150, 3, 2, 2, 2, 150,
	151, 8, 25, 2, 2, 151, 50, 3, 2, 2, 2, 152, 154, 9, 5, 2, 2, 153, 152,
	3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2,
	2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 8, 26, 2, 2, 158, 52, 3, 2, 2, 2,
	159, 163, 7, 37, 2, 2, 160, 162, 10, 6, 2, 2, 161, 160, 3, 2, 2, 2, 162,
	165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 166,
	3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 167, 8, 27, 2, 2, 167, 54, 3, 2,
	2, 2, 11, 2, 113, 118, 125, 136, 142, 146, 155, 163, 3, 8, 2, 2,
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
	"", "';'", "':='", "':'", "'='", "'int'", "'real'", "'string'", "'bool'",
	"'func'", "'('", "','", "')'", "'{'", "'}'", "'-'", "'*'", "'/'", "'%'",
	"'+'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "INTEGER_LITERAL", "REAL_LITERAL", "BOOLEAN_LITERAL", "IDENTIFIER",
	"LINENDING", "WHITESPACE", "COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "INTEGER_LITERAL", "REAL_LITERAL", "BOOLEAN_LITERAL",
	"IDENTIFIER", "LINENDING", "WHITESPACE", "COMMENT",
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
	RuneLexerT__6            = 7
	RuneLexerT__7            = 8
	RuneLexerT__8            = 9
	RuneLexerT__9            = 10
	RuneLexerT__10           = 11
	RuneLexerT__11           = 12
	RuneLexerT__12           = 13
	RuneLexerT__13           = 14
	RuneLexerT__14           = 15
	RuneLexerT__15           = 16
	RuneLexerT__16           = 17
	RuneLexerT__17           = 18
	RuneLexerT__18           = 19
	RuneLexerINTEGER_LITERAL = 20
	RuneLexerREAL_LITERAL    = 21
	RuneLexerBOOLEAN_LITERAL = 22
	RuneLexerIDENTIFIER      = 23
	RuneLexerLINENDING       = 24
	RuneLexerWHITESPACE      = 25
	RuneLexerCOMMENT         = 26
)
