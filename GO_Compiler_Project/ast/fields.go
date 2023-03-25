package ast

import (
	"bytes"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)

type Fields struct {
	*token.Token
	decls	[]*Decl //might be single Decl or list of Decl
}

func NewFields(decls []*Decl, token *token.Token) *Fields {
	return &Fields{token,decls}
}
func (f *Fields) String() string {
	var out bytes.Buffer
	for _, decl := range f.decls{
		out.WriteString((*decl).String())
	}
	return out.String()
}

func (f *Fields) BuildSymbolTable(tables *st.SymbolTables, searchId string) {
	for _, decl := range f.decls {
		decl.BuildSymbolTable(tables,searchId)
	}
}

func (f *Fields) GetType() []ty.Type {
	var lType []ty.Type
	for _,dec := range f.decls {
		t := dec.GetType()
		lType = append(lType, t)
	}
	return lType
}

//Type check for fields will never be called because type check will be done at typedecl
//only need to confirm that fields does not have multiple variables with the same name
//in the same struct declaration, confirmed this in typedecl typecheck
func (f *Fields) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	return errors 
}

//Have this function to return the list of types for the list fo decls for typedecl to input
//into the block for that specific struct
func (f *Fields) BuildBlocks()string{
	var strTyFields string
	strTyFields = (strTyFields+"{")
	for i,d := range f.decls {
		frontEndType := d.ty.String()
		if frontEndType == "int"{
			backEndType := "i64"
			strTyFields = strTyFields+backEndType
		}else if frontEndType == "bool" {
			backEndType := "i64"
			strTyFields= strTyFields+backEndType
		}else if frontEndType == "unknown" {
			backEndType := "nil"
			strTyFields = strTyFields+backEndType
		}else if frontEndType == "void" {
			backEndType := "null"
			strTyFields = strTyFields+backEndType
		}else if frontEndType == "struct" {
			backEndType := "%struct."+d.id+"*"
			strTyFields = strTyFields+backEndType
		}
		if i != len(f.decls) -1 {//confirm that you are not on last iteration
			strTyFields = strTyFields+","
		}else{//if you are then put a bracket on it
			strTyFields = (strTyFields+"}")
		}
	}
	return strTyFields
}