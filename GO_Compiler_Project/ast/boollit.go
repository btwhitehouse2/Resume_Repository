package ast

import (
	"fmt"
	"golite/token"
	"golite/types"
	st "golite/symboltable"
)

type BoolLiteral struct {
	*token.Token      // The token for the integer literal
	Value        bool // The value for the integer literal
	ty           types.Type
}

func NewBoolLit(value bool, token *token.Token) *BoolLiteral {
	return &BoolLiteral{token, value, nil}
}
func (bl *BoolLiteral) String() string {
	return fmt.Sprintf("%d", bl.Value)
}

func (a *BoolLiteral) BuildSymbolTable(tables *st.SymbolTables) {}


func (bl *BoolLiteral) GetType() types.Type {
	return bl.ty
}

func (bl *BoolLiteral) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string { 
	boolSingle := types.StringToType("bool")
	ty := boolSingle
	bl.ty = ty
	return errors
}
