package ast

import (
	"golite/token"
	"bytes"
	st "golite/symboltable"
// 	ty "golite/types"
)


type ForLiteral struct {
	*token.Token       // The token for the for literal
	expr		Expression
	blck        *Block 
}

func NewForLit(e Expression, b *Block, token *token.Token) *ForLiteral {
	return &ForLiteral{token, e, b}
}
func (f *ForLiteral) String() string {
	var out bytes.Buffer

	out.WriteString("For Literal")
	out.WriteString((f.expr).String())
	out.WriteString("\n")
	out.WriteString((*f.blck).String())
	return out.String()
}
func (f *ForLiteral) BuildSymbolTable(tables *st.SymbolTables, fId string){}

func (f *ForLiteral) TypeCheck(errors []string, tables *st.SymbolTables, funcId string)[]string{
	errors = (f.expr).TypeCheck(errors,tables,funcId)
	return errors
}