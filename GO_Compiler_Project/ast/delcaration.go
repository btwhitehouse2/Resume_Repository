package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	//ty "golite/types"
	//"golite/ir"
)

//used in declarations and functions
//var id {, id} type
type Declaration struct {
	*token.Token
	ids		[]*Variable // The lvalue for the declaration
	Ty		types.Type
}

func NewDeclaration(ids []*Variable, ty types.Type, token *token.Token) *Declaration {
	return &Declaration{token, ids, ty}
}
func (d *Declaration) String() string {

	var out bytes.Buffer
	out.WriteString(d.Ty.String() + " ")
	for _,id := range d.ids {
		out.WriteString((*id).String())
	}
	out.WriteString(";")
	return out.String()
}

func (d *Declaration) BuildSymbolTable(tables *st.SymbolTables, funcId string) {
	//if we are defining the declaration here then it is being called by function and will be local
	//if it were global it would have been defined by program
	if funcEntry := tables.Funcs.Contains(funcId); funcEntry !=nil {
		//now iterate through each id and define a variable in the funcs table
		for _,id := range d.ids {
			funcEntry.Variables.Insert(id.String(),&st.VarEntry{id.String(),d.Ty,st.LOCAL})
			// variableTable := funcEntry.Variables
			// vEntry := &st.VarEntry{id.String(),d.Ty,st.LOCAL}
			// variableTable[vEntry] = nil
			//funcEntry.Variables[id.String()] = &st.VarEntry{id.String(),d.Ty,st.LOCAL}
		}
	}
}

func (d *Declaration) GetType() types.Type {
	return d.Ty}

//duplicate global already in errors	
func (d *Declaration) TypeCheck(errors []string, tables *st.SymbolTables, searchId string) []string {
	//need to type check in the event that parameters calls a type check
	for _,id := range d.ids{
		id.TypeCheck(errors,tables,searchId)
	}

	//need to look through the parameters and local declarations for a function cannot have same id for both
	//a param and a local variable
	if searchId != ""{ //we know that we are looking for a function
		funcEnt := tables.Funcs.Contains(searchId)
		for _,varEntry := range funcEnt.Parameters {
			if duplicateEntry := funcEnt.Variables.Contains(varEntry.String()); duplicateEntry!= nil{
				//Function contains a parameter and a variable declaration with the same name, error
				errorMessage := ("Function: "+searchId+"contains a parameter and a variable declaration with the same id:"+duplicateEntry.String())
				errors = append(errors, errorMessage)
			}
		}
	} //checked global variables for duplicate functions when I was building ST
	
	return errors
}

//single block
// func (d *Declaration) BuildBlocks(tables *st.SymbolTables, blocks []ir.BasicBlock,funcId string)[]string{
// 	for _,v := range d.ids {
// 		blocks = v.BuildBlocks(tables,blocks,funcId)
// 	}
	
// 	return blocks
// }