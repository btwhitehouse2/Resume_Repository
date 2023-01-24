package token

type TokenType string
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" //nothing follows
	INT     = "INT"
	LET 	= "LET" //let 
	PRINT 	= "PRINT" //print
	ID 		= "ID" //variables
	EQUALS 	= "EQUALS"
	PLUS    = "PLUS"
	MINUS   = "MINUS"
	DIVIDE  = "DIVIDE"
	ASTERISK = "ASTERISK"
	SCOLON 	= "SEMICOLON"
)

type Token struct {
	Type    TokenType
	Count 	int
	Literal string
	Line    string
}