package ast

import (
	"fmt"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)


type ReturnLit struct {
	*token.Token       // The token for the return literal
	Values       	Expression // The value or nil
	ty				ty.Type
}

func NewReturnLit(v Expression, token *token.Token) *ReturnLit {
	return &ReturnLit{token, v, nil}
}
func (r *ReturnLit) String() string {
	return fmt.Sprintf("%d", (r.Values))
}

//don't care about values that it returns because variables will be saved
func (r *ReturnLit) BuildSymbolTable(tables *st.SymbolTables, fId string) {}

func (r *ReturnLit) GetType() ty.Type {
	return r.ty}


func (r *ReturnLit) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	errors = r.Values.TypeCheck(errors, tables, funcId)
	
	funcEnt := tables.Funcs.Contains(funcId)
	if funcEnt.RetTy.String() != (r.Values).GetType().String() {
		fmt.Println("error found in return type check")
		fmt.Println(funcEnt.RetTy.String(), (r.Values).GetType().String())
		r.ty = ty.StringToType("unknown")
		errors = append(errors,"return lit type did not match return type")
		return errors
	}else{
		r.ty = funcEnt.RetTy
		return errors
	} 	

	}
