package ast

import (
	"fmt"
	"golite/token"
	"golite/types"
	st "golite/symboltable"
)

type IntLiteral struct {
	*token.Token       // The token for the integer literal
	Value        int64 // The value for the integer literal
	ty           types.Type
}

func NewIntLit(value int64, token *token.Token) *IntLiteral {
	intTy := types.StringToType("int")
	return &IntLiteral{token, value, intTy}
}
func (il *IntLiteral) String() string {
	return fmt.Sprintf("%d", il.Value)
}
func (il *IntLiteral) GetType() types.Type {
	return il.ty
}

func (il *IntLiteral) TypeCheck(errors []string, tables *st.SymbolTables,funcID string) []string {
	intSingle := types.StringToType("int")
	ty := intSingle
	il.ty = ty
	return errors
}