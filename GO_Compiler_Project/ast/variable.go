package ast

import (
	"fmt"
	st "golite/symboltable"
	"golite/token"
	ty "golite/types"
	//"golite/ir"
)

type Variable struct {
	*token.Token
	Identifier 	string // The variable name
	ty 			ty.Type
}

func NewVariable(identifier string, token *token.Token) *Variable {
	unk := ty.StringToType("unknown")
	return &Variable{token, identifier, unk}
}
func (v *Variable) String() string { return v.Identifier }


func (v *Variable) BuildSymbolTable(tables *st.SymbolTables) {}

func (v *Variable) GetType() ty.Type{
	return v.ty
}

//need to update from unknown
//global variables of the same name cannot be defined
//local variables of the same name of the same function cannot be defined
func (v *Variable) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	//search locally for the variable that we are trying to declare
	//fmt.Println("type chek in variable funcid ",funcId," var id ",v.Identifier," ty ",v.ty)
	
	//first confirm that you are not recieving funcs and structs and variables
	if fEntry := tables.Funcs.Contains(v.Identifier); fEntry != nil{
		v.ty = fEntry.RetTy
		return errors
	}else if sEntry :=tables.Structs.Contains(v.Identifier); sEntry != nil {
		v.ty = sEntry.SType
		return errors
		//now check using the func ID
	}else if fEntry := tables.Funcs.Contains(funcId); fEntry != nil {
		
		for _, entry := range fEntry.Parameters {
			fmt.Println("entry ",entry)
			if entry.Name == v.Identifier{
				fmt.Println("func param ",entry.Name," varialbe ", v.Identifier)
				v.ty = entry.Ty
				return errors
			}
		}//searched through all of the parameters did not find
		if varEntry := fEntry.Variables.Contains(v.Identifier); varEntry != nil{
			fmt.Println("Funcs",funcId," id ",v.Identifier," varEntry type ",varEntry.Ty," Name ",varEntry.Name)
			//now check to make sure that the var entry type is not in the structs table
			if strucEnt := tables.Structs.Contains(varEntry.Name); strucEnt != nil{
				fmt.Println("a specific struct is the return type")
				v.ty = varEntry.Ty
				return errors
			}else{
				v.ty = varEntry.Ty
				fmt.Println(">>>>>>>")
				fmt.Println("Funcs",funcId," id ",v.Identifier," varEntry type ",varEntry.Ty," Name ",varEntry.Name)
				return errors

			}

		}
		//else not a locally declared variable then we need to keep searcing
		//search global variables or structs next
	}else if varEntry := tables.Globals.Contains(v.Identifier); varEntry != nil {
		v.ty = varEntry.Ty
		return errors
	} else if structEntry := tables.Structs.Contains(funcId); structEntry != nil {
		//we know it is part of a struct
		fmt.Println("struct in variable ",structEntry)
		for _, entry := range structEntry.Fields {
			if entry.Name == v.Identifier{
				v.ty = entry.Ty
				return errors
			}
		}


	}else {//You don't find the entry so therefore you need to define it and add to errors
			fmt.Println("Got to errors here.....................................")
			v.ty = ty.StringToType("unknown")
			errorMessage := ("When checking "+v.Identifier+" in function: "+funcId+" did not find in the symbol table")
			errors = append(errors, errorMessage)
			v.BuildSymbolTable(tables)
			return errors 
		}
		return errors
}

//single block
// func (v *Variable) BuildBlocks(tables *st.SymbolTables, blocks []ir.BasicBlock,funcId string)[]string{
// 	//append to last block
// 	//all local declarations will occur at the beginning of a function
// 	if funcId == ""{//we know this is a global variable
// 		currentBlock := blocks[len(blocks)-1]
// 		currentInstr := currentBlock.instruc
// 		globalVar := "@global"+v.Identifier+" = common global i64 0 \n"
// 		currentInstr = currentInstr+globalVar
// 		currentBlock.instruc = currentInstr 
// 	}
	
// 	//will need initial value
// 	//int = 0
// 	//bool = false
// 	//pointer = null //not nil
// 	return blocks
// }