package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
)

type UnaryOp struct {
	*token.Token 
	operator     Operator // The operator for the binary expression
	right        Expression 	// The right operand expression
	ty           types.Type
}


func NewUnaryOp(op Operator, right Expression, token *token.Token) *UnaryOp {
	return &UnaryOp{token, op, right, nil}
}

func (unOp *UnaryOp) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(" " + OpString(unOp.operator) + " ")
	out.WriteString((unOp.right).String())
	out.WriteString(")")

	return out.String()
}

func (u *UnaryOp) BuildSymbolTable(tables *st.SymbolTables) {}

//should only return type already exists simple call to the struct
func (unOp *UnaryOp) GetType() types.Type {
	return unOp.ty
}

func (unOp *UnaryOp) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	//call type check for left and right side of binOp
	errors = (unOp.right).TypeCheck(errors,tables, funcId)
	//afterwards call get type for left and right side of binOp

	intSingle := types.StringToType("int")
	boolSingle := types.StringToType("bool")
	unkSingle := types.StringToType("unknown")

	if (unOp.right).GetType() == intSingle{
		unOp.ty = intSingle
	}else if (unOp.right).GetType() == boolSingle{
		unOp.ty = boolSingle
	}else {
		unOp.ty = unkSingle
	}
		
	return errors
}