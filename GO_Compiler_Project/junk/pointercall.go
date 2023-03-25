package ast

import (
	//"fmt"
	"golite/types"
	"golite/token"
	"bytes"
	st "golite/symboltable"
)


type PointerLit struct {
	*token.Token       // The token for the pointer literal
	attrName        string //variable //The current variable we are calling a memory address for
	ty 				types.Type
}

func NewPointerCall(attrN string, token *token.Token) *PointerLit {
	return &PointerLit{token, attrN, nil}
}

func (a *PointerLit) String() string {

	var out bytes.Buffer

	out.WriteString("attr name =")
	out.WriteString(a.attrName)
	out.WriteString("\n")
	out.WriteString(a.ty.String())
	out.WriteString("\n")
	return out.String()
}

func (a *PointerLit) BuildSymbolTable(tables *st.SymbolTables){}


//need to implement type checking
func (p *PointerLit) TypeCheck(errors []string, tables *st.SymbolTables){}