package lexer

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"golite/context"
)

type Lexer interface {
	GetTokenStream() *antlr.CommonTokenStream
	PrintTokens()
	GetErrors() []*context.CompilerError
}

type lexer struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	lexer                       *GoliteLexer
	stream                      *antlr.CommonTokenStream
	errors                      []*context.CompilerError
}

func (lex *lexer) PrintTokens() {
	lex.stream.Fill()
	for _, token := range lex.stream.GetAllTokens() {
		fmt.Println(token.GetText())
	}
}
func (lex *lexer) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	lex.errors = append(lex.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.LEXER,
	})
}

func (lexer *lexer) GetTokenStream() *antlr.CommonTokenStream {

	return lexer.stream
}
func (lexer *lexer) GetErrors() []*context.CompilerError {
	return lexer.errors
}
func NewLexer(inputSourcePath string) Lexer {
	input, _ := antlr.NewFileStream(inputSourcePath)
	lex := &lexer{antlr.NewDefaultErrorListener(), nil, nil, nil}
	antrLexer := NewGoliteLexer(input)
	antrLexer.RemoveErrorListeners()
	antrLexer.AddErrorListener(lex)
	stream := antlr.NewCommonTokenStream(antrLexer, 0)
	lex.lexer = antrLexer
	lex.stream = stream
	return lex
}
