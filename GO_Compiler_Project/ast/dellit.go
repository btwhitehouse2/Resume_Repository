package ast

import (
	"bytes"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)


type DelLiteral struct {
	*token.Token       // The token for the integer literal
	exp 		Expression // The value for the integer literal
}

func NewDelLit(exp Expression, token *token.Token) *DelLiteral {
	return &DelLiteral{token, exp}
}
func (del *DelLiteral) String() string {
	var out bytes.Buffer

	out.WriteString("Delete Literal:")
	out.WriteString("\n")
	out.WriteString((del.exp).String())
	return out.String()
}

func (r *DelLiteral) GetType() ty.Type {
	rty := ty.StringToType("unknown")
	return rty}

func (a *DelLiteral) BuildSymbolTable(tables *st.SymbolTables, fId string){
}

//don't need to type check for del lit
func (p *DelLiteral) TypeCheck(errors []string, tables *st.SymbolTables, funcId string)[]string{
	errors = (p.exp).TypeCheck(errors,tables,funcId)
	return errors
}
