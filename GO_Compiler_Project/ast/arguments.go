package ast

import (
	"golite/token"
	st "golite/symboltable"
	"bytes"
)

//don't need to write a string because arguments might be 
//just: (), (expr), (expr,expr,...)
//maybe do only 1 expressions force later term to select additional expression?
type Arguments struct {
	*token.Token
	Args    []Expression // The value assigned to the variable of this statements
}

func (a *Arguments) String() string {

	var out bytes.Buffer
	for _,arg := range a.Args{
		out.WriteString((arg).String())
		out.WriteString("\n")
	}
	out.WriteString(";")
	return out.String()
}

//what account for parenthesis when are those removed
func NewArgStm( args []Expression, token *token.Token) *Arguments {
	return &Arguments{token, args}
}
func (a *Arguments) BuildSymbolTable(tables *st.SymbolTables) {}


//no need to type check
func (a *Arguments) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	for _,arg := range a.Args{
		(arg).TypeCheck(errors,tables,funcId)
	}
	return errors
}