package ast

import (
	"golite/token"
	"bytes"
	ty "golite/types"
	st "golite/symboltable"
)


type TypeLiteral struct {
	*token.Token       
	Value        string 
	ty			ty.Type
}

func NewTypeLit(value string, ty ty.Type, token *token.Token) *TypeLiteral {
	return &TypeLiteral{token, value, ty}
}
func (t *TypeLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(t.Value)
	out.WriteString("Type: ")
	out.WriteString(t.ty.String())
	return out.String()
}

func (t *TypeLiteral) TypeCheck(errors []string, tables *st.SymbolTables){}

func (t *TypeLiteral) BuildSymbolTable(tables *st.SymbolTables) {}