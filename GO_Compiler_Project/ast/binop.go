package ast

import (
	"bytes"
	//"errors"
	"fmt"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)

type BinOpExpr struct {
	*token.Token            // The token information
	operator     Operator   // The operator for the binary expression
	right        Expression // The right operand expression
	left         Expression // The left operand expression
	ty           ty.Type
}

func isIntOp(op Operator) bool {
	return op == PLUS ||
		op == SUB ||
		op == ASTERISK ||
		op == FSLASH
}
func isComparOp(op Operator) bool {
	return op == GT ||
		op == GE ||
		op == LT ||
		op == LE ||
		op == NOT ||
		op == OR
}

func isBoolOp(op Operator) bool {
	return op == NOT ||
		op == OR ||
		op == ANDCOMP ||
		op == ISEQUAL ||
		op == NOTEQUAL ||
		op == OR
}


func NewBinOp(op Operator, right Expression, left Expression, token *token.Token) *BinOpExpr {
	return &BinOpExpr{token, op, right, left, nil}
}

func (binOp *BinOpExpr) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString((binOp.left).String())
	out.WriteString(" " + OpString(binOp.operator) + " ")
	out.WriteString((binOp.right).String())
	out.WriteString(")")

	return out.String()
}

//attribute call should already exist in the symbol table 
//should not need to add anything
func (b *BinOpExpr) BuildSymbolTable(tables *st.SymbolTables){}

//should only return type already exists simple call to the struct
func (binOp *BinOpExpr) GetType() ty.Type {
	return binOp.ty
}

func (binOp *BinOpExpr) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	//call type check for left and right side of binOp
	errors = (binOp.left).TypeCheck(errors, tables, funcId)
	errors = (binOp.right).TypeCheck(errors,tables, funcId )
	//afterwards call get type for left and right side of binOp

	intSingle := ty.StringToType("int")
	boolSingle := ty.StringToType("bool")
	unkSingle := ty.StringToType("unknown")
	voidSingle := ty.StringToType("void")

	if (binOp.left).GetType() == intSingle  {
		//check the type of left and right side
		if (binOp.right).GetType() == intSingle{
			binOp.ty = intSingle
		} else {binOp.ty = unkSingle} 

	} else if (binOp.left).GetType()  == boolSingle {
		if (binOp.right).GetType() == boolSingle{
			binOp.ty = boolSingle
		}else {
			binOp.ty = unkSingle
		}
	} else {
		fmt.Println("BINOP: no type selected return void")
		binOp.ty = voidSingle
	}

	return errors
}

