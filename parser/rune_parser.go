// Generated from Rune.g4 by ANTLR 4.7.

package parser // Rune

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import (
	"github.com/mikijov/rune/vm"
)

var _ vm.Type // inhibit unused import error

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 260,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 7, 2, 40, 10, 2, 12, 2, 14, 2, 43, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 59, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 69, 10, 4, 5, 4, 71, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 84, 10, 6, 12, 6, 14, 6, 87, 11, 6, 5,
	6, 89, 10, 6, 3, 6, 3, 6, 3, 6, 5, 6, 94, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6,
	99, 10, 6, 12, 6, 14, 6, 102, 11, 6, 3, 6, 5, 6, 105, 10, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 112, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	7, 8, 120, 10, 8, 12, 8, 14, 8, 123, 11, 8, 5, 8, 125, 10, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 7, 9, 132, 10, 9, 12, 9, 14, 9, 135, 11, 9, 3, 9,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 143, 10, 10, 12, 10, 14, 10, 146,
	11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 154, 10, 11, 12,
	11, 14, 11, 157, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 5, 12, 163, 10, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 173, 10,
	13, 12, 13, 14, 13, 176, 11, 13, 3, 13, 3, 13, 5, 13, 180, 10, 13, 3, 14,
	3, 14, 3, 14, 5, 14, 185, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 201,
	10, 16, 12, 16, 14, 16, 204, 11, 16, 5, 16, 206, 10, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 217, 10, 16, 3,
	16, 3, 16, 3, 16, 5, 16, 222, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 246, 10, 16, 12, 16,
	14, 16, 249, 11, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 5,
	19, 258, 10, 19, 3, 19, 2, 3, 30, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 2, 6, 3, 2, 24, 27, 3, 2, 28, 31, 3, 2,
	32, 37, 4, 2, 6, 6, 40, 47, 2, 285, 2, 41, 3, 2, 2, 2, 4, 58, 3, 2, 2,
	2, 6, 70, 3, 2, 2, 2, 8, 72, 3, 2, 2, 2, 10, 104, 3, 2, 2, 2, 12, 106,
	3, 2, 2, 2, 14, 115, 3, 2, 2, 2, 16, 128, 3, 2, 2, 2, 18, 139, 3, 2, 2,
	2, 20, 151, 3, 2, 2, 2, 22, 160, 3, 2, 2, 2, 24, 164, 3, 2, 2, 2, 26, 181,
	3, 2, 2, 2, 28, 188, 3, 2, 2, 2, 30, 221, 3, 2, 2, 2, 32, 250, 3, 2, 2,
	2, 34, 252, 3, 2, 2, 2, 36, 257, 3, 2, 2, 2, 38, 40, 5, 4, 3, 2, 39, 38,
	3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2,
	42, 44, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 7, 2, 2, 3, 45, 3, 3, 2,
	2, 2, 46, 47, 5, 6, 4, 2, 47, 48, 7, 3, 2, 2, 48, 59, 3, 2, 2, 2, 49, 50,
	5, 22, 12, 2, 50, 51, 7, 3, 2, 2, 51, 59, 3, 2, 2, 2, 52, 53, 5, 28, 15,
	2, 53, 54, 7, 3, 2, 2, 54, 59, 3, 2, 2, 2, 55, 59, 5, 12, 7, 2, 56, 59,
	5, 24, 13, 2, 57, 59, 5, 26, 14, 2, 58, 46, 3, 2, 2, 2, 58, 49, 3, 2, 2,
	2, 58, 52, 3, 2, 2, 2, 58, 55, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 57,
	3, 2, 2, 2, 59, 5, 3, 2, 2, 2, 60, 61, 7, 51, 2, 2, 61, 62, 7, 4, 2, 2,
	62, 71, 5, 28, 15, 2, 63, 64, 7, 51, 2, 2, 64, 65, 7, 5, 2, 2, 65, 68,
	5, 8, 5, 2, 66, 67, 7, 6, 2, 2, 67, 69, 5, 28, 15, 2, 68, 66, 3, 2, 2,
	2, 68, 69, 3, 2, 2, 2, 69, 71, 3, 2, 2, 2, 70, 60, 3, 2, 2, 2, 70, 63,
	3, 2, 2, 2, 71, 7, 3, 2, 2, 2, 72, 73, 5, 10, 6, 2, 73, 9, 3, 2, 2, 2,
	74, 105, 7, 7, 2, 2, 75, 105, 7, 8, 2, 2, 76, 105, 7, 9, 2, 2, 77, 105,
	7, 10, 2, 2, 78, 79, 7, 11, 2, 2, 79, 88, 7, 12, 2, 2, 80, 85, 5, 10, 6,
	2, 81, 82, 7, 13, 2, 2, 82, 84, 5, 10, 6, 2, 83, 81, 3, 2, 2, 2, 84, 87,
	3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2,
	87, 85, 3, 2, 2, 2, 88, 80, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 93, 7, 14, 2, 2, 91, 92, 7, 5, 2, 2, 92, 94, 5, 10, 6, 2,
	93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 105, 3, 2, 2, 2, 95, 96, 7,
	15, 2, 2, 96, 100, 7, 16, 2, 2, 97, 99, 5, 18, 10, 2, 98, 97, 3, 2, 2,
	2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101,
	103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 7, 17, 2, 2, 104, 74,
	3, 2, 2, 2, 104, 75, 3, 2, 2, 2, 104, 76, 3, 2, 2, 2, 104, 77, 3, 2, 2,
	2, 104, 78, 3, 2, 2, 2, 104, 95, 3, 2, 2, 2, 105, 11, 3, 2, 2, 2, 106,
	107, 7, 11, 2, 2, 107, 108, 7, 51, 2, 2, 108, 111, 5, 14, 8, 2, 109, 110,
	7, 5, 2, 2, 110, 112, 5, 8, 5, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2,
	2, 2, 112, 113, 3, 2, 2, 2, 113, 114, 5, 20, 11, 2, 114, 13, 3, 2, 2, 2,
	115, 124, 7, 12, 2, 2, 116, 121, 5, 16, 9, 2, 117, 118, 7, 13, 2, 2, 118,
	120, 5, 16, 9, 2, 119, 117, 3, 2, 2, 2, 120, 123, 3, 2, 2, 2, 121, 119,
	3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2,
	2, 2, 124, 116, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2,
	126, 127, 7, 14, 2, 2, 127, 15, 3, 2, 2, 2, 128, 133, 7, 51, 2, 2, 129,
	130, 7, 13, 2, 2, 130, 132, 7, 51, 2, 2, 131, 129, 3, 2, 2, 2, 132, 135,
	3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 3, 2,
	2, 2, 135, 133, 3, 2, 2, 2, 136, 137, 7, 5, 2, 2, 137, 138, 5, 8, 5, 2,
	138, 17, 3, 2, 2, 2, 139, 144, 7, 51, 2, 2, 140, 141, 7, 13, 2, 2, 141,
	143, 7, 51, 2, 2, 142, 140, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142,
	3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 146, 144, 3, 2,
	2, 2, 147, 148, 7, 5, 2, 2, 148, 149, 5, 8, 5, 2, 149, 150, 7, 3, 2, 2,
	150, 19, 3, 2, 2, 2, 151, 155, 7, 16, 2, 2, 152, 154, 5, 4, 3, 2, 153,
	152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156,
	3, 2, 2, 2, 156, 158, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 159, 7, 17,
	2, 2, 159, 21, 3, 2, 2, 2, 160, 162, 7, 18, 2, 2, 161, 163, 5, 28, 15,
	2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 23, 3, 2, 2, 2, 164,
	165, 7, 19, 2, 2, 165, 166, 5, 28, 15, 2, 166, 174, 5, 20, 11, 2, 167,
	168, 7, 20, 2, 2, 168, 169, 7, 19, 2, 2, 169, 170, 5, 28, 15, 2, 170, 171,
	5, 20, 11, 2, 171, 173, 3, 2, 2, 2, 172, 167, 3, 2, 2, 2, 173, 176, 3,
	2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 179, 3, 2, 2,
	2, 176, 174, 3, 2, 2, 2, 177, 178, 7, 20, 2, 2, 178, 180, 5, 20, 11, 2,
	179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 25, 3, 2, 2, 2, 181, 184,
	7, 21, 2, 2, 182, 183, 7, 22, 2, 2, 183, 185, 5, 28, 15, 2, 184, 182, 3,
	2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 5, 20, 11,
	2, 187, 27, 3, 2, 2, 2, 188, 189, 5, 30, 16, 2, 189, 29, 3, 2, 2, 2, 190,
	191, 8, 16, 1, 2, 191, 192, 7, 12, 2, 2, 192, 193, 5, 30, 16, 2, 193, 194,
	7, 14, 2, 2, 194, 222, 3, 2, 2, 2, 195, 196, 7, 51, 2, 2, 196, 205, 7,
	12, 2, 2, 197, 202, 5, 30, 16, 2, 198, 199, 7, 13, 2, 2, 199, 201, 5, 30,
	16, 2, 200, 198, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2,
	202, 203, 3, 2, 2, 2, 203, 206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205,
	197, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 222,
	7, 14, 2, 2, 208, 209, 5, 32, 17, 2, 209, 210, 5, 30, 16, 12, 210, 222,
	3, 2, 2, 2, 211, 222, 5, 36, 19, 2, 212, 213, 7, 11, 2, 2, 213, 216, 5,
	14, 8, 2, 214, 215, 7, 5, 2, 2, 215, 217, 5, 8, 5, 2, 216, 214, 3, 2, 2,
	2, 216, 217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 5, 20, 11, 2,
	219, 222, 3, 2, 2, 2, 220, 222, 7, 51, 2, 2, 221, 190, 3, 2, 2, 2, 221,
	195, 3, 2, 2, 2, 221, 208, 3, 2, 2, 2, 221, 211, 3, 2, 2, 2, 221, 212,
	3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 247, 3, 2, 2, 2, 223, 224, 12, 11,
	2, 2, 224, 225, 9, 2, 2, 2, 225, 246, 5, 30, 16, 12, 226, 227, 12, 10,
	2, 2, 227, 228, 9, 3, 2, 2, 228, 246, 5, 30, 16, 11, 229, 230, 12, 9, 2,
	2, 230, 231, 9, 4, 2, 2, 231, 246, 5, 30, 16, 10, 232, 233, 12, 8, 2, 2,
	233, 234, 7, 38, 2, 2, 234, 246, 5, 30, 16, 9, 235, 236, 12, 7, 2, 2, 236,
	237, 7, 39, 2, 2, 237, 246, 5, 30, 16, 8, 238, 239, 12, 6, 2, 2, 239, 240,
	5, 34, 18, 2, 240, 241, 5, 30, 16, 7, 241, 246, 3, 2, 2, 2, 242, 243, 12,
	13, 2, 2, 243, 244, 7, 23, 2, 2, 244, 246, 7, 51, 2, 2, 245, 223, 3, 2,
	2, 2, 245, 226, 3, 2, 2, 2, 245, 229, 3, 2, 2, 2, 245, 232, 3, 2, 2, 2,
	245, 235, 3, 2, 2, 2, 245, 238, 3, 2, 2, 2, 245, 242, 3, 2, 2, 2, 246,
	249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 31, 3,
	2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251, 7, 29, 2, 2, 251, 33, 3, 2, 2,
	2, 252, 253, 9, 5, 2, 2, 253, 35, 3, 2, 2, 2, 254, 258, 7, 49, 2, 2, 255,
	258, 7, 48, 2, 2, 256, 258, 7, 50, 2, 2, 257, 254, 3, 2, 2, 2, 257, 255,
	3, 2, 2, 2, 257, 256, 3, 2, 2, 2, 258, 37, 3, 2, 2, 2, 28, 41, 58, 68,
	70, 85, 88, 93, 100, 104, 111, 121, 124, 133, 144, 155, 162, 174, 179,
	184, 202, 205, 216, 221, 245, 247, 257,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "':='", "':'", "'='", "'int'", "'real'", "'string'", "'bool'",
	"'func'", "'('", "','", "')'", "'struct'", "'{'", "'}'", "'return'", "'if'",
	"'else'", "'loop'", "'while'", "'.'", "'*'", "'/'", "'%'", "'&'", "'+'",
	"'-'", "'|'", "'^'", "'<'", "'>'", "'=='", "'>='", "'<='", "'!='", "'and'",
	"'or'", "'+='", "'-='", "'*='", "'/='", "'%='", "'&='", "'|='", "'^='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "INTEGER_LITERAL", "REAL_LITERAL",
	"BOOLEAN_LITERAL", "IDENTIFIER", "LINENDING", "WHITESPACE", "COMMENT",
}

var ruleNames = []string{
	"program", "statement", "declaration", "typeName", "typeName2", "function",
	"paramDecl", "combinedParam", "combinedField", "scope", "returnStatement",
	"ifStatement", "loop", "expression", "expression2", "unaryOp", "assignmentOp",
	"literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type RuneParser struct {
	*antlr.BaseParser
}

func NewRuneParser(input antlr.TokenStream) *RuneParser {
	this := new(RuneParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Rune.g4"

	return this
}

// RuneParser tokens.
const (
	RuneParserEOF             = antlr.TokenEOF
	RuneParserT__0            = 1
	RuneParserT__1            = 2
	RuneParserT__2            = 3
	RuneParserT__3            = 4
	RuneParserT__4            = 5
	RuneParserT__5            = 6
	RuneParserT__6            = 7
	RuneParserT__7            = 8
	RuneParserT__8            = 9
	RuneParserT__9            = 10
	RuneParserT__10           = 11
	RuneParserT__11           = 12
	RuneParserT__12           = 13
	RuneParserT__13           = 14
	RuneParserT__14           = 15
	RuneParserT__15           = 16
	RuneParserT__16           = 17
	RuneParserT__17           = 18
	RuneParserT__18           = 19
	RuneParserT__19           = 20
	RuneParserT__20           = 21
	RuneParserT__21           = 22
	RuneParserT__22           = 23
	RuneParserT__23           = 24
	RuneParserT__24           = 25
	RuneParserT__25           = 26
	RuneParserT__26           = 27
	RuneParserT__27           = 28
	RuneParserT__28           = 29
	RuneParserT__29           = 30
	RuneParserT__30           = 31
	RuneParserT__31           = 32
	RuneParserT__32           = 33
	RuneParserT__33           = 34
	RuneParserT__34           = 35
	RuneParserT__35           = 36
	RuneParserT__36           = 37
	RuneParserT__37           = 38
	RuneParserT__38           = 39
	RuneParserT__39           = 40
	RuneParserT__40           = 41
	RuneParserT__41           = 42
	RuneParserT__42           = 43
	RuneParserT__43           = 44
	RuneParserT__44           = 45
	RuneParserINTEGER_LITERAL = 46
	RuneParserREAL_LITERAL    = 47
	RuneParserBOOLEAN_LITERAL = 48
	RuneParserIDENTIFIER      = 49
	RuneParserLINENDING       = 50
	RuneParserWHITESPACE      = 51
	RuneParserCOMMENT         = 52
)

// RuneParser rules.
const (
	RuneParserRULE_program         = 0
	RuneParserRULE_statement       = 1
	RuneParserRULE_declaration     = 2
	RuneParserRULE_typeName        = 3
	RuneParserRULE_typeName2       = 4
	RuneParserRULE_function        = 5
	RuneParserRULE_paramDecl       = 6
	RuneParserRULE_combinedParam   = 7
	RuneParserRULE_combinedField   = 8
	RuneParserRULE_scope           = 9
	RuneParserRULE_returnStatement = 10
	RuneParserRULE_ifStatement     = 11
	RuneParserRULE_loop            = 12
	RuneParserRULE_expression      = 13
	RuneParserRULE_expression2     = 14
	RuneParserRULE_unaryOp         = 15
	RuneParserRULE_assignmentOp    = 16
	RuneParserRULE_literal         = 17
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_statement IStatementContext
	statements []IStatementContext
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Get_statement() IStatementContext { return s._statement }

func (s *ProgramContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *ProgramContext) GetStatements() []IStatementContext { return s.statements }

func (s *ProgramContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuneParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuneParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__8)|(1<<RuneParserT__9)|(1<<RuneParserT__15)|(1<<RuneParserT__16)|(1<<RuneParserT__18)|(1<<RuneParserT__26))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(RuneParserINTEGER_LITERAL-46))|(1<<(RuneParserREAL_LITERAL-46))|(1<<(RuneParserBOOLEAN_LITERAL-46))|(1<<(RuneParserIDENTIFIER-46)))) != 0) {
		{
			p.SetState(36)

			var _x = p.Statement()

			localctx.(*ProgramContext)._statement = _x
		}
		localctx.(*ProgramContext).statements = append(localctx.(*ProgramContext).statements, localctx.(*ProgramContext)._statement)

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(RuneParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuneParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Declaration()
		}
		{
			p.SetState(45)
			p.Match(RuneParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.ReturnStatement()
		}
		{
			p.SetState(48)
			p.Match(RuneParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Expression()
		}
		{
			p.SetState(51)
			p.Match(RuneParserT__0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.Function()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)
			p.IfStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(55)
			p.Loop()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifier returns the identifier token.
	GetIdentifier() antlr.Token

	// SetIdentifier sets the identifier token.
	SetIdentifier(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// GetType_ returns the type_ rule contexts.
	GetType_() ITypeNameContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// SetType_ sets the type_ rule contexts.
	SetType_(ITypeNameContext)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	identifier antlr.Token
	value      IExpressionContext
	type_      ITypeNameContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *DeclarationContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *DeclarationContext) GetValue() IExpressionContext { return s.value }

func (s *DeclarationContext) GetType_() ITypeNameContext { return s.type_ }

func (s *DeclarationContext) SetValue(v IExpressionContext) { s.value = v }

func (s *DeclarationContext) SetType_(v ITypeNameContext) { s.type_ = v }

func (s *DeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RuneParserRULE_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)

			var _m = p.Match(RuneParserIDENTIFIER)

			localctx.(*DeclarationContext).identifier = _m
		}
		{
			p.SetState(59)
			p.Match(RuneParserT__1)
		}
		{
			p.SetState(60)

			var _x = p.Expression()

			localctx.(*DeclarationContext).value = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)

			var _m = p.Match(RuneParserIDENTIFIER)

			localctx.(*DeclarationContext).identifier = _m
		}
		{
			p.SetState(62)
			p.Match(RuneParserT__2)
		}
		{
			p.SetState(63)

			var _x = p.TypeName()

			localctx.(*DeclarationContext).type_ = _x
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RuneParserT__3 {
			{
				p.SetState(64)
				p.Match(RuneParserT__3)
			}
			{
				p.SetState(65)

				var _x = p.Expression()

				localctx.(*DeclarationContext).value = _x
			}

		}

	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) TypeName2() ITypeName2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeName2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeName2Context)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RuneParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.TypeName2()
	}

	return localctx
}

// ITypeName2Context is an interface to support dynamic dispatch.
type ITypeName2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeName2Context differentiates from other interfaces.
	IsTypeName2Context()
}

type TypeName2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeName2Context() *TypeName2Context {
	var p = new(TypeName2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_typeName2
	return p
}

func (*TypeName2Context) IsTypeName2Context() {}

func NewTypeName2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeName2Context {
	var p = new(TypeName2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_typeName2

	return p
}

func (s *TypeName2Context) GetParser() antlr.Parser { return s.parser }

func (s *TypeName2Context) CopyFrom(ctx *TypeName2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeName2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeName2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleTypeContext struct {
	*TypeName2Context
}

func NewSimpleTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleTypeContext {
	var p = new(SimpleTypeContext)

	p.TypeName2Context = NewEmptyTypeName2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeName2Context))

	return p
}

func (s *SimpleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitSimpleType(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructTypeContext struct {
	*TypeName2Context
}

func NewStructTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeContext {
	var p = new(StructTypeContext)

	p.TypeName2Context = NewEmptyTypeName2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeName2Context))

	return p
}

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeContext) AllCombinedField() []ICombinedFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICombinedFieldContext)(nil)).Elem())
	var tst = make([]ICombinedFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICombinedFieldContext)
		}
	}

	return tst
}

func (s *StructTypeContext) CombinedField(i int) ICombinedFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICombinedFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICombinedFieldContext)
}

func (s *StructTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitStructType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionTypeContext struct {
	*TypeName2Context
	_typeName2 ITypeName2Context
	paramTypes []ITypeName2Context
	returnType ITypeName2Context
}

func NewFunctionTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.TypeName2Context = NewEmptyTypeName2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeName2Context))

	return p
}

func (s *FunctionTypeContext) Get_typeName2() ITypeName2Context { return s._typeName2 }

func (s *FunctionTypeContext) GetReturnType() ITypeName2Context { return s.returnType }

func (s *FunctionTypeContext) Set_typeName2(v ITypeName2Context) { s._typeName2 = v }

func (s *FunctionTypeContext) SetReturnType(v ITypeName2Context) { s.returnType = v }

func (s *FunctionTypeContext) GetParamTypes() []ITypeName2Context { return s.paramTypes }

func (s *FunctionTypeContext) SetParamTypes(v []ITypeName2Context) { s.paramTypes = v }

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) AllTypeName2() []ITypeName2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeName2Context)(nil)).Elem())
	var tst = make([]ITypeName2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeName2Context)
		}
	}

	return tst
}

func (s *FunctionTypeContext) TypeName2(i int) ITypeName2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeName2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeName2Context)
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) TypeName2() (localctx ITypeName2Context) {
	localctx = NewTypeName2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RuneParserRULE_typeName2)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuneParserT__4:
		localctx = NewSimpleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(RuneParserT__4)
		}

	case RuneParserT__5:
		localctx = NewSimpleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(RuneParserT__5)
		}

	case RuneParserT__6:
		localctx = NewSimpleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.Match(RuneParserT__6)
		}

	case RuneParserT__7:
		localctx = NewSimpleTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(RuneParserT__7)
		}

	case RuneParserT__8:
		localctx = NewFunctionTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.Match(RuneParserT__8)
		}
		{
			p.SetState(77)
			p.Match(RuneParserT__9)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__4)|(1<<RuneParserT__5)|(1<<RuneParserT__6)|(1<<RuneParserT__7)|(1<<RuneParserT__8)|(1<<RuneParserT__12))) != 0 {
			{
				p.SetState(78)

				var _x = p.TypeName2()

				localctx.(*FunctionTypeContext)._typeName2 = _x
			}
			localctx.(*FunctionTypeContext).paramTypes = append(localctx.(*FunctionTypeContext).paramTypes, localctx.(*FunctionTypeContext)._typeName2)
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == RuneParserT__10 {
				{
					p.SetState(79)
					p.Match(RuneParserT__10)
				}
				{
					p.SetState(80)

					var _x = p.TypeName2()

					localctx.(*FunctionTypeContext)._typeName2 = _x
				}
				localctx.(*FunctionTypeContext).paramTypes = append(localctx.(*FunctionTypeContext).paramTypes, localctx.(*FunctionTypeContext)._typeName2)

				p.SetState(85)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(88)
			p.Match(RuneParserT__11)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RuneParserT__2 {
			{
				p.SetState(89)
				p.Match(RuneParserT__2)
			}
			{
				p.SetState(90)

				var _x = p.TypeName2()

				localctx.(*FunctionTypeContext).returnType = _x
			}

		}

	case RuneParserT__12:
		localctx = NewStructTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.Match(RuneParserT__12)
		}
		{
			p.SetState(94)
			p.Match(RuneParserT__13)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RuneParserIDENTIFIER {
			{
				p.SetState(95)
				p.CombinedField()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(RuneParserT__14)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifier returns the identifier token.
	GetIdentifier() antlr.Token

	// SetIdentifier sets the identifier token.
	SetIdentifier(antlr.Token)

	// GetParams returns the params rule contexts.
	GetParams() IParamDeclContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeNameContext

	// GetBody returns the body rule contexts.
	GetBody() IScopeContext

	// SetParams sets the params rule contexts.
	SetParams(IParamDeclContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeNameContext)

	// SetBody sets the body rule contexts.
	SetBody(IScopeContext)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	identifier antlr.Token
	params     IParamDeclContext
	returnType ITypeNameContext
	body       IScopeContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *FunctionContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *FunctionContext) GetParams() IParamDeclContext { return s.params }

func (s *FunctionContext) GetReturnType() ITypeNameContext { return s.returnType }

func (s *FunctionContext) GetBody() IScopeContext { return s.body }

func (s *FunctionContext) SetParams(v IParamDeclContext) { s.params = v }

func (s *FunctionContext) SetReturnType(v ITypeNameContext) { s.returnType = v }

func (s *FunctionContext) SetBody(v IScopeContext) { s.body = v }

func (s *FunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, 0)
}

func (s *FunctionContext) ParamDecl() IParamDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *FunctionContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *FunctionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RuneParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(RuneParserT__8)
	}
	{
		p.SetState(105)

		var _m = p.Match(RuneParserIDENTIFIER)

		localctx.(*FunctionContext).identifier = _m
	}
	{
		p.SetState(106)

		var _x = p.ParamDecl()

		localctx.(*FunctionContext).params = _x
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuneParserT__2 {
		{
			p.SetState(107)
			p.Match(RuneParserT__2)
		}
		{
			p.SetState(108)

			var _x = p.TypeName()

			localctx.(*FunctionContext).returnType = _x
		}

	}
	{
		p.SetState(111)

		var _x = p.Scope()

		localctx.(*FunctionContext).body = _x
	}

	return localctx
}

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_combinedParam returns the _combinedParam rule contexts.
	Get_combinedParam() ICombinedParamContext

	// Set_combinedParam sets the _combinedParam rule contexts.
	Set_combinedParam(ICombinedParamContext)

	// GetParamGroup returns the paramGroup rule context list.
	GetParamGroup() []ICombinedParamContext

	// SetParamGroup sets the paramGroup rule context list.
	SetParamGroup([]ICombinedParamContext)

	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_combinedParam ICombinedParamContext
	paramGroup     []ICombinedParamContext
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_paramDecl
	return p
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) Get_combinedParam() ICombinedParamContext { return s._combinedParam }

func (s *ParamDeclContext) Set_combinedParam(v ICombinedParamContext) { s._combinedParam = v }

func (s *ParamDeclContext) GetParamGroup() []ICombinedParamContext { return s.paramGroup }

func (s *ParamDeclContext) SetParamGroup(v []ICombinedParamContext) { s.paramGroup = v }

func (s *ParamDeclContext) AllCombinedParam() []ICombinedParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICombinedParamContext)(nil)).Elem())
	var tst = make([]ICombinedParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICombinedParamContext)
		}
	}

	return tst
}

func (s *ParamDeclContext) CombinedParam(i int) ICombinedParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICombinedParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICombinedParamContext)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitParamDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RuneParserRULE_paramDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(RuneParserT__9)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuneParserIDENTIFIER {
		{
			p.SetState(114)

			var _x = p.CombinedParam()

			localctx.(*ParamDeclContext)._combinedParam = _x
		}
		localctx.(*ParamDeclContext).paramGroup = append(localctx.(*ParamDeclContext).paramGroup, localctx.(*ParamDeclContext)._combinedParam)
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == RuneParserT__10 {
			{
				p.SetState(115)
				p.Match(RuneParserT__10)
			}
			{
				p.SetState(116)

				var _x = p.CombinedParam()

				localctx.(*ParamDeclContext)._combinedParam = _x
			}
			localctx.(*ParamDeclContext).paramGroup = append(localctx.(*ParamDeclContext).paramGroup, localctx.(*ParamDeclContext)._combinedParam)

			p.SetState(121)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(124)
		p.Match(RuneParserT__11)
	}

	return localctx
}

// ICombinedParamContext is an interface to support dynamic dispatch.
type ICombinedParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// GetNames returns the names token list.
	GetNames() []antlr.Token

	// SetNames sets the names token list.
	SetNames([]antlr.Token)

	// GetParamType returns the paramType rule contexts.
	GetParamType() ITypeNameContext

	// SetParamType sets the paramType rule contexts.
	SetParamType(ITypeNameContext)

	// IsCombinedParamContext differentiates from other interfaces.
	IsCombinedParamContext()
}

type CombinedParamContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_IDENTIFIER antlr.Token
	names       []antlr.Token
	paramType   ITypeNameContext
}

func NewEmptyCombinedParamContext() *CombinedParamContext {
	var p = new(CombinedParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_combinedParam
	return p
}

func (*CombinedParamContext) IsCombinedParamContext() {}

func NewCombinedParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CombinedParamContext {
	var p = new(CombinedParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_combinedParam

	return p
}

func (s *CombinedParamContext) GetParser() antlr.Parser { return s.parser }

func (s *CombinedParamContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *CombinedParamContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *CombinedParamContext) GetNames() []antlr.Token { return s.names }

func (s *CombinedParamContext) SetNames(v []antlr.Token) { s.names = v }

func (s *CombinedParamContext) GetParamType() ITypeNameContext { return s.paramType }

func (s *CombinedParamContext) SetParamType(v ITypeNameContext) { s.paramType = v }

func (s *CombinedParamContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(RuneParserIDENTIFIER)
}

func (s *CombinedParamContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, i)
}

func (s *CombinedParamContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CombinedParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CombinedParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CombinedParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitCombinedParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) CombinedParam() (localctx ICombinedParamContext) {
	localctx = NewCombinedParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RuneParserRULE_combinedParam)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)

		var _m = p.Match(RuneParserIDENTIFIER)

		localctx.(*CombinedParamContext)._IDENTIFIER = _m
	}
	localctx.(*CombinedParamContext).names = append(localctx.(*CombinedParamContext).names, localctx.(*CombinedParamContext)._IDENTIFIER)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuneParserT__10 {
		{
			p.SetState(127)
			p.Match(RuneParserT__10)
		}
		{
			p.SetState(128)

			var _m = p.Match(RuneParserIDENTIFIER)

			localctx.(*CombinedParamContext)._IDENTIFIER = _m
		}
		localctx.(*CombinedParamContext).names = append(localctx.(*CombinedParamContext).names, localctx.(*CombinedParamContext)._IDENTIFIER)

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(134)
		p.Match(RuneParserT__2)
	}
	{
		p.SetState(135)

		var _x = p.TypeName()

		localctx.(*CombinedParamContext).paramType = _x
	}

	return localctx
}

// ICombinedFieldContext is an interface to support dynamic dispatch.
type ICombinedFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IDENTIFIER returns the _IDENTIFIER token.
	Get_IDENTIFIER() antlr.Token

	// Set_IDENTIFIER sets the _IDENTIFIER token.
	Set_IDENTIFIER(antlr.Token)

	// GetNames returns the names token list.
	GetNames() []antlr.Token

	// SetNames sets the names token list.
	SetNames([]antlr.Token)

	// GetParamType returns the paramType rule contexts.
	GetParamType() ITypeNameContext

	// SetParamType sets the paramType rule contexts.
	SetParamType(ITypeNameContext)

	// IsCombinedFieldContext differentiates from other interfaces.
	IsCombinedFieldContext()
}

type CombinedFieldContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_IDENTIFIER antlr.Token
	names       []antlr.Token
	paramType   ITypeNameContext
}

func NewEmptyCombinedFieldContext() *CombinedFieldContext {
	var p = new(CombinedFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_combinedField
	return p
}

func (*CombinedFieldContext) IsCombinedFieldContext() {}

func NewCombinedFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CombinedFieldContext {
	var p = new(CombinedFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_combinedField

	return p
}

func (s *CombinedFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *CombinedFieldContext) Get_IDENTIFIER() antlr.Token { return s._IDENTIFIER }

func (s *CombinedFieldContext) Set_IDENTIFIER(v antlr.Token) { s._IDENTIFIER = v }

func (s *CombinedFieldContext) GetNames() []antlr.Token { return s.names }

func (s *CombinedFieldContext) SetNames(v []antlr.Token) { s.names = v }

func (s *CombinedFieldContext) GetParamType() ITypeNameContext { return s.paramType }

func (s *CombinedFieldContext) SetParamType(v ITypeNameContext) { s.paramType = v }

func (s *CombinedFieldContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(RuneParserIDENTIFIER)
}

func (s *CombinedFieldContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, i)
}

func (s *CombinedFieldContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CombinedFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CombinedFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CombinedFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitCombinedField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) CombinedField() (localctx ICombinedFieldContext) {
	localctx = NewCombinedFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RuneParserRULE_combinedField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)

		var _m = p.Match(RuneParserIDENTIFIER)

		localctx.(*CombinedFieldContext)._IDENTIFIER = _m
	}
	localctx.(*CombinedFieldContext).names = append(localctx.(*CombinedFieldContext).names, localctx.(*CombinedFieldContext)._IDENTIFIER)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RuneParserT__10 {
		{
			p.SetState(138)
			p.Match(RuneParserT__10)
		}
		{
			p.SetState(139)

			var _m = p.Match(RuneParserIDENTIFIER)

			localctx.(*CombinedFieldContext)._IDENTIFIER = _m
		}
		localctx.(*CombinedFieldContext).names = append(localctx.(*CombinedFieldContext).names, localctx.(*CombinedFieldContext)._IDENTIFIER)

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
		p.Match(RuneParserT__2)
	}
	{
		p.SetState(146)

		var _x = p.TypeName()

		localctx.(*CombinedFieldContext).paramType = _x
	}
	{
		p.SetState(147)
		p.Match(RuneParserT__0)
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_statement IStatementContext
	statements []IStatementContext
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) Get_statement() IStatementContext { return s._statement }

func (s *ScopeContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *ScopeContext) GetStatements() []IStatementContext { return s.statements }

func (s *ScopeContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *ScopeContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ScopeContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Scope() (localctx IScopeContext) {
	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RuneParserRULE_scope)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(RuneParserT__13)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__8)|(1<<RuneParserT__9)|(1<<RuneParserT__15)|(1<<RuneParserT__16)|(1<<RuneParserT__18)|(1<<RuneParserT__26))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(RuneParserINTEGER_LITERAL-46))|(1<<(RuneParserREAL_LITERAL-46))|(1<<(RuneParserBOOLEAN_LITERAL-46))|(1<<(RuneParserIDENTIFIER-46)))) != 0) {
		{
			p.SetState(150)

			var _x = p.Statement()

			localctx.(*ScopeContext)._statement = _x
		}
		localctx.(*ScopeContext).statements = append(localctx.(*ScopeContext).statements, localctx.(*ScopeContext)._statement)

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(RuneParserT__14)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRetVal returns the retVal rule contexts.
	GetRetVal() IExpressionContext

	// SetRetVal sets the retVal rule contexts.
	SetRetVal(IExpressionContext)

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	retVal IExpressionContext
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) GetRetVal() IExpressionContext { return s.retVal }

func (s *ReturnStatementContext) SetRetVal(v IExpressionContext) { s.retVal = v }

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RuneParserRULE_returnStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(RuneParserT__15)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__8)|(1<<RuneParserT__9)|(1<<RuneParserT__26))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(RuneParserINTEGER_LITERAL-46))|(1<<(RuneParserREAL_LITERAL-46))|(1<<(RuneParserBOOLEAN_LITERAL-46))|(1<<(RuneParserIDENTIFIER-46)))) != 0) {
		{
			p.SetState(159)

			var _x = p.Expression()

			localctx.(*ReturnStatementContext).retVal = _x
		}

	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_scope returns the _scope rule contexts.
	Get_scope() IScopeContext

	// GetAlternative returns the alternative rule contexts.
	GetAlternative() IScopeContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_scope sets the _scope rule contexts.
	Set_scope(IScopeContext)

	// SetAlternative sets the alternative rule contexts.
	SetAlternative(IScopeContext)

	// GetConditions returns the conditions rule context list.
	GetConditions() []IExpressionContext

	// GetEffects returns the effects rule context list.
	GetEffects() []IScopeContext

	// SetConditions sets the conditions rule context list.
	SetConditions([]IExpressionContext)

	// SetEffects sets the effects rule context list.
	SetEffects([]IScopeContext)

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_expression IExpressionContext
	conditions  []IExpressionContext
	_scope      IScopeContext
	effects     []IScopeContext
	alternative IScopeContext
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) Get_expression() IExpressionContext { return s._expression }

func (s *IfStatementContext) Get_scope() IScopeContext { return s._scope }

func (s *IfStatementContext) GetAlternative() IScopeContext { return s.alternative }

func (s *IfStatementContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *IfStatementContext) Set_scope(v IScopeContext) { s._scope = v }

func (s *IfStatementContext) SetAlternative(v IScopeContext) { s.alternative = v }

func (s *IfStatementContext) GetConditions() []IExpressionContext { return s.conditions }

func (s *IfStatementContext) GetEffects() []IScopeContext { return s.effects }

func (s *IfStatementContext) SetConditions(v []IExpressionContext) { s.conditions = v }

func (s *IfStatementContext) SetEffects(v []IScopeContext) { s.effects = v }

func (s *IfStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllScope() []IScopeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IScopeContext)(nil)).Elem())
	var tst = make([]IScopeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScopeContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Scope(i int) IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RuneParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(RuneParserT__16)
	}
	{
		p.SetState(163)

		var _x = p.Expression()

		localctx.(*IfStatementContext)._expression = _x
	}
	localctx.(*IfStatementContext).conditions = append(localctx.(*IfStatementContext).conditions, localctx.(*IfStatementContext)._expression)
	{
		p.SetState(164)

		var _x = p.Scope()

		localctx.(*IfStatementContext)._scope = _x
	}
	localctx.(*IfStatementContext).effects = append(localctx.(*IfStatementContext).effects, localctx.(*IfStatementContext)._scope)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(165)
				p.Match(RuneParserT__17)
			}
			{
				p.SetState(166)
				p.Match(RuneParserT__16)
			}
			{
				p.SetState(167)

				var _x = p.Expression()

				localctx.(*IfStatementContext)._expression = _x
			}
			localctx.(*IfStatementContext).conditions = append(localctx.(*IfStatementContext).conditions, localctx.(*IfStatementContext)._expression)
			{
				p.SetState(168)

				var _x = p.Scope()

				localctx.(*IfStatementContext)._scope = _x
			}
			localctx.(*IfStatementContext).effects = append(localctx.(*IfStatementContext).effects, localctx.(*IfStatementContext)._scope)

		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuneParserT__17 {
		{
			p.SetState(175)
			p.Match(RuneParserT__17)
		}
		{
			p.SetState(176)

			var _x = p.Scope()

			localctx.(*IfStatementContext).alternative = _x
		}

	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKind returns the kind token.
	GetKind() antlr.Token

	// SetKind sets the kind token.
	SetKind(antlr.Token)

	// GetCondition returns the condition rule contexts.
	GetCondition() IExpressionContext

	// GetBody returns the body rule contexts.
	GetBody() IScopeContext

	// SetCondition sets the condition rule contexts.
	SetCondition(IExpressionContext)

	// SetBody sets the body rule contexts.
	SetBody(IScopeContext)

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	kind      antlr.Token
	condition IExpressionContext
	body      IScopeContext
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) GetKind() antlr.Token { return s.kind }

func (s *LoopContext) SetKind(v antlr.Token) { s.kind = v }

func (s *LoopContext) GetCondition() IExpressionContext { return s.condition }

func (s *LoopContext) GetBody() IScopeContext { return s.body }

func (s *LoopContext) SetCondition(v IExpressionContext) { s.condition = v }

func (s *LoopContext) SetBody(v IScopeContext) { s.body = v }

func (s *LoopContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RuneParserRULE_loop)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(RuneParserT__18)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RuneParserT__19 {
		{
			p.SetState(180)

			var _m = p.Match(RuneParserT__19)

			localctx.(*LoopContext).kind = _m
		}
		{
			p.SetState(181)

			var _x = p.Expression()

			localctx.(*LoopContext).condition = _x
		}

	}
	{
		p.SetState(184)

		var _x = p.Scope()

		localctx.(*LoopContext).body = _x
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RuneParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.expression2(0)
	}

	return localctx
}

// IExpression2Context is an interface to support dynamic dispatch.
type IExpression2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression2Context differentiates from other interfaces.
	IsExpression2Context()
}

type Expression2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression2Context() *Expression2Context {
	var p = new(Expression2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_expression2
	return p
}

func (*Expression2Context) IsExpression2Context() {}

func NewExpression2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression2Context {
	var p = new(Expression2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_expression2

	return p
}

func (s *Expression2Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression2Context) CopyFrom(ctx *Expression2Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Expression2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralPassthroughContext struct {
	*Expression2Context
	value ILiteralContext
}

func NewLiteralPassthroughContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralPassthroughContext {
	var p = new(LiteralPassthroughContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *LiteralPassthroughContext) GetValue() ILiteralContext { return s.value }

func (s *LiteralPassthroughContext) SetValue(v ILiteralContext) { s.value = v }

func (s *LiteralPassthroughContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralPassthroughContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralPassthroughContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitLiteralPassthrough(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentContext struct {
	*Expression2Context
	left  IExpression2Context
	op    IAssignmentOpContext
	right IExpression2Context
}

func NewAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentContext {
	var p = new(AssignmentContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *AssignmentContext) GetLeft() IExpression2Context { return s.left }

func (s *AssignmentContext) GetOp() IAssignmentOpContext { return s.op }

func (s *AssignmentContext) GetRight() IExpression2Context { return s.right }

func (s *AssignmentContext) SetLeft(v IExpression2Context) { s.left = v }

func (s *AssignmentContext) SetOp(v IAssignmentOpContext) { s.op = v }

func (s *AssignmentContext) SetRight(v IExpression2Context) { s.right = v }

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) AllExpression2() []IExpression2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression2Context)(nil)).Elem())
	var tst = make([]IExpression2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression2Context)
		}
	}

	return tst
}

func (s *AssignmentContext) Expression2(i int) IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *AssignmentContext) AssignmentOp() IAssignmentOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOpContext)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpressionContext struct {
	*Expression2Context
	left  IExpression2Context
	op    antlr.Token
	right IExpression2Context
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *BinaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpressionContext) GetLeft() IExpression2Context { return s.left }

func (s *BinaryExpressionContext) GetRight() IExpression2Context { return s.right }

func (s *BinaryExpressionContext) SetLeft(v IExpression2Context) { s.left = v }

func (s *BinaryExpressionContext) SetRight(v IExpression2Context) { s.right = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression2() []IExpression2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression2Context)(nil)).Elem())
	var tst = make([]IExpression2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression2Context)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression2(i int) IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionPassthroughContext struct {
	*Expression2Context
	value IExpression2Context
}

func NewExpressionPassthroughContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionPassthroughContext {
	var p = new(ExpressionPassthroughContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *ExpressionPassthroughContext) GetValue() IExpression2Context { return s.value }

func (s *ExpressionPassthroughContext) SetValue(v IExpression2Context) { s.value = v }

func (s *ExpressionPassthroughContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionPassthroughContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *ExpressionPassthroughContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitExpressionPassthrough(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	*Expression2Context
	op    IUnaryOpContext
	value IExpression2Context
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *UnaryExpressionContext) GetOp() IUnaryOpContext { return s.op }

func (s *UnaryExpressionContext) GetValue() IExpression2Context { return s.value }

func (s *UnaryExpressionContext) SetOp(v IUnaryOpContext) { s.op = v }

func (s *UnaryExpressionContext) SetValue(v IExpression2Context) { s.value = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) UnaryOp() IUnaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOpContext)
}

func (s *UnaryExpressionContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableExpressionContext struct {
	*Expression2Context
	name antlr.Token
}

func NewVariableExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableExpressionContext {
	var p = new(VariableExpressionContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *VariableExpressionContext) GetName() antlr.Token { return s.name }

func (s *VariableExpressionContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, 0)
}

func (s *VariableExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitVariableExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallContext struct {
	*Expression2Context
	name         antlr.Token
	_expression2 IExpression2Context
	params       []IExpression2Context
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *FunctionCallContext) GetName() antlr.Token { return s.name }

func (s *FunctionCallContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionCallContext) Get_expression2() IExpression2Context { return s._expression2 }

func (s *FunctionCallContext) Set_expression2(v IExpression2Context) { s._expression2 = v }

func (s *FunctionCallContext) GetParams() []IExpression2Context { return s.params }

func (s *FunctionCallContext) SetParams(v []IExpression2Context) { s.params = v }

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) AllExpression2() []IExpression2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression2Context)(nil)).Elem())
	var tst = make([]IExpression2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression2Context)
		}
	}

	return tst
}

func (s *FunctionCallContext) Expression2(i int) IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldSelectorContext struct {
	*Expression2Context
	base  IExpression2Context
	field antlr.Token
}

func NewFieldSelectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldSelectorContext {
	var p = new(FieldSelectorContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *FieldSelectorContext) GetField() antlr.Token { return s.field }

func (s *FieldSelectorContext) SetField(v antlr.Token) { s.field = v }

func (s *FieldSelectorContext) GetBase() IExpression2Context { return s.base }

func (s *FieldSelectorContext) SetBase(v IExpression2Context) { s.base = v }

func (s *FieldSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldSelectorContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *FieldSelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RuneParserIDENTIFIER, 0)
}

func (s *FieldSelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitFieldSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaContext struct {
	*Expression2Context
	params     IParamDeclContext
	returnType ITypeNameContext
	body       IScopeContext
}

func NewLambdaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaContext {
	var p = new(LambdaContext)

	p.Expression2Context = NewEmptyExpression2Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Expression2Context))

	return p
}

func (s *LambdaContext) GetParams() IParamDeclContext { return s.params }

func (s *LambdaContext) GetReturnType() ITypeNameContext { return s.returnType }

func (s *LambdaContext) GetBody() IScopeContext { return s.body }

func (s *LambdaContext) SetParams(v IParamDeclContext) { s.params = v }

func (s *LambdaContext) SetReturnType(v ITypeNameContext) { s.returnType = v }

func (s *LambdaContext) SetBody(v IScopeContext) { s.body = v }

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ParamDecl() IParamDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamDeclContext)
}

func (s *LambdaContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *LambdaContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Expression2() (localctx IExpression2Context) {
	return p.expression2(0)
}

func (p *RuneParser) expression2(_p int) (localctx IExpression2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpression2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpression2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, RuneParserRULE_expression2, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionPassthroughContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(189)
			p.Match(RuneParserT__9)
		}
		{
			p.SetState(190)

			var _x = p.expression2(0)

			localctx.(*ExpressionPassthroughContext).value = _x
		}
		{
			p.SetState(191)
			p.Match(RuneParserT__11)
		}

	case 2:
		localctx = NewFunctionCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(193)

			var _m = p.Match(RuneParserIDENTIFIER)

			localctx.(*FunctionCallContext).name = _m
		}
		{
			p.SetState(194)
			p.Match(RuneParserT__9)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__8)|(1<<RuneParserT__9)|(1<<RuneParserT__26))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(RuneParserINTEGER_LITERAL-46))|(1<<(RuneParserREAL_LITERAL-46))|(1<<(RuneParserBOOLEAN_LITERAL-46))|(1<<(RuneParserIDENTIFIER-46)))) != 0) {
			{
				p.SetState(195)

				var _x = p.expression2(0)

				localctx.(*FunctionCallContext)._expression2 = _x
			}
			localctx.(*FunctionCallContext).params = append(localctx.(*FunctionCallContext).params, localctx.(*FunctionCallContext)._expression2)
			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == RuneParserT__10 {
				{
					p.SetState(196)
					p.Match(RuneParserT__10)
				}
				{
					p.SetState(197)

					var _x = p.expression2(0)

					localctx.(*FunctionCallContext)._expression2 = _x
				}
				localctx.(*FunctionCallContext).params = append(localctx.(*FunctionCallContext).params, localctx.(*FunctionCallContext)._expression2)

				p.SetState(202)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(205)
			p.Match(RuneParserT__11)
		}

	case 3:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)

			var _x = p.UnaryOp()

			localctx.(*UnaryExpressionContext).op = _x
		}
		{
			p.SetState(207)

			var _x = p.expression2(10)

			localctx.(*UnaryExpressionContext).value = _x
		}

	case 4:
		localctx = NewLiteralPassthroughContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)

			var _x = p.Literal()

			localctx.(*LiteralPassthroughContext).value = _x
		}

	case 5:
		localctx = NewLambdaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(210)
			p.Match(RuneParserT__8)
		}
		{
			p.SetState(211)

			var _x = p.ParamDecl()

			localctx.(*LambdaContext).params = _x
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RuneParserT__2 {
			{
				p.SetState(212)
				p.Match(RuneParserT__2)
			}
			{
				p.SetState(213)

				var _x = p.TypeName()

				localctx.(*LambdaContext).returnType = _x
			}

		}
		{
			p.SetState(216)

			var _x = p.Scope()

			localctx.(*LambdaContext).body = _x
		}

	case 6:
		localctx = NewVariableExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(218)

			var _m = p.Match(RuneParserIDENTIFIER)

			localctx.(*VariableExpressionContext).name = _m
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(222)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryExpressionContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__21)|(1<<RuneParserT__22)|(1<<RuneParserT__23)|(1<<RuneParserT__24))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryExpressionContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(223)

					var _x = p.expression2(10)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(225)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryExpressionContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<RuneParserT__25)|(1<<RuneParserT__26)|(1<<RuneParserT__27)|(1<<RuneParserT__28))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryExpressionContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(226)

					var _x = p.expression2(9)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(228)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*BinaryExpressionContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(RuneParserT__29-30))|(1<<(RuneParserT__30-30))|(1<<(RuneParserT__31-30))|(1<<(RuneParserT__32-30))|(1<<(RuneParserT__33-30))|(1<<(RuneParserT__34-30)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*BinaryExpressionContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(229)

					var _x = p.expression2(8)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(231)

					var _m = p.Match(RuneParserT__35)

					localctx.(*BinaryExpressionContext).op = _m
				}
				{
					p.SetState(232)

					var _x = p.expression2(7)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 5:
				localctx = NewBinaryExpressionContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(234)

					var _m = p.Match(RuneParserT__36)

					localctx.(*BinaryExpressionContext).op = _m
				}
				{
					p.SetState(235)

					var _x = p.expression2(6)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 6:
				localctx = NewAssignmentContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*AssignmentContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(237)

					var _x = p.AssignmentOp()

					localctx.(*AssignmentContext).op = _x
				}
				{
					p.SetState(238)

					var _x = p.expression2(5)

					localctx.(*AssignmentContext).right = _x
				}

			case 7:
				localctx = NewFieldSelectorContext(p, NewExpression2Context(p, _parentctx, _parentState))
				localctx.(*FieldSelectorContext).base = _prevctx

				p.PushNewRecursionContext(localctx, _startState, RuneParserRULE_expression2)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(241)
					p.Match(RuneParserT__20)
				}
				{
					p.SetState(242)

					var _m = p.Match(RuneParserIDENTIFIER)

					localctx.(*FieldSelectorContext).field = _m
				}

			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryOpContext is an interface to support dynamic dispatch.
type IUnaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOpContext differentiates from other interfaces.
	IsUnaryOpContext()
}

type UnaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOpContext() *UnaryOpContext {
	var p = new(UnaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_unaryOp
	return p
}

func (*UnaryOpContext) IsUnaryOpContext() {}

func NewUnaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOpContext {
	var p = new(UnaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_unaryOp

	return p
}

func (s *UnaryOpContext) GetParser() antlr.Parser { return s.parser }
func (s *UnaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitUnaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) UnaryOp() (localctx IUnaryOpContext) {
	localctx = NewUnaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RuneParserRULE_unaryOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(RuneParserT__26)
	}

	return localctx
}

// IAssignmentOpContext is an interface to support dynamic dispatch.
type IAssignmentOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOpContext differentiates from other interfaces.
	IsAssignmentOpContext()
}

type AssignmentOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOpContext() *AssignmentOpContext {
	var p = new(AssignmentOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_assignmentOp
	return p
}

func (*AssignmentOpContext) IsAssignmentOpContext() {}

func NewAssignmentOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOpContext {
	var p = new(AssignmentOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_assignmentOp

	return p
}

func (s *AssignmentOpContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignmentOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitAssignmentOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) AssignmentOp() (localctx IAssignmentOpContext) {
	localctx = NewAssignmentOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RuneParserRULE_assignmentOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(250)
	_la = p.GetTokenStream().LA(1)

	if !(_la == RuneParserT__3 || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(RuneParserT__37-38))|(1<<(RuneParserT__38-38))|(1<<(RuneParserT__39-38))|(1<<(RuneParserT__40-38))|(1<<(RuneParserT__41-38))|(1<<(RuneParserT__42-38))|(1<<(RuneParserT__43-38))|(1<<(RuneParserT__44-38)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuneParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuneParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RealLiteralContext struct {
	*LiteralContext
	value antlr.Token
}

func NewRealLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *RealLiteralContext) GetValue() antlr.Token { return s.value }

func (s *RealLiteralContext) SetValue(v antlr.Token) { s.value = v }

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(RuneParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitRealLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanLiteralContext struct {
	*LiteralContext
	value antlr.Token
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *BooleanLiteralContext) GetValue() antlr.Token { return s.value }

func (s *BooleanLiteralContext) SetValue(v antlr.Token) { s.value = v }

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) BOOLEAN_LITERAL() antlr.TerminalNode {
	return s.GetToken(RuneParserBOOLEAN_LITERAL, 0)
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegerLiteralContext struct {
	*LiteralContext
	value antlr.Token
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *IntegerLiteralContext) GetValue() antlr.Token { return s.value }

func (s *IntegerLiteralContext) SetValue(v antlr.Token) { s.value = v }

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(RuneParserINTEGER_LITERAL, 0)
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RuneVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RuneParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RuneParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RuneParserREAL_LITERAL:
		localctx = NewRealLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)

			var _m = p.Match(RuneParserREAL_LITERAL)

			localctx.(*RealLiteralContext).value = _m
		}

	case RuneParserINTEGER_LITERAL:
		localctx = NewIntegerLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)

			var _m = p.Match(RuneParserINTEGER_LITERAL)

			localctx.(*IntegerLiteralContext).value = _m
		}

	case RuneParserBOOLEAN_LITERAL:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(254)

			var _m = p.Match(RuneParserBOOLEAN_LITERAL)

			localctx.(*BooleanLiteralContext).value = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *RuneParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *Expression2Context = nil
		if localctx != nil {
			t = localctx.(*Expression2Context)
		}
		return p.Expression2_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RuneParser) Expression2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
