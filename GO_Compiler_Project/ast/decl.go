package ast

import (
	"bytes"
	"golite/token"
	"golite/types"
	st "golite/symboltable"
	ty "golite/types"
)
//use for fields in a struct
//'id' ('int' | 'bool' | '*''id')
type Decl struct {
	*token.Token
	id 		string // The lvalue for the declaration
	ty   	ty.Type
}

func NewDeclStm(id string, ty ty.Type, token *token.Token) *Decl {
	return &Decl{token, id, ty}
}
func (d *Decl) String() string {
	var out bytes.Buffer

	out.WriteString(d.ty.String() + " ")
	out.WriteString(d.id)
	out.WriteString(";")

	return out.String()
}

func (d *Decl) BuildSymbolTable(tables *st.SymbolTables, searchId string) {
	//being defined as part of a struct or a param for a function
	//look up struct first
	if structEntry := tables.Structs.Contains(searchId); structEntry !=nil {
		//if struct exists add a new decl field to thes struct
		structEntry.Fields = append(structEntry.Fields, &st.VarEntry{d.id,d.ty,st.GLOBAL})
		//next look to see if id is in funcs table as a parameter
		//variables not parameters
	}else if funcEntry := tables.Funcs.Contains(searchId); funcEntry !=nil {
		funcEntry.Variables.Insert(d.id, &st.VarEntry{d.id,d.ty,st.LOCAL})
	}//error
}

func (d *Decl) GetType() types.Type {
	return d.ty}


	//need to type check in the event that parameters calls a type check
	//however since the type is declared along with the variable only need to confirm that
	//there are not two local instances of the same variable
func (d *Decl) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	if structEntry := tables.Structs.Contains(funcId); structEntry !=nil {
		for _,fld := range structEntry.Fields {
			if fld.Name == d.id{ //found fields with matching names
				if fld.Ty == d.ty {
					return errors
				}else{//found the field with matching names and same struct name wrong types
					errorMessage := ("When assigning "+d.id+" in decl found the correct struct"+funcId+"and field mismatch of types")
					errors = append(errors, errorMessage)
					//structEntry.SType = ty.StringToType("unknown")
					return errors
				}
			}//assuming a func can have the same name as a struct no error goes here
		}
	}else if funcEntry := tables.Funcs.Contains(funcId); funcEntry !=nil {
		//check parameters first if not in parameter check variables
		for _,par := range funcEntry.Parameters{
			if par.Name == d.id{
				if par.Ty == d.ty{
					return errors
				}
			}
		}
		//check variables now
		if varEntry := funcEntry.Variables.Contains(d.id) ; varEntry !=nil { 
			if varEntry.Ty == d.ty{
				return errors
			}else {//found delcaration in ST, but did not match type
				errorMessage := ("When assigning "+d.id+" in decl found the correct func "+funcId+" and variable, but mismatch types")
				errors = append(errors, errorMessage)
				//structEntry.SType = ty.StringToType("unknown")
				return errors
			}
		}else{//searched func, func had matching name, but did not have the decl in sT
			errorMessage := ("When assigning "+d.id+" in decl found the correct func "+funcId+", but no correct variable")
			errors = append(errors, errorMessage)
			//structEntry.SType = ty.StringToType("unknown")
			return errors
		}
	}else{
		errorMessage := ("When assigning "+d.id+" in decl cound not find decl using funcId "+funcId+"")
		errors = append(errors, errorMessage)
		//funcEntry.RetTy = ty.StringToType("unknown")
		return errors
	}

	return errors
}
