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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 36, 212,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	6, 29, 156, 10, 29, 13, 29, 14, 29, 157, 3, 30, 7, 30, 161, 10, 30, 12,
	30, 14, 30, 164, 11, 30, 3, 30, 3, 30, 6, 30, 168, 10, 30, 13, 30, 14,
	30, 169, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	5, 31, 181, 10, 31, 3, 32, 3, 32, 7, 32, 185, 10, 32, 12, 32, 14, 32, 188,
	11, 32, 3, 33, 5, 33, 191, 10, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 6,
	34, 198, 10, 34, 13, 34, 14, 34, 199, 3, 34, 3, 34, 3, 35, 3, 35, 7, 35,
	206, 10, 35, 12, 35, 14, 35, 209, 11, 35, 3, 35, 3, 35, 2, 2, 36, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 3, 2, 7, 3, 2, 50, 59, 5, 2, 67,
	92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 11, 11,
	34, 34, 4, 2, 12, 12, 15, 15, 2, 219, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 3, 71, 3, 2, 2, 2, 5, 73, 3, 2, 2, 2, 7, 76,
	3, 2, 2, 2, 9, 78, 3, 2, 2, 2, 11, 80, 3, 2, 2, 2, 13, 84, 3, 2, 2, 2,
	15, 89, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19, 101, 3, 2, 2, 2, 21, 106, 3,
	2, 2, 2, 23, 108, 3, 2, 2, 2, 25, 110, 3, 2, 2, 2, 27, 112, 3, 2, 2, 2,
	29, 114, 3, 2, 2, 2, 31, 116, 3, 2, 2, 2, 33, 123, 3, 2, 2, 2, 35, 126,
	3, 2, 2, 2, 37, 131, 3, 2, 2, 2, 39, 133, 3, 2, 2, 2, 41, 136, 3, 2, 2,
	2, 43, 140, 3, 2, 2, 2, 45, 142, 3, 2, 2, 2, 47, 144, 3, 2, 2, 2, 49, 146,
	3, 2, 2, 2, 51, 148, 3, 2, 2, 2, 53, 150, 3, 2, 2, 2, 55, 152, 3, 2, 2,
	2, 57, 155, 3, 2, 2, 2, 59, 162, 3, 2, 2, 2, 61, 180, 3, 2, 2, 2, 63, 182,
	3, 2, 2, 2, 65, 190, 3, 2, 2, 2, 67, 197, 3, 2, 2, 2, 69, 203, 3, 2, 2,
	2, 71, 72, 7, 61, 2, 2, 72, 4, 3, 2, 2, 2, 73, 74, 7, 60, 2, 2, 74, 75,
	7, 63, 2, 2, 75, 6, 3, 2, 2, 2, 76, 77, 7, 60, 2, 2, 77, 8, 3, 2, 2, 2,
	78, 79, 7, 63, 2, 2, 79, 10, 3, 2, 2, 2, 80, 81, 7, 107, 2, 2, 81, 82,
	7, 112, 2, 2, 82, 83, 7, 118, 2, 2, 83, 12, 3, 2, 2, 2, 84, 85, 7, 116,
	2, 2, 85, 86, 7, 103, 2, 2, 86, 87, 7, 99, 2, 2, 87, 88, 7, 110, 2, 2,
	88, 14, 3, 2, 2, 2, 89, 90, 7, 117, 2, 2, 90, 91, 7, 118, 2, 2, 91, 92,
	7, 116, 2, 2, 92, 93, 7, 107, 2, 2, 93, 94, 7, 112, 2, 2, 94, 95, 7, 105,
	2, 2, 95, 16, 3, 2, 2, 2, 96, 97, 7, 100, 2, 2, 97, 98, 7, 113, 2, 2, 98,
	99, 7, 113, 2, 2, 99, 100, 7, 110, 2, 2, 100, 18, 3, 2, 2, 2, 101, 102,
	7, 104, 2, 2, 102, 103, 7, 119, 2, 2, 103, 104, 7, 112, 2, 2, 104, 105,
	7, 101, 2, 2, 105, 20, 3, 2, 2, 2, 106, 107, 7, 42, 2, 2, 107, 22, 3, 2,
	2, 2, 108, 109, 7, 46, 2, 2, 109, 24, 3, 2, 2, 2, 110, 111, 7, 43, 2, 2,
	111, 26, 3, 2, 2, 2, 112, 113, 7, 125, 2, 2, 113, 28, 3, 2, 2, 2, 114,
	115, 7, 127, 2, 2, 115, 30, 3, 2, 2, 2, 116, 117, 7, 116, 2, 2, 117, 118,
	7, 103, 2, 2, 118, 119, 7, 118, 2, 2, 119, 120, 7, 119, 2, 2, 120, 121,
	7, 116, 2, 2, 121, 122, 7, 112, 2, 2, 122, 32, 3, 2, 2, 2, 123, 124, 7,
	107, 2, 2, 124, 125, 7, 104, 2, 2, 125, 34, 3, 2, 2, 2, 126, 127, 7, 103,
	2, 2, 127, 128, 7, 110, 2, 2, 128, 129, 7, 117, 2, 2, 129, 130, 7, 103,
	2, 2, 130, 36, 3, 2, 2, 2, 131, 132, 7, 47, 2, 2, 132, 38, 3, 2, 2, 2,
	133, 134, 7, 113, 2, 2, 134, 135, 7, 116, 2, 2, 135, 40, 3, 2, 2, 2, 136,
	137, 7, 99, 2, 2, 137, 138, 7, 112, 2, 2, 138, 139, 7, 102, 2, 2, 139,
	42, 3, 2, 2, 2, 140, 141, 7, 126, 2, 2, 141, 44, 3, 2, 2, 2, 142, 143,
	7, 96, 2, 2, 143, 46, 3, 2, 2, 2, 144, 145, 7, 40, 2, 2, 145, 48, 3, 2,
	2, 2, 146, 147, 7, 44, 2, 2, 147, 50, 3, 2, 2, 2, 148, 149, 7, 49, 2, 2,
	149, 52, 3, 2, 2, 2, 150, 151, 7, 39, 2, 2, 151, 54, 3, 2, 2, 2, 152, 153,
	7, 45, 2, 2, 153, 56, 3, 2, 2, 2, 154, 156, 9, 2, 2, 2, 155, 154, 3, 2,
	2, 2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2,
	158, 58, 3, 2, 2, 2, 159, 161, 9, 2, 2, 2, 160, 159, 3, 2, 2, 2, 161, 164,
	3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2,
	2, 2, 164, 162, 3, 2, 2, 2, 165, 167, 7, 48, 2, 2, 166, 168, 9, 2, 2, 2,
	167, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169,
	170, 3, 2, 2, 2, 170, 60, 3, 2, 2, 2, 171, 172, 7, 118, 2, 2, 172, 173,
	7, 116, 2, 2, 173, 174, 7, 119, 2, 2, 174, 181, 7, 103, 2, 2, 175, 176,
	7, 104, 2, 2, 176, 177, 7, 99, 2, 2, 177, 178, 7, 110, 2, 2, 178, 179,
	7, 117, 2, 2, 179, 181, 7, 103, 2, 2, 180, 171, 3, 2, 2, 2, 180, 175, 3,
	2, 2, 2, 181, 62, 3, 2, 2, 2, 182, 186, 9, 3, 2, 2, 183, 185, 9, 4, 2,
	2, 184, 183, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 64, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 191, 7,
	15, 2, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2,
	2, 192, 193, 7, 12, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 8, 33, 2, 2,
	195, 66, 3, 2, 2, 2, 196, 198, 9, 5, 2, 2, 197, 196, 3, 2, 2, 2, 198, 199,
	3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 201, 3, 2,
	2, 2, 201, 202, 8, 34, 2, 2, 202, 68, 3, 2, 2, 2, 203, 207, 7, 37, 2, 2,
	204, 206, 10, 6, 2, 2, 205, 204, 3, 2, 2, 2, 206, 209, 3, 2, 2, 2, 207,
	205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2, 2, 2, 209, 207,
	3, 2, 2, 2, 210, 211, 8, 35, 2, 2, 211, 70, 3, 2, 2, 2, 11, 2, 157, 162,
	169, 180, 186, 190, 199, 207, 3, 8, 2, 2,
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
	"'func'", "'('", "','", "')'", "'{'", "'}'", "'return'", "'if'", "'else'",
	"'-'", "'or'", "'and'", "'|'", "'^'", "'&'", "'*'", "'/'", "'%'", "'+'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "INTEGER_LITERAL", "REAL_LITERAL",
	"BOOLEAN_LITERAL", "IDENTIFIER", "LINENDING", "WHITESPACE", "COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "INTEGER_LITERAL", "REAL_LITERAL", "BOOLEAN_LITERAL",
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
	RuneLexerT__19           = 20
	RuneLexerT__20           = 21
	RuneLexerT__21           = 22
	RuneLexerT__22           = 23
	RuneLexerT__23           = 24
	RuneLexerT__24           = 25
	RuneLexerT__25           = 26
	RuneLexerT__26           = 27
	RuneLexerINTEGER_LITERAL = 28
	RuneLexerREAL_LITERAL    = 29
	RuneLexerBOOLEAN_LITERAL = 30
	RuneLexerIDENTIFIER      = 31
	RuneLexerLINENDING       = 32
	RuneLexerWHITESPACE      = 33
	RuneLexerCOMMENT         = 34
)
