// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoliteParser struct {
	*antlr.BaseParser
}

var goliteparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goliteparserParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.literalNames = []string{
		"", "", "'printf'", "'fmt'", "'type'", "'struct'", "'var'", "'nil'",
		"'return'", "'for'", "'if'", "'else'", "'scan'", "'int'", "'func'",
		"'bool'", "'true'", "'false'", "'delete'", "'new'", "", "", "'||'",
		"'&&'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'!'", "'='",
		"'+'", "'-'", "'/'", "'*'", "';'", "','", "'.'", "'{'", "'}'", "'('",
		"')'",
	}
	staticData.symbolicNames = []string{
		"", "NUMBER", "PRINTF", "FMT", "TYPE", "STRUCT", "VAR", "NIL", "RETURN",
		"FOR", "IF", "ELSE", "SCAN", "INT", "FUNC", "BOOL", "TRUE", "FALSE",
		"DELETE", "NEW", "ID", "STRING", "OR", "ANDCOMP", "ISEQUAL", "NOTEQUAL",
		"GREATERTHAN", "LESSTHAN", "GREATEREQUAL", "LESSEQUAL", "NOT", "EQUALS",
		"PLUS", "SUB", "FSLASH", "ASTERISK", "SEMICOLON", "COMMA", "PERIOD",
		"LCURLYBRACE", "RCURLYBRACE", "LPAREN", "RPAREN", "WS", "COMMENTS",
	}
	staticData.ruleNames = []string{
		"program", "typedeclaration", "fields", "fieldsprime", "decl", "type",
		"declaration", "declarationprime", "function", "funcprime", "parameters",
		"parametersprime", "statement", "block", "delete", "assignment", "assignmentprime",
		"print", "conditional", "conditionalprime", "loop", "rtrn", "rtrnprime",
		"read", "invocation", "arguments", "argumentsprime", "expression", "expressionprime",
		"boolterm", "boolprime", "equalterm", "equalprime", "relationterm",
		"relationprime", "simpleterm", "simpleprime", "term", "termprime", "unaryterm",
		"selectorterm", "selectorprime", "factor",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 44, 379, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 0, 5, 0, 94,
		8, 0, 10, 0, 12, 0, 97, 9, 0, 1, 0, 5, 0, 100, 8, 0, 10, 0, 12, 0, 103,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 5, 2, 118, 8, 2, 10, 2, 12, 2, 121, 9, 2, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1, 6,
		1, 6, 5, 6, 138, 8, 6, 10, 6, 12, 6, 141, 9, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 153, 8, 8, 1, 8, 1, 8, 5, 8,
		157, 8, 8, 10, 8, 12, 8, 160, 9, 8, 1, 8, 5, 8, 163, 8, 8, 10, 8, 12, 8,
		166, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 5, 10, 175, 8,
		10, 10, 10, 12, 10, 178, 9, 10, 3, 10, 180, 8, 10, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 3, 12, 196, 8, 12, 1, 13, 1, 13, 5, 13, 200, 8, 13, 10, 13, 12, 13,
		203, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 5,
		15, 213, 8, 15, 10, 15, 12, 15, 216, 9, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 230, 8,
		17, 10, 17, 12, 17, 233, 9, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 244, 8, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 257, 8, 21, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 5, 25, 274, 8, 25, 10, 25, 12, 25, 277, 9, 25,
		3, 25, 279, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5,
		27, 288, 8, 27, 10, 27, 12, 27, 291, 9, 27, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 5, 29, 298, 8, 29, 10, 29, 12, 29, 301, 9, 29, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 5, 31, 308, 8, 31, 10, 31, 12, 31, 311, 9, 31, 1, 32,
		1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 318, 8, 33, 10, 33, 12, 33, 321, 9,
		33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 328, 8, 35, 10, 35, 12, 35,
		331, 9, 35, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 338, 8, 37, 10, 37,
		12, 37, 341, 9, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 3, 39, 351, 8, 39, 1, 40, 1, 40, 5, 40, 355, 8, 40, 10, 40, 12, 40,
		358, 9, 40, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 3, 42, 369, 8, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42,
		377, 8, 42, 1, 42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 4, 1, 0, 24, 25,
		1, 0, 26, 29, 1, 0, 32, 33, 1, 0, 34, 35, 378, 0, 89, 1, 0, 0, 0, 2, 106,
		1, 0, 0, 0, 4, 114, 1, 0, 0, 0, 6, 122, 1, 0, 0, 0, 8, 125, 1, 0, 0, 0,
		10, 132, 1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 145, 1, 0, 0, 0, 16, 148,
		1, 0, 0, 0, 18, 169, 1, 0, 0, 0, 20, 171, 1, 0, 0, 0, 22, 183, 1, 0, 0,
		0, 24, 195, 1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 206, 1, 0, 0, 0, 30, 210,
		1, 0, 0, 0, 32, 221, 1, 0, 0, 0, 34, 224, 1, 0, 0, 0, 36, 237, 1, 0, 0,
		0, 38, 245, 1, 0, 0, 0, 40, 248, 1, 0, 0, 0, 42, 254, 1, 0, 0, 0, 44, 260,
		1, 0, 0, 0, 46, 262, 1, 0, 0, 0, 48, 266, 1, 0, 0, 0, 50, 270, 1, 0, 0,
		0, 52, 282, 1, 0, 0, 0, 54, 285, 1, 0, 0, 0, 56, 292, 1, 0, 0, 0, 58, 295,
		1, 0, 0, 0, 60, 302, 1, 0, 0, 0, 62, 305, 1, 0, 0, 0, 64, 312, 1, 0, 0,
		0, 66, 315, 1, 0, 0, 0, 68, 322, 1, 0, 0, 0, 70, 325, 1, 0, 0, 0, 72, 332,
		1, 0, 0, 0, 74, 335, 1, 0, 0, 0, 76, 342, 1, 0, 0, 0, 78, 350, 1, 0, 0,
		0, 80, 352, 1, 0, 0, 0, 82, 359, 1, 0, 0, 0, 84, 376, 1, 0, 0, 0, 86, 88,
		3, 2, 1, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 95, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 94, 3,
		12, 6, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95,
		96, 1, 0, 0, 0, 96, 101, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 100, 3, 16,
		8, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101,
		102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105,
		5, 0, 0, 1, 105, 1, 1, 0, 0, 0, 106, 107, 5, 4, 0, 0, 107, 108, 5, 20,
		0, 0, 108, 109, 5, 5, 0, 0, 109, 110, 5, 39, 0, 0, 110, 111, 3, 4, 2, 0,
		111, 112, 5, 40, 0, 0, 112, 113, 5, 36, 0, 0, 113, 3, 1, 0, 0, 0, 114,
		115, 3, 8, 4, 0, 115, 119, 5, 36, 0, 0, 116, 118, 3, 6, 3, 0, 117, 116,
		1, 0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0,
		0, 0, 120, 5, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 3, 8, 4, 0, 123,
		124, 5, 36, 0, 0, 124, 7, 1, 0, 0, 0, 125, 126, 5, 20, 0, 0, 126, 127,
		3, 10, 5, 0, 127, 9, 1, 0, 0, 0, 128, 133, 5, 13, 0, 0, 129, 133, 5, 15,
		0, 0, 130, 131, 5, 35, 0, 0, 131, 133, 5, 20, 0, 0, 132, 128, 1, 0, 0,
		0, 132, 129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133, 11, 1, 0, 0, 0, 134,
		135, 5, 6, 0, 0, 135, 139, 5, 20, 0, 0, 136, 138, 3, 14, 7, 0, 137, 136,
		1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143, 3, 10, 5, 0,
		143, 144, 5, 36, 0, 0, 144, 13, 1, 0, 0, 0, 145, 146, 5, 37, 0, 0, 146,
		147, 5, 20, 0, 0, 147, 15, 1, 0, 0, 0, 148, 149, 5, 14, 0, 0, 149, 150,
		5, 20, 0, 0, 150, 152, 3, 20, 10, 0, 151, 153, 3, 18, 9, 0, 152, 151, 1,
		0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 158, 5, 39, 0,
		0, 155, 157, 3, 12, 6, 0, 156, 155, 1, 0, 0, 0, 157, 160, 1, 0, 0, 0, 158,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 164, 1, 0, 0, 0, 160, 158,
		1, 0, 0, 0, 161, 163, 3, 24, 12, 0, 162, 161, 1, 0, 0, 0, 163, 166, 1,
		0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 167, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 167, 168, 5, 40, 0, 0, 168, 17, 1, 0, 0, 0, 169,
		170, 3, 10, 5, 0, 170, 19, 1, 0, 0, 0, 171, 179, 5, 41, 0, 0, 172, 176,
		3, 8, 4, 0, 173, 175, 3, 22, 11, 0, 174, 173, 1, 0, 0, 0, 175, 178, 1,
		0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 180, 1, 0, 0,
		0, 178, 176, 1, 0, 0, 0, 179, 172, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		181, 1, 0, 0, 0, 181, 182, 5, 42, 0, 0, 182, 21, 1, 0, 0, 0, 183, 184,
		5, 37, 0, 0, 184, 185, 3, 8, 4, 0, 185, 23, 1, 0, 0, 0, 186, 196, 3, 26,
		13, 0, 187, 196, 3, 30, 15, 0, 188, 196, 3, 34, 17, 0, 189, 196, 3, 28,
		14, 0, 190, 196, 3, 36, 18, 0, 191, 196, 3, 40, 20, 0, 192, 196, 3, 42,
		21, 0, 193, 196, 3, 46, 23, 0, 194, 196, 3, 48, 24, 0, 195, 186, 1, 0,
		0, 0, 195, 187, 1, 0, 0, 0, 195, 188, 1, 0, 0, 0, 195, 189, 1, 0, 0, 0,
		195, 190, 1, 0, 0, 0, 195, 191, 1, 0, 0, 0, 195, 192, 1, 0, 0, 0, 195,
		193, 1, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 25, 1, 0, 0, 0, 197, 201, 5,
		39, 0, 0, 198, 200, 3, 24, 12, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0,
		0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 204, 1, 0, 0, 0,
		203, 201, 1, 0, 0, 0, 204, 205, 5, 40, 0, 0, 205, 27, 1, 0, 0, 0, 206,
		207, 5, 18, 0, 0, 207, 208, 3, 54, 27, 0, 208, 209, 5, 36, 0, 0, 209, 29,
		1, 0, 0, 0, 210, 214, 5, 20, 0, 0, 211, 213, 3, 32, 16, 0, 212, 211, 1,
		0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0,
		0, 215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 5, 31, 0, 0, 218,
		219, 3, 54, 27, 0, 219, 220, 5, 36, 0, 0, 220, 31, 1, 0, 0, 0, 221, 222,
		5, 38, 0, 0, 222, 223, 5, 20, 0, 0, 223, 33, 1, 0, 0, 0, 224, 225, 5, 2,
		0, 0, 225, 226, 5, 41, 0, 0, 226, 231, 5, 21, 0, 0, 227, 228, 5, 37, 0,
		0, 228, 230, 3, 54, 27, 0, 229, 227, 1, 0, 0, 0, 230, 233, 1, 0, 0, 0,
		231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 1, 0, 0, 0, 233,
		231, 1, 0, 0, 0, 234, 235, 5, 42, 0, 0, 235, 236, 5, 36, 0, 0, 236, 35,
		1, 0, 0, 0, 237, 238, 5, 10, 0, 0, 238, 239, 5, 41, 0, 0, 239, 240, 3,
		54, 27, 0, 240, 241, 5, 42, 0, 0, 241, 243, 3, 26, 13, 0, 242, 244, 3,
		38, 19, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 37, 1, 0, 0,
		0, 245, 246, 5, 11, 0, 0, 246, 247, 3, 26, 13, 0, 247, 39, 1, 0, 0, 0,
		248, 249, 5, 9, 0, 0, 249, 250, 5, 41, 0, 0, 250, 251, 3, 54, 27, 0, 251,
		252, 5, 42, 0, 0, 252, 253, 3, 26, 13, 0, 253, 41, 1, 0, 0, 0, 254, 256,
		5, 8, 0, 0, 255, 257, 3, 44, 22, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1,
		0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 5, 36, 0, 0, 259, 43, 1, 0, 0,
		0, 260, 261, 3, 54, 27, 0, 261, 45, 1, 0, 0, 0, 262, 263, 5, 12, 0, 0,
		263, 264, 3, 54, 27, 0, 264, 265, 5, 36, 0, 0, 265, 47, 1, 0, 0, 0, 266,
		267, 5, 20, 0, 0, 267, 268, 3, 50, 25, 0, 268, 269, 5, 36, 0, 0, 269, 49,
		1, 0, 0, 0, 270, 278, 5, 41, 0, 0, 271, 275, 3, 54, 27, 0, 272, 274, 3,
		52, 26, 0, 273, 272, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0,
		0, 0, 275, 276, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0,
		278, 271, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		281, 5, 42, 0, 0, 281, 51, 1, 0, 0, 0, 282, 283, 5, 37, 0, 0, 283, 284,
		3, 54, 27, 0, 284, 53, 1, 0, 0, 0, 285, 289, 3, 58, 29, 0, 286, 288, 3,
		56, 28, 0, 287, 286, 1, 0, 0, 0, 288, 291, 1, 0, 0, 0, 289, 287, 1, 0,
		0, 0, 289, 290, 1, 0, 0, 0, 290, 55, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0,
		292, 293, 5, 22, 0, 0, 293, 294, 3, 58, 29, 0, 294, 57, 1, 0, 0, 0, 295,
		299, 3, 62, 31, 0, 296, 298, 3, 60, 30, 0, 297, 296, 1, 0, 0, 0, 298, 301,
		1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 59, 1, 0,
		0, 0, 301, 299, 1, 0, 0, 0, 302, 303, 5, 23, 0, 0, 303, 304, 3, 62, 31,
		0, 304, 61, 1, 0, 0, 0, 305, 309, 3, 66, 33, 0, 306, 308, 3, 64, 32, 0,
		307, 306, 1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309,
		310, 1, 0, 0, 0, 310, 63, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 313, 7,
		0, 0, 0, 313, 314, 3, 66, 33, 0, 314, 65, 1, 0, 0, 0, 315, 319, 3, 70,
		35, 0, 316, 318, 3, 68, 34, 0, 317, 316, 1, 0, 0, 0, 318, 321, 1, 0, 0,
		0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 67, 1, 0, 0, 0, 321,
		319, 1, 0, 0, 0, 322, 323, 7, 1, 0, 0, 323, 324, 3, 70, 35, 0, 324, 69,
		1, 0, 0, 0, 325, 329, 3, 74, 37, 0, 326, 328, 3, 72, 36, 0, 327, 326, 1,
		0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0,
		0, 330, 71, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 332, 333, 7, 2, 0, 0, 333,
		334, 3, 74, 37, 0, 334, 73, 1, 0, 0, 0, 335, 339, 3, 78, 39, 0, 336, 338,
		3, 76, 38, 0, 337, 336, 1, 0, 0, 0, 338, 341, 1, 0, 0, 0, 339, 337, 1,
		0, 0, 0, 339, 340, 1, 0, 0, 0, 340, 75, 1, 0, 0, 0, 341, 339, 1, 0, 0,
		0, 342, 343, 7, 3, 0, 0, 343, 344, 3, 78, 39, 0, 344, 77, 1, 0, 0, 0, 345,
		346, 5, 30, 0, 0, 346, 351, 3, 80, 40, 0, 347, 348, 5, 33, 0, 0, 348, 351,
		3, 80, 40, 0, 349, 351, 3, 80, 40, 0, 350, 345, 1, 0, 0, 0, 350, 347, 1,
		0, 0, 0, 350, 349, 1, 0, 0, 0, 351, 79, 1, 0, 0, 0, 352, 356, 3, 84, 42,
		0, 353, 355, 3, 82, 41, 0, 354, 353, 1, 0, 0, 0, 355, 358, 1, 0, 0, 0,
		356, 354, 1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 81, 1, 0, 0, 0, 358, 356,
		1, 0, 0, 0, 359, 360, 5, 38, 0, 0, 360, 361, 5, 20, 0, 0, 361, 83, 1, 0,
		0, 0, 362, 363, 5, 41, 0, 0, 363, 364, 3, 54, 27, 0, 364, 365, 5, 42, 0,
		0, 365, 377, 1, 0, 0, 0, 366, 368, 5, 20, 0, 0, 367, 369, 3, 50, 25, 0,
		368, 367, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 377, 1, 0, 0, 0, 370,
		377, 5, 1, 0, 0, 371, 372, 5, 19, 0, 0, 372, 377, 5, 20, 0, 0, 373, 377,
		5, 16, 0, 0, 374, 377, 5, 17, 0, 0, 375, 377, 5, 7, 0, 0, 376, 362, 1,
		0, 0, 0, 376, 366, 1, 0, 0, 0, 376, 370, 1, 0, 0, 0, 376, 371, 1, 0, 0,
		0, 376, 373, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0, 0, 0, 377,
		85, 1, 0, 0, 0, 29, 89, 95, 101, 119, 132, 139, 152, 158, 164, 176, 179,
		195, 201, 214, 231, 243, 256, 275, 278, 289, 299, 309, 319, 329, 339, 350,
		356, 368, 376,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoliteParserInit initializes any static state used to implement GoliteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoliteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteParserInit() {
	staticData := &goliteparserParserStaticData
	staticData.once.Do(goliteparserParserInit)
}

// NewGoliteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoliteParser(input antlr.TokenStream) *GoliteParser {
	GoliteParserInit()
	this := new(GoliteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &goliteparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// GoliteParser tokens.
const (
	GoliteParserEOF          = antlr.TokenEOF
	GoliteParserNUMBER       = 1
	GoliteParserPRINTF       = 2
	GoliteParserFMT          = 3
	GoliteParserTYPE         = 4
	GoliteParserSTRUCT       = 5
	GoliteParserVAR          = 6
	GoliteParserNIL          = 7
	GoliteParserRETURN       = 8
	GoliteParserFOR          = 9
	GoliteParserIF           = 10
	GoliteParserELSE         = 11
	GoliteParserSCAN         = 12
	GoliteParserINT          = 13
	GoliteParserFUNC         = 14
	GoliteParserBOOL         = 15
	GoliteParserTRUE         = 16
	GoliteParserFALSE        = 17
	GoliteParserDELETE       = 18
	GoliteParserNEW          = 19
	GoliteParserID           = 20
	GoliteParserSTRING       = 21
	GoliteParserOR           = 22
	GoliteParserANDCOMP      = 23
	GoliteParserISEQUAL      = 24
	GoliteParserNOTEQUAL     = 25
	GoliteParserGREATERTHAN  = 26
	GoliteParserLESSTHAN     = 27
	GoliteParserGREATEREQUAL = 28
	GoliteParserLESSEQUAL    = 29
	GoliteParserNOT          = 30
	GoliteParserEQUALS       = 31
	GoliteParserPLUS         = 32
	GoliteParserSUB          = 33
	GoliteParserFSLASH       = 34
	GoliteParserASTERISK     = 35
	GoliteParserSEMICOLON    = 36
	GoliteParserCOMMA        = 37
	GoliteParserPERIOD       = 38
	GoliteParserLCURLYBRACE  = 39
	GoliteParserRCURLYBRACE  = 40
	GoliteParserLPAREN       = 41
	GoliteParserRPAREN       = 42
	GoliteParserWS           = 43
	GoliteParserCOMMENTS     = 44
)

// GoliteParser rules.
const (
	GoliteParserRULE_program          = 0
	GoliteParserRULE_typedeclaration  = 1
	GoliteParserRULE_fields           = 2
	GoliteParserRULE_fieldsprime      = 3
	GoliteParserRULE_decl             = 4
	GoliteParserRULE_type             = 5
	GoliteParserRULE_declaration      = 6
	GoliteParserRULE_declarationprime = 7
	GoliteParserRULE_function         = 8
	GoliteParserRULE_funcprime        = 9
	GoliteParserRULE_parameters       = 10
	GoliteParserRULE_parametersprime  = 11
	GoliteParserRULE_statement        = 12
	GoliteParserRULE_block            = 13
	GoliteParserRULE_delete           = 14
	GoliteParserRULE_assignment       = 15
	GoliteParserRULE_assignmentprime  = 16
	GoliteParserRULE_print            = 17
	GoliteParserRULE_conditional      = 18
	GoliteParserRULE_conditionalprime = 19
	GoliteParserRULE_loop             = 20
	GoliteParserRULE_rtrn             = 21
	GoliteParserRULE_rtrnprime        = 22
	GoliteParserRULE_read             = 23
	GoliteParserRULE_invocation       = 24
	GoliteParserRULE_arguments        = 25
	GoliteParserRULE_argumentsprime   = 26
	GoliteParserRULE_expression       = 27
	GoliteParserRULE_expressionprime  = 28
	GoliteParserRULE_boolterm         = 29
	GoliteParserRULE_boolprime        = 30
	GoliteParserRULE_equalterm        = 31
	GoliteParserRULE_equalprime       = 32
	GoliteParserRULE_relationterm     = 33
	GoliteParserRULE_relationprime    = 34
	GoliteParserRULE_simpleterm       = 35
	GoliteParserRULE_simpleprime      = 36
	GoliteParserRULE_term             = 37
	GoliteParserRULE_termprime        = 38
	GoliteParserRULE_unaryterm        = 39
	GoliteParserRULE_selectorterm     = 40
	GoliteParserRULE_selectorprime    = 41
	GoliteParserRULE_factor           = 42
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTyDec returns the tyDec rule contexts.
	GetTyDec() ITypedeclarationContext

	// GetDec returns the dec rule contexts.
	GetDec() IDeclarationContext

	// GetFunc_ returns the func_ rule contexts.
	GetFunc_() IFunctionContext

	// SetTyDec sets the tyDec rule contexts.
	SetTyDec(ITypedeclarationContext)

	// SetDec sets the dec rule contexts.
	SetDec(IDeclarationContext)

	// SetFunc_ sets the func_ rule contexts.
	SetFunc_(IFunctionContext)

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tyDec  ITypedeclarationContext
	dec    IDeclarationContext
	func_  IFunctionContext
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) GetTyDec() ITypedeclarationContext { return s.tyDec }

func (s *ProgramContext) GetDec() IDeclarationContext { return s.dec }

func (s *ProgramContext) GetFunc_() IFunctionContext { return s.func_ }

func (s *ProgramContext) SetTyDec(v ITypedeclarationContext) { s.tyDec = v }

func (s *ProgramContext) SetDec(v IDeclarationContext) { s.dec = v }

func (s *ProgramContext) SetFunc_(v IFunctionContext) { s.func_ = v }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoliteParserEOF, 0)
}

func (s *ProgramContext) AllTypedeclaration() []ITypedeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypedeclarationContext); ok {
			len++
		}
	}

	tst := make([]ITypedeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypedeclarationContext); ok {
			tst[i] = t.(ITypedeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Typedeclaration(i int) ITypedeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedeclarationContext)
}

func (s *ProgramContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ProgramContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GoliteParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoliteParserRULE_program)
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserTYPE {
		{
			p.SetState(86)

			var _x = p.Typedeclaration()

			localctx.(*ProgramContext).tyDec = _x
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(92)

			var _x = p.Declaration()

			localctx.(*ProgramContext).dec = _x
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFUNC {
		{
			p.SetState(98)

			var _x = p.Function()

			localctx.(*ProgramContext).func_ = _x
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Match(GoliteParserEOF)
	}

	return localctx
}

// ITypedeclarationContext is an interface to support dynamic dispatch.
type ITypedeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFlds returns the flds rule contexts.
	GetFlds() IFieldsContext

	// SetFlds sets the flds rule contexts.
	SetFlds(IFieldsContext)

	// IsTypedeclarationContext differentiates from other interfaces.
	IsTypedeclarationContext()
}

type TypedeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	flds   IFieldsContext
}

func NewEmptyTypedeclarationContext() *TypedeclarationContext {
	var p = new(TypedeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_typedeclaration
	return p
}

func (*TypedeclarationContext) IsTypedeclarationContext() {}

func NewTypedeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedeclarationContext {
	var p = new(TypedeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_typedeclaration

	return p
}

func (s *TypedeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedeclarationContext) GetFlds() IFieldsContext { return s.flds }

func (s *TypedeclarationContext) SetFlds(v IFieldsContext) { s.flds = v }

func (s *TypedeclarationContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTYPE, 0)
}

func (s *TypedeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypedeclarationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRUCT, 0)
}

func (s *TypedeclarationContext) LCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLCURLYBRACE, 0)
}

func (s *TypedeclarationContext) RCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRCURLYBRACE, 0)
}

func (s *TypedeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *TypedeclarationContext) Fields() IFieldsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsContext)
}

func (s *TypedeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTypedeclaration(s)
	}
}

func (s *TypedeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTypedeclaration(s)
	}
}

func (p *GoliteParser) Typedeclaration() (localctx ITypedeclarationContext) {
	this := p
	_ = this

	localctx = NewTypedeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoliteParserRULE_typedeclaration)

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
		p.SetState(106)
		p.Match(GoliteParserTYPE)
	}
	{
		p.SetState(107)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(108)
		p.Match(GoliteParserSTRUCT)
	}
	{
		p.SetState(109)
		p.Match(GoliteParserLCURLYBRACE)
	}
	{
		p.SetState(110)

		var _x = p.Fields()

		localctx.(*TypedeclarationContext).flds = _x
	}
	{
		p.SetState(111)
		p.Match(GoliteParserRCURLYBRACE)
	}
	{
		p.SetState(112)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IFieldsContext is an interface to support dynamic dispatch.
type IFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDec returns the dec rule contexts.
	GetDec() IDeclContext

	// SetDec sets the dec rule contexts.
	SetDec(IDeclContext)

	// IsFieldsContext differentiates from other interfaces.
	IsFieldsContext()
}

type FieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dec    IDeclContext
}

func NewEmptyFieldsContext() *FieldsContext {
	var p = new(FieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_fields
	return p
}

func (*FieldsContext) IsFieldsContext() {}

func NewFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsContext {
	var p = new(FieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fields

	return p
}

func (s *FieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsContext) GetDec() IDeclContext { return s.dec }

func (s *FieldsContext) SetDec(v IDeclContext) { s.dec = v }

func (s *FieldsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *FieldsContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsContext) AllFieldsprime() []IFieldsprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldsprimeContext); ok {
			len++
		}
	}

	tst := make([]IFieldsprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldsprimeContext); ok {
			tst[i] = t.(IFieldsprimeContext)
			i++
		}
	}

	return tst
}

func (s *FieldsContext) Fieldsprime(i int) IFieldsprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsprimeContext)
}

func (s *FieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFields(s)
	}
}

func (s *FieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFields(s)
	}
}

func (p *GoliteParser) Fields() (localctx IFieldsContext) {
	this := p
	_ = this

	localctx = NewFieldsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoliteParserRULE_fields)
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
		p.SetState(114)

		var _x = p.Decl()

		localctx.(*FieldsContext).dec = _x
	}
	{
		p.SetState(115)
		p.Match(GoliteParserSEMICOLON)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserID {
		{
			p.SetState(116)
			p.Fieldsprime()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldsprimeContext is an interface to support dynamic dispatch.
type IFieldsprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsprimeContext differentiates from other interfaces.
	IsFieldsprimeContext()
}

type FieldsprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsprimeContext() *FieldsprimeContext {
	var p = new(FieldsprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_fieldsprime
	return p
}

func (*FieldsprimeContext) IsFieldsprimeContext() {}

func NewFieldsprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsprimeContext {
	var p = new(FieldsprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_fieldsprime

	return p
}

func (s *FieldsprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsprimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *FieldsprimeContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *FieldsprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFieldsprime(s)
	}
}

func (s *FieldsprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFieldsprime(s)
	}
}

func (p *GoliteParser) Fieldsprime() (localctx IFieldsprimeContext) {
	this := p
	_ = this

	localctx = NewFieldsprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoliteParserRULE_fieldsprime)

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
		p.SetState(122)
		p.Decl()
	}
	{
		p.SetState(123)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetTy returns the ty rule contexts.
	GetTy() ITypeContext

	// SetTy sets the ty rule contexts.
	SetTy(ITypeContext)

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	ty     ITypeContext
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) GetId() antlr.Token { return s.id }

func (s *DeclContext) SetId(v antlr.Token) { s.id = v }

func (s *DeclContext) GetTy() ITypeContext { return s.ty }

func (s *DeclContext) SetTy(v ITypeContext) { s.ty = v }

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *DeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *GoliteParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoliteParserRULE_decl)

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
		p.SetState(125)

		var _m = p.Match(GoliteParserID)

		localctx.(*DeclContext).id = _m
	}
	{
		p.SetState(126)

		var _x = p.Type_()

		localctx.(*DeclContext).ty = _x
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(GoliteParserINT, 0)
}

func (s *TypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GoliteParserBOOL, 0)
}

func (s *TypeContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *GoliteParser) Type_() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoliteParserRULE_type)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(GoliteParserINT)
		}

	case GoliteParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(GoliteParserBOOL)
		}

	case GoliteParserASTERISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(GoliteParserASTERISK)
		}
		{
			p.SetState(131)
			p.Match(GoliteParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVarlit returns the varlit token.
	GetVarlit() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetVarlit sets the varlit token.
	SetVarlit(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	varlit antlr.Token
	id     antlr.Token
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetVarlit() antlr.Token { return s.varlit }

func (s *DeclarationContext) GetId() antlr.Token { return s.id }

func (s *DeclarationContext) SetVarlit(v antlr.Token) { s.varlit = v }

func (s *DeclarationContext) SetId(v antlr.Token) { s.id = v }

func (s *DeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GoliteParserVAR, 0)
}

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *DeclarationContext) AllDeclarationprime() []IDeclarationprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationprimeContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationprimeContext); ok {
			tst[i] = t.(IDeclarationprimeContext)
			i++
		}
	}

	return tst
}

func (s *DeclarationContext) Declarationprime(i int) IDeclarationprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationprimeContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *GoliteParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoliteParserRULE_declaration)
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
		p.SetState(134)

		var _m = p.Match(GoliteParserVAR)

		localctx.(*DeclarationContext).varlit = _m
	}
	{
		p.SetState(135)

		var _m = p.Match(GoliteParserID)

		localctx.(*DeclarationContext).id = _m
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(136)
			p.Declarationprime()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Type_()
	}
	{
		p.SetState(143)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IDeclarationprimeContext is an interface to support dynamic dispatch.
type IDeclarationprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsDeclarationprimeContext differentiates from other interfaces.
	IsDeclarationprimeContext()
}

type DeclarationprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyDeclarationprimeContext() *DeclarationprimeContext {
	var p = new(DeclarationprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_declarationprime
	return p
}

func (*DeclarationprimeContext) IsDeclarationprimeContext() {}

func NewDeclarationprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationprimeContext {
	var p = new(DeclarationprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_declarationprime

	return p
}

func (s *DeclarationprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationprimeContext) GetId() antlr.Token { return s.id }

func (s *DeclarationprimeContext) SetId(v antlr.Token) { s.id = v }

func (s *DeclarationprimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *DeclarationprimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *DeclarationprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDeclarationprime(s)
	}
}

func (s *DeclarationprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDeclarationprime(s)
	}
}

func (p *GoliteParser) Declarationprime() (localctx IDeclarationprimeContext) {
	this := p
	_ = this

	localctx = NewDeclarationprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoliteParserRULE_declarationprime)

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
		p.SetState(145)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(146)

		var _m = p.Match(GoliteParserID)

		localctx.(*DeclarationprimeContext).id = _m
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoliteParserFUNC, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) LCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLCURLYBRACE, 0)
}

func (s *FunctionContext) RCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRCURLYBRACE, 0)
}

func (s *FunctionContext) Funcprime() IFuncprimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncprimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncprimeContext)
}

func (s *FunctionContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *FunctionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *GoliteParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoliteParserRULE_function)
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
		p.SetState(148)
		p.Match(GoliteParserFUNC)
	}
	{
		p.SetState(149)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(150)
		p.Parameters()
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359779328) != 0 {
		{
			p.SetState(151)
			p.Funcprime()
		}

	}
	{
		p.SetState(154)
		p.Match(GoliteParserLCURLYBRACE)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserVAR {
		{
			p.SetState(155)
			p.Declaration()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549757130500) != 0 {
		{
			p.SetState(161)
			p.Statement()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(GoliteParserRCURLYBRACE)
	}

	return localctx
}

// IFuncprimeContext is an interface to support dynamic dispatch.
type IFuncprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncprimeContext differentiates from other interfaces.
	IsFuncprimeContext()
}

type FuncprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncprimeContext() *FuncprimeContext {
	var p = new(FuncprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_funcprime
	return p
}

func (*FuncprimeContext) IsFuncprimeContext() {}

func NewFuncprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncprimeContext {
	var p = new(FuncprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_funcprime

	return p
}

func (s *FuncprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncprimeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FuncprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFuncprime(s)
	}
}

func (s *FuncprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFuncprime(s)
	}
}

func (p *GoliteParser) Funcprime() (localctx IFuncprimeContext) {
	this := p
	_ = this

	localctx = NewFuncprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoliteParserRULE_funcprime)

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
		p.SetState(169)
		p.Type_()
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDec returns the dec rule contexts.
	GetDec() IDeclContext

	// SetDec sets the dec rule contexts.
	SetDec(IDeclContext)

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dec    IDeclContext
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) GetDec() IDeclContext { return s.dec }

func (s *ParametersContext) SetDec(v IDeclContext) { s.dec = v }

func (s *ParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ParametersContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersContext) AllParametersprime() []IParametersprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametersprimeContext); ok {
			len++
		}
	}

	tst := make([]IParametersprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametersprimeContext); ok {
			tst[i] = t.(IParametersprimeContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) Parametersprime(i int) IParametersprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersprimeContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *GoliteParser) Parameters() (localctx IParametersContext) {
	this := p
	_ = this

	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoliteParserRULE_parameters)
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
		p.SetState(171)
		p.Match(GoliteParserLPAREN)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserID {
		{
			p.SetState(172)

			var _x = p.Decl()

			localctx.(*ParametersContext).dec = _x
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(173)
				p.Parametersprime()
			}

			p.SetState(178)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(181)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IParametersprimeContext is an interface to support dynamic dispatch.
type IParametersprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersprimeContext differentiates from other interfaces.
	IsParametersprimeContext()
}

type ParametersprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersprimeContext() *ParametersprimeContext {
	var p = new(ParametersprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_parametersprime
	return p
}

func (*ParametersprimeContext) IsParametersprimeContext() {}

func NewParametersprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersprimeContext {
	var p = new(ParametersprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_parametersprime

	return p
}

func (s *ParametersprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersprimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ParametersprimeContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParametersprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterParametersprime(s)
	}
}

func (s *ParametersprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitParametersprime(s)
	}
}

func (p *GoliteParser) Parametersprime() (localctx IParametersprimeContext) {
	this := p
	_ = this

	localctx = NewParametersprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoliteParserRULE_parametersprime)

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
		p.SetState(183)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(184)
		p.Decl()
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAss returns the ass rule contexts.
	GetAss() IAssignmentContext

	// GetDel returns the del rule contexts.
	GetDel() IDeleteContext

	// GetInv returns the inv rule contexts.
	GetInv() IInvocationContext

	// SetAss sets the ass rule contexts.
	SetAss(IAssignmentContext)

	// SetDel sets the del rule contexts.
	SetDel(IDeleteContext)

	// SetInv sets the inv rule contexts.
	SetInv(IInvocationContext)

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ass    IAssignmentContext
	del    IDeleteContext
	inv    IInvocationContext
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetAss() IAssignmentContext { return s.ass }

func (s *StatementContext) GetDel() IDeleteContext { return s.del }

func (s *StatementContext) GetInv() IInvocationContext { return s.inv }

func (s *StatementContext) SetAss(v IAssignmentContext) { s.ass = v }

func (s *StatementContext) SetDel(v IDeleteContext) { s.del = v }

func (s *StatementContext) SetInv(v IInvocationContext) { s.inv = v }

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) Rtrn() IRtrnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtrnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtrnContext)
}

func (s *StatementContext) Read() IReadContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReadContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReadContext)
}

func (s *StatementContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *GoliteParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoliteParserRULE_statement)

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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)

			var _x = p.Assignment()

			localctx.(*StatementContext).ass = _x
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(188)
			p.Print_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)

			var _x = p.Delete_()

			localctx.(*StatementContext).del = _x
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(190)
			p.Conditional()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(191)
			p.Loop()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(192)
			p.Rtrn()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(193)
			p.Read()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(194)

			var _x = p.Invocation()

			localctx.(*StatementContext).inv = _x
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserLCURLYBRACE, 0)
}

func (s *BlockContext) RCURLYBRACE() antlr.TerminalNode {
	return s.GetToken(GoliteParserRCURLYBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *GoliteParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoliteParserRULE_block)
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
		p.SetState(197)
		p.Match(GoliteParserLCURLYBRACE)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&549757130500) != 0 {
		{
			p.SetState(198)
			p.Statement()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(204)
		p.Match(GoliteParserRCURLYBRACE)
	}

	return localctx
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_delete
	return p
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) GetExpr() IExpressionContext { return s.expr }

func (s *DeleteContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoliteParserDELETE, 0)
}

func (s *DeleteContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *DeleteContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *GoliteParser) Delete_() (localctx IDeleteContext) {
	this := p
	_ = this

	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoliteParserRULE_delete)

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
		p.SetState(206)
		p.Match(GoliteParserDELETE)
	}
	{
		p.SetState(207)

		var _x = p.Expression()

		localctx.(*DeleteContext).expr = _x
	}
	{
		p.SetState(208)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GoliteParserEQUALS, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *AssignmentContext) AllAssignmentprime() []IAssignmentprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentprimeContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentprimeContext); ok {
			tst[i] = t.(IAssignmentprimeContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentContext) Assignmentprime(i int) IAssignmentprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentprimeContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *GoliteParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoliteParserRULE_assignment)
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
		p.SetState(210)
		p.Match(GoliteParserID)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPERIOD {
		{
			p.SetState(211)
			p.Assignmentprime()
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
		p.Match(GoliteParserEQUALS)
	}
	{
		p.SetState(218)
		p.Expression()
	}
	{
		p.SetState(219)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IAssignmentprimeContext is an interface to support dynamic dispatch.
type IAssignmentprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentprimeContext differentiates from other interfaces.
	IsAssignmentprimeContext()
}

type AssignmentprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentprimeContext() *AssignmentprimeContext {
	var p = new(AssignmentprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_assignmentprime
	return p
}

func (*AssignmentprimeContext) IsAssignmentprimeContext() {}

func NewAssignmentprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentprimeContext {
	var p = new(AssignmentprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_assignmentprime

	return p
}

func (s *AssignmentprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentprimeContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(GoliteParserPERIOD, 0)
}

func (s *AssignmentprimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *AssignmentprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterAssignmentprime(s)
	}
}

func (s *AssignmentprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitAssignmentprime(s)
	}
}

func (p *GoliteParser) Assignmentprime() (localctx IAssignmentprimeContext) {
	this := p
	_ = this

	localctx = NewAssignmentprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoliteParserRULE_assignmentprime)

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
		p.SetState(221)
		p.Match(GoliteParserPERIOD)
	}
	{
		p.SetState(222)
		p.Match(GoliteParserID)
	}

	return localctx
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	exp    IExpressionContext
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_print
	return p
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) GetExp() IExpressionContext { return s.exp }

func (s *PrintContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *PrintContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GoliteParserPRINTF, 0)
}

func (s *PrintContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *PrintContext) STRING() antlr.TerminalNode {
	return s.GetToken(GoliteParserSTRING, 0)
}

func (s *PrintContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *PrintContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *PrintContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoliteParserCOMMA)
}

func (s *PrintContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, i)
}

func (s *PrintContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *GoliteParser) Print_() (localctx IPrintContext) {
	this := p
	_ = this

	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoliteParserRULE_print)
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
		p.SetState(224)
		p.Match(GoliteParserPRINTF)
	}
	{
		p.SetState(225)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(226)
		p.Match(GoliteParserSTRING)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserCOMMA {
		{
			p.SetState(227)
			p.Match(GoliteParserCOMMA)
		}
		{
			p.SetState(228)

			var _x = p.Expression()

			localctx.(*PrintContext).exp = _x
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(234)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(235)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond token.
	GetCond() antlr.Token

	// SetCond sets the cond token.
	SetCond(antlr.Token)

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   antlr.Token
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) GetCond() antlr.Token { return s.cond }

func (s *ConditionalContext) SetCond(v antlr.Token) { s.cond = v }

func (s *ConditionalContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ConditionalContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ConditionalContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalContext) IF() antlr.TerminalNode {
	return s.GetToken(GoliteParserIF, 0)
}

func (s *ConditionalContext) Conditionalprime() IConditionalprimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalprimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalprimeContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *GoliteParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoliteParserRULE_conditional)
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
		p.SetState(237)

		var _m = p.Match(GoliteParserIF)

		localctx.(*ConditionalContext).cond = _m
	}
	{
		p.SetState(238)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(239)
		p.Expression()
	}
	{
		p.SetState(240)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(241)
		p.Block()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoliteParserELSE {
		{
			p.SetState(242)
			p.Conditionalprime()
		}

	}

	return localctx
}

// IConditionalprimeContext is an interface to support dynamic dispatch.
type IConditionalprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond token.
	GetCond() antlr.Token

	// SetCond sets the cond token.
	SetCond(antlr.Token)

	// IsConditionalprimeContext differentiates from other interfaces.
	IsConditionalprimeContext()
}

type ConditionalprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   antlr.Token
}

func NewEmptyConditionalprimeContext() *ConditionalprimeContext {
	var p = new(ConditionalprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_conditionalprime
	return p
}

func (*ConditionalprimeContext) IsConditionalprimeContext() {}

func NewConditionalprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalprimeContext {
	var p = new(ConditionalprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_conditionalprime

	return p
}

func (s *ConditionalprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalprimeContext) GetCond() antlr.Token { return s.cond }

func (s *ConditionalprimeContext) SetCond(v antlr.Token) { s.cond = v }

func (s *ConditionalprimeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConditionalprimeContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserELSE, 0)
}

func (s *ConditionalprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterConditionalprime(s)
	}
}

func (s *ConditionalprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitConditionalprime(s)
	}
}

func (p *GoliteParser) Conditionalprime() (localctx IConditionalprimeContext) {
	this := p
	_ = this

	localctx = NewConditionalprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoliteParserRULE_conditionalprime)

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
		p.SetState(245)

		var _m = p.Match(GoliteParserELSE)

		localctx.(*ConditionalprimeContext).cond = _m
	}
	{
		p.SetState(246)
		p.Block()
	}

	return localctx
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLp returns the lp token.
	GetLp() antlr.Token

	// SetLp sets the lp token.
	SetLp(antlr.Token)

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lp     antlr.Token
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_loop
	return p
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) GetLp() antlr.Token { return s.lp }

func (s *LoopContext) SetLp(v antlr.Token) { s.lp = v }

func (s *LoopContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *LoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LoopContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoliteParserFOR, 0)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (p *GoliteParser) Loop() (localctx ILoopContext) {
	this := p
	_ = this

	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoliteParserRULE_loop)

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

		var _m = p.Match(GoliteParserFOR)

		localctx.(*LoopContext).lp = _m
	}
	{
		p.SetState(249)
		p.Match(GoliteParserLPAREN)
	}
	{
		p.SetState(250)
		p.Expression()
	}
	{
		p.SetState(251)
		p.Match(GoliteParserRPAREN)
	}
	{
		p.SetState(252)
		p.Block()
	}

	return localctx
}

// IRtrnContext is an interface to support dynamic dispatch.
type IRtrnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRtrnContext differentiates from other interfaces.
	IsRtrnContext()
}

type RtrnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRtrnContext() *RtrnContext {
	var p = new(RtrnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_rtrn
	return p
}

func (*RtrnContext) IsRtrnContext() {}

func NewRtrnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RtrnContext {
	var p = new(RtrnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_rtrn

	return p
}

func (s *RtrnContext) GetParser() antlr.Parser { return s.parser }

func (s *RtrnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRETURN, 0)
}

func (s *RtrnContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *RtrnContext) Rtrnprime() IRtrnprimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRtrnprimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRtrnprimeContext)
}

func (s *RtrnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RtrnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RtrnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRtrn(s)
	}
}

func (s *RtrnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRtrn(s)
	}
}

func (p *GoliteParser) Rtrn() (localctx IRtrnContext) {
	this := p
	_ = this

	localctx = NewRtrnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoliteParserRULE_rtrn)
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
		p.SetState(254)
		p.Match(GoliteParserRETURN)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2208688701570) != 0 {
		{
			p.SetState(255)
			p.Rtrnprime()
		}

	}
	{
		p.SetState(258)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IRtrnprimeContext is an interface to support dynamic dispatch.
type IRtrnprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRtrnprimeContext differentiates from other interfaces.
	IsRtrnprimeContext()
}

type RtrnprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRtrnprimeContext() *RtrnprimeContext {
	var p = new(RtrnprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_rtrnprime
	return p
}

func (*RtrnprimeContext) IsRtrnprimeContext() {}

func NewRtrnprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RtrnprimeContext {
	var p = new(RtrnprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_rtrnprime

	return p
}

func (s *RtrnprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RtrnprimeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RtrnprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RtrnprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RtrnprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRtrnprime(s)
	}
}

func (s *RtrnprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRtrnprime(s)
	}
}

func (p *GoliteParser) Rtrnprime() (localctx IRtrnprimeContext) {
	this := p
	_ = this

	localctx = NewRtrnprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoliteParserRULE_rtrnprime)

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
		p.SetState(260)
		p.Expression()
	}

	return localctx
}

// IReadContext is an interface to support dynamic dispatch.
type IReadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReadContext differentiates from other interfaces.
	IsReadContext()
}

type ReadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReadContext() *ReadContext {
	var p = new(ReadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_read
	return p
}

func (*ReadContext) IsReadContext() {}

func NewReadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReadContext {
	var p = new(ReadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_read

	return p
}

func (s *ReadContext) GetParser() antlr.Parser { return s.parser }

func (s *ReadContext) SCAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserSCAN, 0)
}

func (s *ReadContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReadContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRead(s)
	}
}

func (p *GoliteParser) Read() (localctx IReadContext) {
	this := p
	_ = this

	localctx = NewReadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoliteParserRULE_read)

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
		p.SetState(262)
		p.Match(GoliteParserSCAN)
	}
	{
		p.SetState(263)
		p.Expression()
	}
	{
		p.SetState(264)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *InvocationContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *InvocationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GoliteParserSEMICOLON, 0)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterInvocation(s)
	}
}

func (s *InvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitInvocation(s)
	}
}

func (p *GoliteParser) Invocation() (localctx IInvocationContext) {
	this := p
	_ = this

	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoliteParserRULE_invocation)

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
		p.SetState(266)
		p.Match(GoliteParserID)
	}
	{
		p.SetState(267)
		p.Arguments()
	}
	{
		p.SetState(268)
		p.Match(GoliteParserSEMICOLON)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *ArgumentsContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) AllArgumentsprime() []IArgumentsprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgumentsprimeContext); ok {
			len++
		}
	}

	tst := make([]IArgumentsprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgumentsprimeContext); ok {
			tst[i] = t.(IArgumentsprimeContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Argumentsprime(i int) IArgumentsprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsprimeContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GoliteParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoliteParserRULE_arguments)
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
		p.SetState(270)
		p.Match(GoliteParserLPAREN)
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2208688701570) != 0 {
		{
			p.SetState(271)
			p.Expression()
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoliteParserCOMMA {
			{
				p.SetState(272)
				p.Argumentsprime()
			}

			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(280)
		p.Match(GoliteParserRPAREN)
	}

	return localctx
}

// IArgumentsprimeContext is an interface to support dynamic dispatch.
type IArgumentsprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsprimeContext differentiates from other interfaces.
	IsArgumentsprimeContext()
}

type ArgumentsprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsprimeContext() *ArgumentsprimeContext {
	var p = new(ArgumentsprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_argumentsprime
	return p
}

func (*ArgumentsprimeContext) IsArgumentsprimeContext() {}

func NewArgumentsprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsprimeContext {
	var p = new(ArgumentsprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_argumentsprime

	return p
}

func (s *ArgumentsprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsprimeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoliteParserCOMMA, 0)
}

func (s *ArgumentsprimeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterArgumentsprime(s)
	}
}

func (s *ArgumentsprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitArgumentsprime(s)
	}
}

func (p *GoliteParser) Argumentsprime() (localctx IArgumentsprimeContext) {
	this := p
	_ = this

	localctx = NewArgumentsprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoliteParserRULE_argumentsprime)

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
		p.SetState(282)
		p.Match(GoliteParserCOMMA)
	}
	{
		p.SetState(283)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBt returns the bt rule contexts.
	GetBt() IBooltermContext

	// SetBt sets the bt rule contexts.
	SetBt(IBooltermContext)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	bt     IBooltermContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetBt() IBooltermContext { return s.bt }

func (s *ExpressionContext) SetBt(v IBooltermContext) { s.bt = v }

func (s *ExpressionContext) Boolterm() IBooltermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooltermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooltermContext)
}

func (s *ExpressionContext) AllExpressionprime() []IExpressionprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionprimeContext); ok {
			len++
		}
	}

	tst := make([]IExpressionprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionprimeContext); ok {
			tst[i] = t.(IExpressionprimeContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expressionprime(i int) IExpressionprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionprimeContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GoliteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoliteParserRULE_expression)
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
		p.SetState(285)

		var _x = p.Boolterm()

		localctx.(*ExpressionContext).bt = _x
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserOR {
		{
			p.SetState(286)
			p.Expressionprime()
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionprimeContext is an interface to support dynamic dispatch.
type IExpressionprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOr returns the or token.
	GetOr() antlr.Token

	// SetOr sets the or token.
	SetOr(antlr.Token)

	// GetTrm returns the trm rule contexts.
	GetTrm() IBooltermContext

	// SetTrm sets the trm rule contexts.
	SetTrm(IBooltermContext)

	// IsExpressionprimeContext differentiates from other interfaces.
	IsExpressionprimeContext()
}

type ExpressionprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	or     antlr.Token
	trm    IBooltermContext
}

func NewEmptyExpressionprimeContext() *ExpressionprimeContext {
	var p = new(ExpressionprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_expressionprime
	return p
}

func (*ExpressionprimeContext) IsExpressionprimeContext() {}

func NewExpressionprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionprimeContext {
	var p = new(ExpressionprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_expressionprime

	return p
}

func (s *ExpressionprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionprimeContext) GetOr() antlr.Token { return s.or }

func (s *ExpressionprimeContext) SetOr(v antlr.Token) { s.or = v }

func (s *ExpressionprimeContext) GetTrm() IBooltermContext { return s.trm }

func (s *ExpressionprimeContext) SetTrm(v IBooltermContext) { s.trm = v }

func (s *ExpressionprimeContext) OR() antlr.TerminalNode {
	return s.GetToken(GoliteParserOR, 0)
}

func (s *ExpressionprimeContext) Boolterm() IBooltermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooltermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooltermContext)
}

func (s *ExpressionprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterExpressionprime(s)
	}
}

func (s *ExpressionprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitExpressionprime(s)
	}
}

func (p *GoliteParser) Expressionprime() (localctx IExpressionprimeContext) {
	this := p
	_ = this

	localctx = NewExpressionprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoliteParserRULE_expressionprime)

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
		p.SetState(292)

		var _m = p.Match(GoliteParserOR)

		localctx.(*ExpressionprimeContext).or = _m
	}
	{
		p.SetState(293)

		var _x = p.Boolterm()

		localctx.(*ExpressionprimeContext).trm = _x
	}

	return localctx
}

// IBooltermContext is an interface to support dynamic dispatch.
type IBooltermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEq returns the eq rule contexts.
	GetEq() IEqualtermContext

	// SetEq sets the eq rule contexts.
	SetEq(IEqualtermContext)

	// IsBooltermContext differentiates from other interfaces.
	IsBooltermContext()
}

type BooltermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	eq     IEqualtermContext
}

func NewEmptyBooltermContext() *BooltermContext {
	var p = new(BooltermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_boolterm
	return p
}

func (*BooltermContext) IsBooltermContext() {}

func NewBooltermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooltermContext {
	var p = new(BooltermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolterm

	return p
}

func (s *BooltermContext) GetParser() antlr.Parser { return s.parser }

func (s *BooltermContext) GetEq() IEqualtermContext { return s.eq }

func (s *BooltermContext) SetEq(v IEqualtermContext) { s.eq = v }

func (s *BooltermContext) Equalterm() IEqualtermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualtermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualtermContext)
}

func (s *BooltermContext) AllBoolprime() []IBoolprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolprimeContext); ok {
			len++
		}
	}

	tst := make([]IBoolprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolprimeContext); ok {
			tst[i] = t.(IBoolprimeContext)
			i++
		}
	}

	return tst
}

func (s *BooltermContext) Boolprime(i int) IBoolprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolprimeContext)
}

func (s *BooltermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooltermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooltermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolterm(s)
	}
}

func (s *BooltermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolterm(s)
	}
}

func (p *GoliteParser) Boolterm() (localctx IBooltermContext) {
	this := p
	_ = this

	localctx = NewBooltermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoliteParserRULE_boolterm)
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
		p.SetState(295)

		var _x = p.Equalterm()

		localctx.(*BooltermContext).eq = _x
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserANDCOMP {
		{
			p.SetState(296)
			p.Boolprime()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBoolprimeContext is an interface to support dynamic dispatch.
type IBoolprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAnd returns the and token.
	GetAnd() antlr.Token

	// SetAnd sets the and token.
	SetAnd(antlr.Token)

	// GetTrm returns the trm rule contexts.
	GetTrm() IEqualtermContext

	// SetTrm sets the trm rule contexts.
	SetTrm(IEqualtermContext)

	// IsBoolprimeContext differentiates from other interfaces.
	IsBoolprimeContext()
}

type BoolprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	and    antlr.Token
	trm    IEqualtermContext
}

func NewEmptyBoolprimeContext() *BoolprimeContext {
	var p = new(BoolprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_boolprime
	return p
}

func (*BoolprimeContext) IsBoolprimeContext() {}

func NewBoolprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolprimeContext {
	var p = new(BoolprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_boolprime

	return p
}

func (s *BoolprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolprimeContext) GetAnd() antlr.Token { return s.and }

func (s *BoolprimeContext) SetAnd(v antlr.Token) { s.and = v }

func (s *BoolprimeContext) GetTrm() IEqualtermContext { return s.trm }

func (s *BoolprimeContext) SetTrm(v IEqualtermContext) { s.trm = v }

func (s *BoolprimeContext) ANDCOMP() antlr.TerminalNode {
	return s.GetToken(GoliteParserANDCOMP, 0)
}

func (s *BoolprimeContext) Equalterm() IEqualtermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualtermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualtermContext)
}

func (s *BoolprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterBoolprime(s)
	}
}

func (s *BoolprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitBoolprime(s)
	}
}

func (p *GoliteParser) Boolprime() (localctx IBoolprimeContext) {
	this := p
	_ = this

	localctx = NewBoolprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoliteParserRULE_boolprime)

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
		p.SetState(302)

		var _m = p.Match(GoliteParserANDCOMP)

		localctx.(*BoolprimeContext).and = _m
	}
	{
		p.SetState(303)

		var _x = p.Equalterm()

		localctx.(*BoolprimeContext).trm = _x
	}

	return localctx
}

// IEqualtermContext is an interface to support dynamic dispatch.
type IEqualtermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRt returns the rt rule contexts.
	GetRt() IRelationtermContext

	// SetRt sets the rt rule contexts.
	SetRt(IRelationtermContext)

	// IsEqualtermContext differentiates from other interfaces.
	IsEqualtermContext()
}

type EqualtermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	rt     IRelationtermContext
}

func NewEmptyEqualtermContext() *EqualtermContext {
	var p = new(EqualtermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_equalterm
	return p
}

func (*EqualtermContext) IsEqualtermContext() {}

func NewEqualtermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualtermContext {
	var p = new(EqualtermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalterm

	return p
}

func (s *EqualtermContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualtermContext) GetRt() IRelationtermContext { return s.rt }

func (s *EqualtermContext) SetRt(v IRelationtermContext) { s.rt = v }

func (s *EqualtermContext) Relationterm() IRelationtermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationtermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationtermContext)
}

func (s *EqualtermContext) AllEqualprime() []IEqualprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualprimeContext); ok {
			len++
		}
	}

	tst := make([]IEqualprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualprimeContext); ok {
			tst[i] = t.(IEqualprimeContext)
			i++
		}
	}

	return tst
}

func (s *EqualtermContext) Equalprime(i int) IEqualprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualprimeContext)
}

func (s *EqualtermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualtermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualtermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualterm(s)
	}
}

func (s *EqualtermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualterm(s)
	}
}

func (p *GoliteParser) Equalterm() (localctx IEqualtermContext) {
	this := p
	_ = this

	localctx = NewEqualtermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoliteParserRULE_equalterm)
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
		p.SetState(305)

		var _x = p.Relationterm()

		localctx.(*EqualtermContext).rt = _x
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserISEQUAL || _la == GoliteParserNOTEQUAL {
		{
			p.SetState(306)
			p.Equalprime()
		}

		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEqualprimeContext is an interface to support dynamic dispatch.
type IEqualprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEq returns the eq token.
	GetEq() antlr.Token

	// SetEq sets the eq token.
	SetEq(antlr.Token)

	// GetTrm returns the trm rule contexts.
	GetTrm() IRelationtermContext

	// SetTrm sets the trm rule contexts.
	SetTrm(IRelationtermContext)

	// IsEqualprimeContext differentiates from other interfaces.
	IsEqualprimeContext()
}

type EqualprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	eq     antlr.Token
	trm    IRelationtermContext
}

func NewEmptyEqualprimeContext() *EqualprimeContext {
	var p = new(EqualprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_equalprime
	return p
}

func (*EqualprimeContext) IsEqualprimeContext() {}

func NewEqualprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualprimeContext {
	var p = new(EqualprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_equalprime

	return p
}

func (s *EqualprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualprimeContext) GetEq() antlr.Token { return s.eq }

func (s *EqualprimeContext) SetEq(v antlr.Token) { s.eq = v }

func (s *EqualprimeContext) GetTrm() IRelationtermContext { return s.trm }

func (s *EqualprimeContext) SetTrm(v IRelationtermContext) { s.trm = v }

func (s *EqualprimeContext) Relationterm() IRelationtermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationtermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationtermContext)
}

func (s *EqualprimeContext) ISEQUAL() antlr.TerminalNode {
	return s.GetToken(GoliteParserISEQUAL, 0)
}

func (s *EqualprimeContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(GoliteParserNOTEQUAL, 0)
}

func (s *EqualprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterEqualprime(s)
	}
}

func (s *EqualprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitEqualprime(s)
	}
}

func (p *GoliteParser) Equalprime() (localctx IEqualprimeContext) {
	this := p
	_ = this

	localctx = NewEqualprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoliteParserRULE_equalprime)
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
		p.SetState(312)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*EqualprimeContext).eq = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserISEQUAL || _la == GoliteParserNOTEQUAL) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*EqualprimeContext).eq = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(313)

		var _x = p.Relationterm()

		localctx.(*EqualprimeContext).trm = _x
	}

	return localctx
}

// IRelationtermContext is an interface to support dynamic dispatch.
type IRelationtermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSt returns the st rule contexts.
	GetSt() ISimpletermContext

	// SetSt sets the st rule contexts.
	SetSt(ISimpletermContext)

	// IsRelationtermContext differentiates from other interfaces.
	IsRelationtermContext()
}

type RelationtermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	st     ISimpletermContext
}

func NewEmptyRelationtermContext() *RelationtermContext {
	var p = new(RelationtermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_relationterm
	return p
}

func (*RelationtermContext) IsRelationtermContext() {}

func NewRelationtermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationtermContext {
	var p = new(RelationtermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationterm

	return p
}

func (s *RelationtermContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationtermContext) GetSt() ISimpletermContext { return s.st }

func (s *RelationtermContext) SetSt(v ISimpletermContext) { s.st = v }

func (s *RelationtermContext) Simpleterm() ISimpletermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpletermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpletermContext)
}

func (s *RelationtermContext) AllRelationprime() []IRelationprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationprimeContext); ok {
			len++
		}
	}

	tst := make([]IRelationprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationprimeContext); ok {
			tst[i] = t.(IRelationprimeContext)
			i++
		}
	}

	return tst
}

func (s *RelationtermContext) Relationprime(i int) IRelationprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationprimeContext)
}

func (s *RelationtermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationtermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationtermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationterm(s)
	}
}

func (s *RelationtermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationterm(s)
	}
}

func (p *GoliteParser) Relationterm() (localctx IRelationtermContext) {
	this := p
	_ = this

	localctx = NewRelationtermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoliteParserRULE_relationterm)
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
		p.SetState(315)

		var _x = p.Simpleterm()

		localctx.(*RelationtermContext).st = _x
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0 {
		{
			p.SetState(316)
			p.Relationprime()
		}

		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRelationprimeContext is an interface to support dynamic dispatch.
type IRelationprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComp returns the comp token.
	GetComp() antlr.Token

	// SetComp sets the comp token.
	SetComp(antlr.Token)

	// GetTrm returns the trm rule contexts.
	GetTrm() ISimpletermContext

	// SetTrm sets the trm rule contexts.
	SetTrm(ISimpletermContext)

	// IsRelationprimeContext differentiates from other interfaces.
	IsRelationprimeContext()
}

type RelationprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	comp   antlr.Token
	trm    ISimpletermContext
}

func NewEmptyRelationprimeContext() *RelationprimeContext {
	var p = new(RelationprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_relationprime
	return p
}

func (*RelationprimeContext) IsRelationprimeContext() {}

func NewRelationprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationprimeContext {
	var p = new(RelationprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_relationprime

	return p
}

func (s *RelationprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationprimeContext) GetComp() antlr.Token { return s.comp }

func (s *RelationprimeContext) SetComp(v antlr.Token) { s.comp = v }

func (s *RelationprimeContext) GetTrm() ISimpletermContext { return s.trm }

func (s *RelationprimeContext) SetTrm(v ISimpletermContext) { s.trm = v }

func (s *RelationprimeContext) Simpleterm() ISimpletermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpletermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpletermContext)
}

func (s *RelationprimeContext) GREATERTHAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserGREATERTHAN, 0)
}

func (s *RelationprimeContext) LESSTHAN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLESSTHAN, 0)
}

func (s *RelationprimeContext) LESSEQUAL() antlr.TerminalNode {
	return s.GetToken(GoliteParserLESSEQUAL, 0)
}

func (s *RelationprimeContext) GREATEREQUAL() antlr.TerminalNode {
	return s.GetToken(GoliteParserGREATEREQUAL, 0)
}

func (s *RelationprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterRelationprime(s)
	}
}

func (s *RelationprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitRelationprime(s)
	}
}

func (p *GoliteParser) Relationprime() (localctx IRelationprimeContext) {
	this := p
	_ = this

	localctx = NewRelationprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoliteParserRULE_relationprime)
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
		p.SetState(322)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*RelationprimeContext).comp = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*RelationprimeContext).comp = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(323)

		var _x = p.Simpleterm()

		localctx.(*RelationprimeContext).trm = _x
	}

	return localctx
}

// ISimpletermContext is an interface to support dynamic dispatch.
type ISimpletermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t rule contexts.
	GetT() ITermContext

	// SetT sets the t rule contexts.
	SetT(ITermContext)

	// IsSimpletermContext differentiates from other interfaces.
	IsSimpletermContext()
}

type SimpletermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      ITermContext
}

func NewEmptySimpletermContext() *SimpletermContext {
	var p = new(SimpletermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleterm
	return p
}

func (*SimpletermContext) IsSimpletermContext() {}

func NewSimpletermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpletermContext {
	var p = new(SimpletermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleterm

	return p
}

func (s *SimpletermContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpletermContext) GetT() ITermContext { return s.t }

func (s *SimpletermContext) SetT(v ITermContext) { s.t = v }

func (s *SimpletermContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpletermContext) AllSimpleprime() []ISimpleprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleprimeContext); ok {
			len++
		}
	}

	tst := make([]ISimpleprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleprimeContext); ok {
			tst[i] = t.(ISimpleprimeContext)
			i++
		}
	}

	return tst
}

func (s *SimpletermContext) Simpleprime(i int) ISimpleprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleprimeContext)
}

func (s *SimpletermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpletermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpletermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleterm(s)
	}
}

func (s *SimpletermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleterm(s)
	}
}

func (p *GoliteParser) Simpleterm() (localctx ISimpletermContext) {
	this := p
	_ = this

	localctx = NewSimpletermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoliteParserRULE_simpleterm)
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
		p.SetState(325)

		var _x = p.Term()

		localctx.(*SimpletermContext).t = _x
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPLUS || _la == GoliteParserSUB {
		{
			p.SetState(326)
			p.Simpleprime()
		}

		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISimpleprimeContext is an interface to support dynamic dispatch.
type ISimpleprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPm returns the pm token.
	GetPm() antlr.Token

	// SetPm sets the pm token.
	SetPm(antlr.Token)

	// GetTrm returns the trm rule contexts.
	GetTrm() ITermContext

	// SetTrm sets the trm rule contexts.
	SetTrm(ITermContext)

	// IsSimpleprimeContext differentiates from other interfaces.
	IsSimpleprimeContext()
}

type SimpleprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	pm     antlr.Token
	trm    ITermContext
}

func NewEmptySimpleprimeContext() *SimpleprimeContext {
	var p = new(SimpleprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_simpleprime
	return p
}

func (*SimpleprimeContext) IsSimpleprimeContext() {}

func NewSimpleprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleprimeContext {
	var p = new(SimpleprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_simpleprime

	return p
}

func (s *SimpleprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleprimeContext) GetPm() antlr.Token { return s.pm }

func (s *SimpleprimeContext) SetPm(v antlr.Token) { s.pm = v }

func (s *SimpleprimeContext) GetTrm() ITermContext { return s.trm }

func (s *SimpleprimeContext) SetTrm(v ITermContext) { s.trm = v }

func (s *SimpleprimeContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *SimpleprimeContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GoliteParserPLUS, 0)
}

func (s *SimpleprimeContext) SUB() antlr.TerminalNode {
	return s.GetToken(GoliteParserSUB, 0)
}

func (s *SimpleprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSimpleprime(s)
	}
}

func (s *SimpleprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSimpleprime(s)
	}
}

func (p *GoliteParser) Simpleprime() (localctx ISimpleprimeContext) {
	this := p
	_ = this

	localctx = NewSimpleprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoliteParserRULE_simpleprime)
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
		p.SetState(332)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*SimpleprimeContext).pm = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserPLUS || _la == GoliteParserSUB) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*SimpleprimeContext).pm = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(333)

		var _x = p.Term()

		localctx.(*SimpleprimeContext).trm = _x
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUt returns the ut rule contexts.
	GetUt() IUnarytermContext

	// SetUt sets the ut rule contexts.
	SetUt(IUnarytermContext)

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ut     IUnarytermContext
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) GetUt() IUnarytermContext { return s.ut }

func (s *TermContext) SetUt(v IUnarytermContext) { s.ut = v }

func (s *TermContext) Unaryterm() IUnarytermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarytermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarytermContext)
}

func (s *TermContext) AllTermprime() []ITermprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermprimeContext); ok {
			len++
		}
	}

	tst := make([]ITermprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermprimeContext); ok {
			tst[i] = t.(ITermprimeContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Termprime(i int) ITermprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermprimeContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *GoliteParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoliteParserRULE_term)
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
		p.SetState(335)

		var _x = p.Unaryterm()

		localctx.(*TermContext).ut = _x
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserFSLASH || _la == GoliteParserASTERISK {
		{
			p.SetState(336)
			p.Termprime()
		}

		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermprimeContext is an interface to support dynamic dispatch.
type ITermprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAf returns the af token.
	GetAf() antlr.Token

	// SetAf sets the af token.
	SetAf(antlr.Token)

	// GetTrm returns the trm rule contexts.
	GetTrm() IUnarytermContext

	// SetTrm sets the trm rule contexts.
	SetTrm(IUnarytermContext)

	// IsTermprimeContext differentiates from other interfaces.
	IsTermprimeContext()
}

type TermprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	af     antlr.Token
	trm    IUnarytermContext
}

func NewEmptyTermprimeContext() *TermprimeContext {
	var p = new(TermprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_termprime
	return p
}

func (*TermprimeContext) IsTermprimeContext() {}

func NewTermprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermprimeContext {
	var p = new(TermprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_termprime

	return p
}

func (s *TermprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TermprimeContext) GetAf() antlr.Token { return s.af }

func (s *TermprimeContext) SetAf(v antlr.Token) { s.af = v }

func (s *TermprimeContext) GetTrm() IUnarytermContext { return s.trm }

func (s *TermprimeContext) SetTrm(v IUnarytermContext) { s.trm = v }

func (s *TermprimeContext) Unaryterm() IUnarytermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarytermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarytermContext)
}

func (s *TermprimeContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(GoliteParserASTERISK, 0)
}

func (s *TermprimeContext) FSLASH() antlr.TerminalNode {
	return s.GetToken(GoliteParserFSLASH, 0)
}

func (s *TermprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterTermprime(s)
	}
}

func (s *TermprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitTermprime(s)
	}
}

func (p *GoliteParser) Termprime() (localctx ITermprimeContext) {
	this := p
	_ = this

	localctx = NewTermprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoliteParserRULE_termprime)
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
		p.SetState(342)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*TermprimeContext).af = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == GoliteParserFSLASH || _la == GoliteParserASTERISK) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*TermprimeContext).af = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(343)

		var _x = p.Unaryterm()

		localctx.(*TermprimeContext).trm = _x
	}

	return localctx
}

// IUnarytermContext is an interface to support dynamic dispatch.
type IUnarytermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNot returns the not token.
	GetNot() antlr.Token

	// GetNeg returns the neg token.
	GetNeg() antlr.Token

	// SetNot sets the not token.
	SetNot(antlr.Token)

	// SetNeg sets the neg token.
	SetNeg(antlr.Token)

	// GetSt returns the st rule contexts.
	GetSt() ISelectortermContext

	// SetSt sets the st rule contexts.
	SetSt(ISelectortermContext)

	// IsUnarytermContext differentiates from other interfaces.
	IsUnarytermContext()
}

type UnarytermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	not    antlr.Token
	st     ISelectortermContext
	neg    antlr.Token
}

func NewEmptyUnarytermContext() *UnarytermContext {
	var p = new(UnarytermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_unaryterm
	return p
}

func (*UnarytermContext) IsUnarytermContext() {}

func NewUnarytermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnarytermContext {
	var p = new(UnarytermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_unaryterm

	return p
}

func (s *UnarytermContext) GetParser() antlr.Parser { return s.parser }

func (s *UnarytermContext) GetNot() antlr.Token { return s.not }

func (s *UnarytermContext) GetNeg() antlr.Token { return s.neg }

func (s *UnarytermContext) SetNot(v antlr.Token) { s.not = v }

func (s *UnarytermContext) SetNeg(v antlr.Token) { s.neg = v }

func (s *UnarytermContext) GetSt() ISelectortermContext { return s.st }

func (s *UnarytermContext) SetSt(v ISelectortermContext) { s.st = v }

func (s *UnarytermContext) NOT() antlr.TerminalNode {
	return s.GetToken(GoliteParserNOT, 0)
}

func (s *UnarytermContext) Selectorterm() ISelectortermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectortermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectortermContext)
}

func (s *UnarytermContext) SUB() antlr.TerminalNode {
	return s.GetToken(GoliteParserSUB, 0)
}

func (s *UnarytermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarytermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnarytermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterUnaryterm(s)
	}
}

func (s *UnarytermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitUnaryterm(s)
	}
}

func (p *GoliteParser) Unaryterm() (localctx IUnarytermContext) {
	this := p
	_ = this

	localctx = NewUnarytermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoliteParserRULE_unaryterm)

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

	p.SetState(350)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)

			var _m = p.Match(GoliteParserNOT)

			localctx.(*UnarytermContext).not = _m
		}
		{
			p.SetState(346)

			var _x = p.Selectorterm()

			localctx.(*UnarytermContext).st = _x
		}

	case GoliteParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)

			var _m = p.Match(GoliteParserSUB)

			localctx.(*UnarytermContext).neg = _m
		}
		{
			p.SetState(348)
			p.Selectorterm()
		}

	case GoliteParserNUMBER, GoliteParserNIL, GoliteParserTRUE, GoliteParserFALSE, GoliteParserNEW, GoliteParserID, GoliteParserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(349)
			p.Selectorterm()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectortermContext is an interface to support dynamic dispatch.
type ISelectortermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f rule contexts.
	GetF() IFactorContext

	// SetF sets the f rule contexts.
	SetF(IFactorContext)

	// IsSelectortermContext differentiates from other interfaces.
	IsSelectortermContext()
}

type SelectortermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	f      IFactorContext
}

func NewEmptySelectortermContext() *SelectortermContext {
	var p = new(SelectortermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorterm
	return p
}

func (*SelectortermContext) IsSelectortermContext() {}

func NewSelectortermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectortermContext {
	var p = new(SelectortermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorterm

	return p
}

func (s *SelectortermContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectortermContext) GetF() IFactorContext { return s.f }

func (s *SelectortermContext) SetF(v IFactorContext) { s.f = v }

func (s *SelectortermContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *SelectortermContext) AllSelectorprime() []ISelectorprimeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectorprimeContext); ok {
			len++
		}
	}

	tst := make([]ISelectorprimeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectorprimeContext); ok {
			tst[i] = t.(ISelectorprimeContext)
			i++
		}
	}

	return tst
}

func (s *SelectortermContext) Selectorprime(i int) ISelectorprimeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorprimeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorprimeContext)
}

func (s *SelectortermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectortermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectortermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorterm(s)
	}
}

func (s *SelectortermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorterm(s)
	}
}

func (p *GoliteParser) Selectorterm() (localctx ISelectortermContext) {
	this := p
	_ = this

	localctx = NewSelectortermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoliteParserRULE_selectorterm)
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
		p.SetState(352)

		var _x = p.Factor()

		localctx.(*SelectortermContext).f = _x
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoliteParserPERIOD {
		{
			p.SetState(353)
			p.Selectorprime()
		}

		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectorprimeContext is an interface to support dynamic dispatch.
type ISelectorprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsSelectorprimeContext differentiates from other interfaces.
	IsSelectorprimeContext()
}

type SelectorprimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptySelectorprimeContext() *SelectorprimeContext {
	var p = new(SelectorprimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_selectorprime
	return p
}

func (*SelectorprimeContext) IsSelectorprimeContext() {}

func NewSelectorprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorprimeContext {
	var p = new(SelectorprimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_selectorprime

	return p
}

func (s *SelectorprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorprimeContext) GetId() antlr.Token { return s.id }

func (s *SelectorprimeContext) SetId(v antlr.Token) { s.id = v }

func (s *SelectorprimeContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(GoliteParserPERIOD, 0)
}

func (s *SelectorprimeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *SelectorprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterSelectorprime(s)
	}
}

func (s *SelectorprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitSelectorprime(s)
	}
}

func (p *GoliteParser) Selectorprime() (localctx ISelectorprimeContext) {
	this := p
	_ = this

	localctx = NewSelectorprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GoliteParserRULE_selectorprime)

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
		p.SetState(359)
		p.Match(GoliteParserPERIOD)
	}
	{
		p.SetState(360)

		var _m = p.Match(GoliteParserID)

		localctx.(*SelectorprimeContext).id = _m
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// GetArg returns the arg rule contexts.
	GetArg() IArgumentsContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// SetArg sets the arg rule contexts.
	SetArg(IArgumentsContext)

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
	arg    IArgumentsContext
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoliteParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoliteParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) GetExpr() IExpressionContext { return s.expr }

func (s *FactorContext) GetArg() IArgumentsContext { return s.arg }

func (s *FactorContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *FactorContext) SetArg(v IArgumentsContext) { s.arg = v }

func (s *FactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserLPAREN, 0)
}

func (s *FactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GoliteParserRPAREN, 0)
}

func (s *FactorContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(GoliteParserID, 0)
}

func (s *FactorContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FactorContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GoliteParserNUMBER, 0)
}

func (s *FactorContext) NEW() antlr.TerminalNode {
	return s.GetToken(GoliteParserNEW, 0)
}

func (s *FactorContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GoliteParserTRUE, 0)
}

func (s *FactorContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GoliteParserFALSE, 0)
}

func (s *FactorContext) NIL() antlr.TerminalNode {
	return s.GetToken(GoliteParserNIL, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoliteParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *GoliteParser) Factor() (localctx IFactorContext) {
	this := p
	_ = this

	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GoliteParserRULE_factor)
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

	p.SetState(376)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoliteParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.Match(GoliteParserLPAREN)
		}
		{
			p.SetState(363)

			var _x = p.Expression()

			localctx.(*FactorContext).expr = _x
		}
		{
			p.SetState(364)
			p.Match(GoliteParserRPAREN)
		}

	case GoliteParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.Match(GoliteParserID)
		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoliteParserLPAREN {
			{
				p.SetState(367)

				var _x = p.Arguments()

				localctx.(*FactorContext).arg = _x
			}

		}

	case GoliteParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(370)
			p.Match(GoliteParserNUMBER)
		}

	case GoliteParserNEW:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(371)
			p.Match(GoliteParserNEW)
		}
		{
			p.SetState(372)
			p.Match(GoliteParserID)
		}

	case GoliteParserTRUE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(373)
			p.Match(GoliteParserTRUE)
		}

	case GoliteParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(374)
			p.Match(GoliteParserFALSE)
		}

	case GoliteParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(375)
			p.Match(GoliteParserNIL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
