package ast

import (
	"fmt"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)


type ReadLiteral struct {
	*token.Token       // The token for the read literal
	Expr       		Expression // The value for the read literal
}

func NewReadLit(exp Expression, token *token.Token) *ReadLiteral {
	return &ReadLiteral{token, exp}
}
func (r *ReadLiteral) String() string {
	return fmt.Sprintf("%d", (r.Expr))
}

func (r *ReadLiteral) GetType() ty.Type {
	rty := ty.StringToType("unknown")
	return rty}


//don't care about values that it returns because variables will be saved
func (r *ReadLiteral) BuildSymbolTable(tables *st.SymbolTables, funcId string) {}


func (r *ReadLiteral) TypeCheck(errors []string, tables *st.SymbolTables, funcId string)[]string{
	errors = (r.Expr).TypeCheck(errors,tables,funcId) 
	return errors
}