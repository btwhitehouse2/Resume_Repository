package ast

import (
	"bytes"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)

type Parameters struct {
	*token.Token
	decls	[]*Decl //might be single Decl or list of Decl
}

func NewParameters(decls []*Decl, token *token.Token) *Parameters {
	return &Parameters{token,decls}
}
func (f *Parameters) String() string {
	var out bytes.Buffer
	for _, decl := range f.decls{
		out.WriteString((*decl).String())
	}
	return out.String()
}

func (p *Parameters) BuildSymbolTable(tables *st.SymbolTables, funcId string) {
	//declare here instead of decl to avoid local delcaration vs parameter declaration
	for _, decl := range p.decls {
		if funcEntry := tables.Funcs.Contains(funcId); funcEntry !=nil {
			funcEntry.Parameters = append(funcEntry.Parameters, &st.VarEntry{decl.id,decl.ty,st.LOCAL})
	}
}
}

func (p *Parameters) GetType() {
	// var lType []ty.Type
	// for _,dec := range p.decls {
	// 	t := dec.GetType()
	// 	lType = append(lType, t)
	// }
	// return lType
}

func (p *Parameters) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	//check the funcs for a paramater entry
	//do I need to check every func table and call each paramater in each func table?
	for _, decl := range p.decls {
		errors = decl.TypeCheck(errors, tables, funcId)
	}
	funcEnt := tables.Funcs.Contains(funcId)
	tableParams := funcEnt.Parameters
	if len(tableParams) == len(p.decls){
		for i,decl := range p.decls {
			//if a parameter does match the type and name given in the symbol table continue to iterate else return error
			if decl.ty == tableParams[i].Ty && decl.id == tableParams[i].Name  {
			}else { //rerturn error
				funcEnt.RetTy = ty.StringToType("unknown")
				errorMessage := ("Function: "+funcId+"had a paramater id: "+decl.id+"type: "+decl.ty.String()+
				"that did not match the type and or name for that function in the ST")
				errors = append(errors,errorMessage)}
		}
	}else { //wrong number of params return error
	funcEnt.RetTy = ty.StringToType("unknown")
	errorMessage := ("Function: "+funcId+"had the wrong number of parameters")
				errors = append(errors,errorMessage)}
	return errors 
}