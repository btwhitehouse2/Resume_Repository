package ast

import (
	//"fmt"
	"bytes"
	"fmt"
	st "golite/symboltable"
	"golite/token"
	ty "golite/types"
)


type AttributeLit struct {
	*token.Token       // The token for the attribute literal
	parent			[]string //array to store if the struct has multiple attributes prior to
	attrName        Expression //The current variable we are calling
	ty 				ty.Type
}

func NewAttributeCall(attrN Expression, parent []string, token *token.Token) *AttributeLit {
	return &AttributeLit{token, parent, attrN, nil}
}

func (a *AttributeLit) String() string {

	var out bytes.Buffer

	out.WriteString("parent attr =")

	for _, variable := range a.parent{
		out.WriteString(variable)
		out.WriteString("\n")
	}
	out.WriteString("\n")
	out.WriteString("attr =")
	out.WriteString((a.attrName).String())
	out.WriteString("\n")
	return out.String()
}

//attribute call should already exist in the symbol table 
//should not need to add anything
func (a *AttributeLit) BuildSymbolTable(tables *st.SymbolTables){
}

func (a *AttributeLit) GetType() ty.Type {
	return a.ty
}

//id.id.id.id
func (a *AttributeLit) TypeCheck(errors []string, tables *st.SymbolTables,funcId string) []string{
	errors = (a.attrName).TypeCheck(errors,tables,funcId)	
	
	//confirm that variable is declared check locally first by checking the function
	// Check if par.Ty is in global Struct Symbol Table, if it is in Symbol Table,
	// Fetch StructEntry, then check next id/parent[next] is in StructEntry.Fields
	// Either get to final id, and return correct, or return error
	if funcEntry := tables.Funcs.Contains(funcId); funcEntry != nil {
		//check parameters first
		for i, variable := range a.parent{//need to range over parent to ensure we are inspecting
			// if newStruc := tables.Structs.Contains(variable)
			//in the case of an id.id.id the first id might be a parameter or local decl after the first instance
			//must be reference to a struct
			
			// //a parameter matches the first part of the id
			// // If the first id is iin parameter
			if i== 0 {
				for _,par := range funcEntry.Parameters {
					if par.Name == variable {
						if i != len(a.parent) -1 { // confirm that we are not on the last value
							if structEnt := tables.Globals.Contains(par.Ty.String());structEnt != nil{
							//search the next value to make sure it is a type that exists
							}else {//the next value does not exist in structs return error
								errorMessage := ("When checking for an attribute lit "+variable+" for function: "+funcId+" did not have the second id in ST")
								errors = append(errors, errorMessage)
								fmt.Print("error: attributeCall 1\n")
								funcEntry.RetTy = ty.StringToType("unknown")
								return errors
							}
						}else {//looking at the last id, meaning first and last, but no errors so return errors
							return errors
						}
					}
					}
				if localDecl := funcEntry.Variables.Contains("variable");localDecl != nil{
					if i != len(a.parent) -1 { // confirm that we are not on the last value
						if structEnt := tables.Globals.Contains(localDecl.Ty.String());structEnt != nil{
						//search the next value to make sure it is a type that exists
						}else {//the next value does not exist in structs return error
							errorMessage := ("When checking for an attribute lit "+variable+" for function: "+funcId+" did not have the second id in ST")
							errors = append(errors, errorMessage)
							fmt.Print("error: attributeCall 2\n")
							funcEntry.RetTy = ty.StringToType("unknown")
							return errors
						}
					}else {//looking at the last id, meaning first and last, but no errors so return errors
						return errors
				}
			}
		}else if i == len(a.parent) -1 { // last id to search
			if structEnt := tables.Globals.Contains(variable);structEnt != nil{
				return errors
			}else {// on the last id found an error
				errorMessage := ("When checking for an attribute lit "+variable+" for function: "+funcId+" did not have the second id in ST")
				errors = append(errors, errorMessage)
				fmt.Print("error: attributeCall 3\n")
				funcEntry.RetTy = ty.StringToType("unknown")
				return errors
			}
		}else { //not the last id to search
			if structEnt := tables.Globals.Contains(variable);structEnt != nil{
			}else{
				errorMessage := ("When checking for an attribute lit "+variable+" for function: "+funcId+" did not have the second id in ST")
				errors = append(errors, errorMessage)
				fmt.Print("error: attributeCall 4\n")
				funcEntry.RetTy = ty.StringToType("unknown")
				return errors
			}
		}
		}
	}else{//function was not contained
		errorMessage := ("When checking for an attribute lit  for function: "+funcId+" did not have the second id in ST")
		errors = append(errors, errorMessage)
		funcEntry.RetTy = ty.StringToType("unknown")
		return errors
	}
	return errors
}
			
			
			