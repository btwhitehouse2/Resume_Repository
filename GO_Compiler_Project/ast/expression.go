package ast

import (
	st "golite/symboltable"
	"golite/types"
)
type Operator int

const (
	PLUS Operator = iota
	SUB
	GT
	GE
	LT
	LE
	NOT
	ASTERISK
	FSLASH
	COMMA
	OR
	ANDCOMP
	ISEQUAL
	NOTEQUAL
)

func StrToOp(op string) Operator {
	switch op {
	case "+":
		return PLUS
	case "-":
		return SUB
	case ">":
		return GT
	case ">=":
		return GE
	case "<":
		return LT
	case "<=":
		return LE
	case "!":
		return NOT 
	case "*":
		return ASTERISK
	case "/":
		return FSLASH 
	case ",":
		return COMMA
	case "||":
		return OR 
	case "&&":
		return ANDCOMP
	case "==":
		return ISEQUAL
	case "!=":
		return NOTEQUAL
	}
	panic("Error: Could not determine operator")
}
func OpString(op Operator) string {

	switch op {
	case PLUS:
		return "+"
	case SUB:
		return "-"
	case GT:
		return ">"
	case GE:
		return ">="
	case LT:
		return "<"
	case LE:
		return "<="
	case NOT:
		return "!"
	case ASTERISK:
		return "*"
	case FSLASH:
		return "/"
	case COMMA:
		return ","
	case OR:
		return "||"
	case ANDCOMP:
		return "&&"
	case ISEQUAL:
		return "=="
	case NOTEQUAL:
		return "!="
	}
	panic("Error: Could not determine operator")
}

// Expression is the base node for all expression specific nodes (i.e., every expression sub node implements this
// interface)
type Expression interface {
	String() string
	TypeCheck([]string, *st.SymbolTables, string) []string 
	GetType()types.Type 
}
