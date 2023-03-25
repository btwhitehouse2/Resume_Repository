package ast

import (
	"bytes"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)

//printf("c=%d\nd=%d",c,d); 

type Print struct {
	*token.Token            // The token for this print statement
	printString		string	//the print formatting
	exprs        	[]Expression // The variables being printed out.
}

func NewPrintStm(exprs []Expression, printLit string, token *token.Token) *Print {
	return &Print{token, printLit, exprs}
}
func (p *Print) String() string {

	var out bytes.Buffer

	out.WriteString(p.printString)
	out.WriteString(" ")
	for _,e := range p.exprs {
		out.WriteString((e).String())
		out.WriteString("/n")
	}
	out.WriteString(";")
	return out.String()

}
func (p *Print) BuildSymbolTable(tables *st.SymbolTables, printId string) {}

func (r *Print) GetType() ty.Type {
	rty := ty.StringToType("unknown")
	return rty}

func (p *Print) TypeCheck(errors []string, tables *st.SymbolTables,funcId string) []string {
	for _, e := range p.exprs {
		errors = (e).TypeCheck(errors,tables, funcId)
		}
	return errors 
}
