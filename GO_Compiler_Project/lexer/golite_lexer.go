// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package lexer

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GoliteLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var golitelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func golitelexerLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"NUMBER", "PRINTF", "FMT", "TYPE", "STRUCT", "VAR", "NIL", "RETURN",
		"FOR", "IF", "ELSE", "SCAN", "INT", "FUNC", "BOOL", "TRUE", "FALSE",
		"DELETE", "NEW", "ID", "STRING", "OR", "ANDCOMP", "ISEQUAL", "NOTEQUAL",
		"GREATERTHAN", "LESSTHAN", "GREATEREQUAL", "LESSEQUAL", "NOT", "EQUALS",
		"PLUS", "SUB", "FSLASH", "ASTERISK", "SEMICOLON", "COMMA", "PERIOD",
		"LCURLYBRACE", "RCURLYBRACE", "LPAREN", "RPAREN", "WS", "COMMENTS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 272, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 0, 5, 0, 93, 8, 0, 10, 0,
		12, 0, 96, 9, 0, 3, 0, 98, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 193,
		8, 19, 10, 19, 12, 19, 196, 9, 19, 1, 20, 1, 20, 5, 20, 200, 8, 20, 10,
		20, 12, 20, 203, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41,
		1, 41, 1, 42, 4, 42, 256, 8, 42, 11, 42, 12, 42, 257, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 43, 1, 43, 5, 43, 266, 8, 43, 10, 43, 12, 43, 269, 9, 43,
		1, 43, 1, 43, 1, 201, 0, 44, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69,
		35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87,
		44, 1, 0, 7, 1, 0, 48, 48, 1, 0, 49, 57, 1, 0, 48, 57, 1, 0, 65, 122, 3,
		0, 48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10,
		13, 13, 277, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 1, 97, 1, 0, 0, 0, 3, 99, 1,
		0, 0, 0, 5, 106, 1, 0, 0, 0, 7, 110, 1, 0, 0, 0, 9, 115, 1, 0, 0, 0, 11,
		122, 1, 0, 0, 0, 13, 126, 1, 0, 0, 0, 15, 130, 1, 0, 0, 0, 17, 137, 1,
		0, 0, 0, 19, 141, 1, 0, 0, 0, 21, 144, 1, 0, 0, 0, 23, 149, 1, 0, 0, 0,
		25, 154, 1, 0, 0, 0, 27, 158, 1, 0, 0, 0, 29, 163, 1, 0, 0, 0, 31, 168,
		1, 0, 0, 0, 33, 173, 1, 0, 0, 0, 35, 179, 1, 0, 0, 0, 37, 186, 1, 0, 0,
		0, 39, 190, 1, 0, 0, 0, 41, 197, 1, 0, 0, 0, 43, 206, 1, 0, 0, 0, 45, 209,
		1, 0, 0, 0, 47, 212, 1, 0, 0, 0, 49, 215, 1, 0, 0, 0, 51, 218, 1, 0, 0,
		0, 53, 220, 1, 0, 0, 0, 55, 222, 1, 0, 0, 0, 57, 225, 1, 0, 0, 0, 59, 228,
		1, 0, 0, 0, 61, 230, 1, 0, 0, 0, 63, 232, 1, 0, 0, 0, 65, 234, 1, 0, 0,
		0, 67, 236, 1, 0, 0, 0, 69, 238, 1, 0, 0, 0, 71, 240, 1, 0, 0, 0, 73, 242,
		1, 0, 0, 0, 75, 244, 1, 0, 0, 0, 77, 246, 1, 0, 0, 0, 79, 248, 1, 0, 0,
		0, 81, 250, 1, 0, 0, 0, 83, 252, 1, 0, 0, 0, 85, 255, 1, 0, 0, 0, 87, 261,
		1, 0, 0, 0, 89, 98, 7, 0, 0, 0, 90, 94, 7, 1, 0, 0, 91, 93, 7, 2, 0, 0,
		92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1,
		0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0, 97,
		90, 1, 0, 0, 0, 98, 2, 1, 0, 0, 0, 99, 100, 5, 112, 0, 0, 100, 101, 5,
		114, 0, 0, 101, 102, 5, 105, 0, 0, 102, 103, 5, 110, 0, 0, 103, 104, 5,
		116, 0, 0, 104, 105, 5, 102, 0, 0, 105, 4, 1, 0, 0, 0, 106, 107, 5, 102,
		0, 0, 107, 108, 5, 109, 0, 0, 108, 109, 5, 116, 0, 0, 109, 6, 1, 0, 0,
		0, 110, 111, 5, 116, 0, 0, 111, 112, 5, 121, 0, 0, 112, 113, 5, 112, 0,
		0, 113, 114, 5, 101, 0, 0, 114, 8, 1, 0, 0, 0, 115, 116, 5, 115, 0, 0,
		116, 117, 5, 116, 0, 0, 117, 118, 5, 114, 0, 0, 118, 119, 5, 117, 0, 0,
		119, 120, 5, 99, 0, 0, 120, 121, 5, 116, 0, 0, 121, 10, 1, 0, 0, 0, 122,
		123, 5, 118, 0, 0, 123, 124, 5, 97, 0, 0, 124, 125, 5, 114, 0, 0, 125,
		12, 1, 0, 0, 0, 126, 127, 5, 110, 0, 0, 127, 128, 5, 105, 0, 0, 128, 129,
		5, 108, 0, 0, 129, 14, 1, 0, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5,
		101, 0, 0, 132, 133, 5, 116, 0, 0, 133, 134, 5, 117, 0, 0, 134, 135, 5,
		114, 0, 0, 135, 136, 5, 110, 0, 0, 136, 16, 1, 0, 0, 0, 137, 138, 5, 102,
		0, 0, 138, 139, 5, 111, 0, 0, 139, 140, 5, 114, 0, 0, 140, 18, 1, 0, 0,
		0, 141, 142, 5, 105, 0, 0, 142, 143, 5, 102, 0, 0, 143, 20, 1, 0, 0, 0,
		144, 145, 5, 101, 0, 0, 145, 146, 5, 108, 0, 0, 146, 147, 5, 115, 0, 0,
		147, 148, 5, 101, 0, 0, 148, 22, 1, 0, 0, 0, 149, 150, 5, 115, 0, 0, 150,
		151, 5, 99, 0, 0, 151, 152, 5, 97, 0, 0, 152, 153, 5, 110, 0, 0, 153, 24,
		1, 0, 0, 0, 154, 155, 5, 105, 0, 0, 155, 156, 5, 110, 0, 0, 156, 157, 5,
		116, 0, 0, 157, 26, 1, 0, 0, 0, 158, 159, 5, 102, 0, 0, 159, 160, 5, 117,
		0, 0, 160, 161, 5, 110, 0, 0, 161, 162, 5, 99, 0, 0, 162, 28, 1, 0, 0,
		0, 163, 164, 5, 98, 0, 0, 164, 165, 5, 111, 0, 0, 165, 166, 5, 111, 0,
		0, 166, 167, 5, 108, 0, 0, 167, 30, 1, 0, 0, 0, 168, 169, 5, 116, 0, 0,
		169, 170, 5, 114, 0, 0, 170, 171, 5, 117, 0, 0, 171, 172, 5, 101, 0, 0,
		172, 32, 1, 0, 0, 0, 173, 174, 5, 102, 0, 0, 174, 175, 5, 97, 0, 0, 175,
		176, 5, 108, 0, 0, 176, 177, 5, 115, 0, 0, 177, 178, 5, 101, 0, 0, 178,
		34, 1, 0, 0, 0, 179, 180, 5, 100, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182,
		5, 108, 0, 0, 182, 183, 5, 101, 0, 0, 183, 184, 5, 116, 0, 0, 184, 185,
		5, 101, 0, 0, 185, 36, 1, 0, 0, 0, 186, 187, 5, 110, 0, 0, 187, 188, 5,
		101, 0, 0, 188, 189, 5, 119, 0, 0, 189, 38, 1, 0, 0, 0, 190, 194, 7, 3,
		0, 0, 191, 193, 7, 4, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0,
		194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 40, 1, 0, 0, 0, 196, 194,
		1, 0, 0, 0, 197, 201, 5, 34, 0, 0, 198, 200, 9, 0, 0, 0, 199, 198, 1, 0,
		0, 0, 200, 203, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0,
		202, 204, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 5, 34, 0, 0, 205,
		42, 1, 0, 0, 0, 206, 207, 5, 124, 0, 0, 207, 208, 5, 124, 0, 0, 208, 44,
		1, 0, 0, 0, 209, 210, 5, 38, 0, 0, 210, 211, 5, 38, 0, 0, 211, 46, 1, 0,
		0, 0, 212, 213, 5, 61, 0, 0, 213, 214, 5, 61, 0, 0, 214, 48, 1, 0, 0, 0,
		215, 216, 5, 33, 0, 0, 216, 217, 5, 61, 0, 0, 217, 50, 1, 0, 0, 0, 218,
		219, 5, 62, 0, 0, 219, 52, 1, 0, 0, 0, 220, 221, 5, 60, 0, 0, 221, 54,
		1, 0, 0, 0, 222, 223, 5, 62, 0, 0, 223, 224, 5, 61, 0, 0, 224, 56, 1, 0,
		0, 0, 225, 226, 5, 60, 0, 0, 226, 227, 5, 61, 0, 0, 227, 58, 1, 0, 0, 0,
		228, 229, 5, 33, 0, 0, 229, 60, 1, 0, 0, 0, 230, 231, 5, 61, 0, 0, 231,
		62, 1, 0, 0, 0, 232, 233, 5, 43, 0, 0, 233, 64, 1, 0, 0, 0, 234, 235, 5,
		45, 0, 0, 235, 66, 1, 0, 0, 0, 236, 237, 5, 47, 0, 0, 237, 68, 1, 0, 0,
		0, 238, 239, 5, 42, 0, 0, 239, 70, 1, 0, 0, 0, 240, 241, 5, 59, 0, 0, 241,
		72, 1, 0, 0, 0, 242, 243, 5, 44, 0, 0, 243, 74, 1, 0, 0, 0, 244, 245, 5,
		46, 0, 0, 245, 76, 1, 0, 0, 0, 246, 247, 5, 123, 0, 0, 247, 78, 1, 0, 0,
		0, 248, 249, 5, 125, 0, 0, 249, 80, 1, 0, 0, 0, 250, 251, 5, 40, 0, 0,
		251, 82, 1, 0, 0, 0, 252, 253, 5, 41, 0, 0, 253, 84, 1, 0, 0, 0, 254, 256,
		7, 5, 0, 0, 255, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 255, 1, 0,
		0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 6, 42, 0, 0,
		260, 86, 1, 0, 0, 0, 261, 262, 5, 47, 0, 0, 262, 263, 5, 47, 0, 0, 263,
		267, 1, 0, 0, 0, 264, 266, 8, 6, 0, 0, 265, 264, 1, 0, 0, 0, 266, 269,
		1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 270, 1, 0,
		0, 0, 269, 267, 1, 0, 0, 0, 270, 271, 6, 43, 0, 0, 271, 88, 1, 0, 0, 0,
		7, 0, 94, 97, 194, 201, 257, 267, 1, 6, 0, 0,
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

// GoliteLexerInit initializes any static state used to implement GoliteLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGoliteLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoliteLexerInit() {
	staticData := &golitelexerLexerStaticData
	staticData.once.Do(golitelexerLexerInit)
}

// NewGoliteLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGoliteLexer(input antlr.CharStream) *GoliteLexer {
	GoliteLexerInit()
	l := new(GoliteLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &golitelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "GoliteLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GoliteLexer tokens.
const (
	GoliteLexerNUMBER       = 1
	GoliteLexerPRINTF       = 2
	GoliteLexerFMT          = 3
	GoliteLexerTYPE         = 4
	GoliteLexerSTRUCT       = 5
	GoliteLexerVAR          = 6
	GoliteLexerNIL          = 7
	GoliteLexerRETURN       = 8
	GoliteLexerFOR          = 9
	GoliteLexerIF           = 10
	GoliteLexerELSE         = 11
	GoliteLexerSCAN         = 12
	GoliteLexerINT          = 13
	GoliteLexerFUNC         = 14
	GoliteLexerBOOL         = 15
	GoliteLexerTRUE         = 16
	GoliteLexerFALSE        = 17
	GoliteLexerDELETE       = 18
	GoliteLexerNEW          = 19
	GoliteLexerID           = 20
	GoliteLexerSTRING       = 21
	GoliteLexerOR           = 22
	GoliteLexerANDCOMP      = 23
	GoliteLexerISEQUAL      = 24
	GoliteLexerNOTEQUAL     = 25
	GoliteLexerGREATERTHAN  = 26
	GoliteLexerLESSTHAN     = 27
	GoliteLexerGREATEREQUAL = 28
	GoliteLexerLESSEQUAL    = 29
	GoliteLexerNOT          = 30
	GoliteLexerEQUALS       = 31
	GoliteLexerPLUS         = 32
	GoliteLexerSUB          = 33
	GoliteLexerFSLASH       = 34
	GoliteLexerASTERISK     = 35
	GoliteLexerSEMICOLON    = 36
	GoliteLexerCOMMA        = 37
	GoliteLexerPERIOD       = 38
	GoliteLexerLCURLYBRACE  = 39
	GoliteLexerRCURLYBRACE  = 40
	GoliteLexerLPAREN       = 41
	GoliteLexerRPAREN       = 42
	GoliteLexerWS           = 43
	GoliteLexerCOMMENTS     = 44
)
