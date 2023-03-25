package ast

import (
	"bytes"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)

type InvocacationExpression struct {
	*token.Token
	variable 	string
	right		*Arguments
	ty 			ty.Type
}

func NewInvocExpr(variable string, right *Arguments, token *token.Token) *InvocacationExpression {
	return &InvocacationExpression{token, variable, right, nil}
}

func (i *InvocacationExpression) String() string {

	var out bytes.Buffer
	out.WriteString(i.variable)
	out.WriteString(" = ")
	out.WriteString((*i.right).String())
	out.WriteString(";")
	return out.String()
}

//variables should have already been declared will look
func (i *InvocacationExpression) BuildSymbolTable(tables *st.SymbolTables, fId string) {}

func (i *InvocacationExpression) GetType() ty.Type {
	return i.ty}


func (i *InvocacationExpression) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {

	errors = i.right.TypeCheck(errors, tables, i.variable)
	
	// if invocEntry := tables.Globals.Contains(i.variable); invocEntry !=nil { 
	// 	if i.ty != invocEntry.Ty{
	// 		invocEntry.Ty = ty.StringToType("unknown")
	// 		errors = append(errors,i.variable)
	// 	}

	// }else 
	if invocEntry := tables.Funcs.Contains(i.variable); invocEntry !=nil {
		if i.ty != invocEntry.RetTy{
			invocEntry.RetTy = ty.StringToType("unknown")
			errors = append(errors,i.variable)
		}
	}
	// else if invocEntry := tables.Structs.Contains(i.variable); invocEntry !=nil {
	// 	if i.ty != invocEntry.SType{
	// 		invocEntry.SType = ty.StringToType("unknown")
	// 		errors = append(errors,i.variable)
	// 	}
	// }
	//need to check the list of expressions on the right to make sure all of the types there allign with variable
	// invocEntry := tables.Structs.Contains(i.variable)
	// for _, exp := range i.right {
	// 	if exp.GetType() != i.ty{
	// 		invocEntry.SType = ty.StringToType("unknown")
	// 		errors = append(errors,i.variable)
	// 	}
	// }
	return errors 
}