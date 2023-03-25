package ast

import (
	"bytes"
	"fmt"
	st "golite/symboltable"
	"golite/token"
	ty "golite/types"
)

type Assignment struct {
	*token.Token
	variable	string     // The lvalue for the assignment statement
	right    	Expression // The value assigned to the variable of this statements
	ty			ty.Type
}

func NewAssignStm(variable string, right Expression, token *token.Token) *Assignment {
	return &Assignment{token, variable, right,nil}
}
func (a *Assignment) String() string {

	var out bytes.Buffer

	out.WriteString(a.variable)
	out.WriteString(" = ")
	out.WriteString((a.right).String())
	out.WriteString(";")

	return out.String()
}

//item will have been declared don't need to update symbol table
func (a *Assignment) BuildSymbolTable(tables *st.SymbolTables, funcId string) {}

func (a *Assignment) GetType() ty.Type {
	return a.ty
}

func (a *Assignment) TypeCheck(errors []string, tables *st.SymbolTables,funcId string) []string {
	errors = a.right.TypeCheck(errors, tables, funcId)

	//confirm that variable is declared check locally first by checking the function
	if funcEntry := tables.Funcs.Contains(funcId); funcEntry != nil {
		//check parameters first
		for _,v := range funcEntry.Parameters {
			if v.Name == a.variable{
				if v.Ty.String() == a.right.GetType().String(){
					//only needs to have one parameter that matches this assignment
					a.ty = v.Ty
					return errors
				}else {// we found the variable was locally delcared in parameter, but has a different type report error
					errorMessage := ("When assigning "+a.variable+" in assignment found variable in parameters of funcentry did not match type!")
					errors = append(errors, errorMessage)
					//funcEntry.RetTy = ty.StringToType("unknown")
					return errors
				}
			}
		}
		//made it through the slice of parameters now we need to check the local declarations
		if localEntry := funcEntry.Variables.Contains(a.variable); localEntry != nil {
			//the expression is empty so no need to compare
			
			if a.right.String() == "nil" {
				return errors
				//next check to see if the two type match
			}else if localEntry.Ty.String() == a.right.GetType().String(){
				a.ty = localEntry.Ty
				return errors
			}else {
				errorMessage := ("When assigning "+a.variable+" in assignment found variable in local funcentry variables did not match type?+ assignment "+a.String())
				errors = append(errors, errorMessage)
				//funcEntry.RetTy = ty.StringToType("unknown")
				return errors
			}
		}
		//searched the function have not found in parameters or local declarations now time to check
		//globals
	} else if varEntry := tables.Globals.Contains(a.variable); varEntry != nil {
		//var entry exists in a table 
		//need to check that the left and right types match
		if varEntry.Ty == (a.right).GetType(){
			return errors
		}else{
			errorMessage := ("When assigning"+a.variable+"in assignment found variable did not match type")
			errors = append(errors, errorMessage)
			varEntry.Ty = ty.StringToType("unknown")
			return errors
		}
	// } //don't believe that assignment can define a new struct
	// else if structEntry := tables.Structs.Contains(funcId); structEntry != nil {
	// 	for _,v := range structEntry.Fields {
	// 		if v.Name == (*a.right).String() {
	// 			if v.Ty == (*a.right).GetType(){
	// 				a.ty = v.Ty
	// 				//only needs to have one parameter that matches this assignment
	// 				return errors
	// 			}else{ //found the declaration however the type did not match
	// 				errorMessage := ("When assigning"+a.variable+"in assignment found struct in structs table did not match type")
	// 				errors = append(errors, errorMessage)
	// 				varEntry.Ty = ty.StringToType("unknown")
	// 				return errors
	// 			}
	// 		}
	// 	}
	}else {
	//could not find the value in any table need to return as an error because it should have 
	//been declared
	errorMessage := ("When assigning"+a.variable+"in assignment could not find the variable being declared")
	errors = append(errors, errorMessage)
	varEntry.Ty = ty.StringToType("unknown")
}

return errors
}

//single block
func (p *Assignment) BuildBlocks(tables *st.SymbolTables){
	//append to last block
	//all local declarations will occur at the beginning of a function
	
}