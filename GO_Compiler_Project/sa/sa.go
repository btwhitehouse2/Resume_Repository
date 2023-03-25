package sa

import (
	"fmt"
	"golite/ast"
	st "golite/symboltable"
)

func reportErrors(errors []string) bool {
	//Exercise: Think about how to implement this method.
	if len(errors) > 0{
		//print out all of the errors for the user
		for _,e := range errors {
			fmt.Println(e)
		}
		return false
	}else{
		return true
	}
}

func Execute(program *ast.Program) *st.SymbolTables {

	//Define the compiler symbol tables
	
	tables := st.NewSymbolTables()
	
	errors := make([]string, 0)
	//return errors from incorrect table entries ie. two funcs same name
	//First Build the Symbol Table(s) for all global declarations
	errors = program.BuildSymbolTable(tables,errors)
	//Second perform type checking.
	errors = program.TypeCheck(errors, tables)
	if mainFunction := tables.Funcs.Contains("main"); mainFunction == nil{
		errorMessage := ("When searching for a main did not find one declared.")
		errors = append(errors, errorMessage)
	}
	//if !reportErrors(errors) {
	if reportErrors(errors) {
		//no errors return table
		return tables
	}
	//otherwise return nil
	return nil
}
