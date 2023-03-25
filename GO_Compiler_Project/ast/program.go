package ast

import (
	"bytes"
	"fmt"
	//"golite/sa"
	//"golite/ir"
	st "golite/symboltable"
	"golite/token"
)

type Program struct {
	*token.Token
	typeDecls 	[]*TypeDecl
	decls		[]*Declaration //these must be global 
	funcs 		[]*Function //Decls in here will be local
}

func NewProgram(typeDecls []*TypeDecl, decls []*Declaration, funcs []*Function, token *token.Token) *Program {
	return &Program{token, typeDecls, decls, funcs}
}
func (p *Program) String() string {

	var out bytes.Buffer
	for _, tydcl := range p.typeDecls {
		out.WriteString((*tydcl).String())
		out.WriteString("\n")
	}
	for _, decl := range p.decls {
		out.WriteString((*decl).String())
		out.WriteString("\n")
	}
	for _, funct := range p.funcs {
		out.WriteString((*funct).String())
		out.WriteString("\n")
	}
	return out.String()
}
func (p *Program) BuildSymbolTable(tables *st.SymbolTables, errors []string)[]string{

	for _, tydcl := range p.typeDecls {
		errors = tydcl.BuildSymbolTable(tables,errors)
	}

	//all declarations that come before function will be global 
	//insert these declarations here to avoid confusion later
	for _, decl := range p.decls {
		for _,id := range decl.ids {
			//check to see if it has already been declared if it has add an error
			if  globalEntry := tables.Globals.Contains(id.String());globalEntry != nil{
				errorMessage := ("The following Global variable:"+id.String()+"has been declared more than once")
				errors = append(errors,errorMessage)
			}else if structEntry := tables.Structs.Contains(id.String());structEntry != nil{
				errorMessage := ("The following Global variable:"+id.String()+"has been declared more than once")
				errors = append(errors,errorMessage)
			}else{//global variable has not add to global variables
				tables.Globals.Insert(id.String(),&st.VarEntry{id.String(),decl.Ty,st.GLOBAL})
			}
		}
	}
	//call the funcs
	for _, funct := range p.funcs {
		funct.BuildSymbolTable(tables,errors)
	}
	return errors
}
func (p *Program) TypeCheck(errors []string, tables *st.SymbolTables)[]string{
	//while making symbol table errors were built for the following if they occurred
	//duplicate global variables
	//structs with duplicate names
	//duplicate functions

	//type check needs to confirm
	//if declared struct and global struct are the same (ie. length and variable names, order and type)
	for _, tyDec := range p.typeDecls {//check for duplicate field names
		errors = tyDec.TypeCheck(errors, tables)
	}
	for _, decl := range p.decls {
		errors = decl.TypeCheck(errors, tables, "")
	}
	for _, fun := range p.funcs {
		errors = fun.TypeCheck(errors, tables)
	}
	fmt.Println("From Program TypeCheck",errors)
	return errors 
}


//target info at the start
//
//call bb for decl typedecl and decl and funcs

//if last block do the following 
// func (p *Program) BuildBlocks(tables *st.SymbolTables, blocks []ir.BasicBlock)[]ir.BasicBlock{
	
// 	//need to inser the following into blocks
// 	//
// 	if len(blocks) == 0 { //need to insert a block
// 		blocks = blocks.Insertblock
// 	}
// 	for  _, tyDec := range p.typeDecls {
// 		structId := tyDec.id 
// 		blocks = tyDec.BuildBlocks(tables, blocks,structId)
// 	}
// 	for _, decl := range p.decls {
// 		funcId := "" //know it is global
// 		blocks = decl.BuildBlocks(tables, blocks,funcId)
// 	}
// 	for _, f := range p.funcs{
// 		blocks = f.BuildBlocks(tables,blocks)
// 	}


// 	//need to add lines to the end for "Function Declarations" to run C-runtime
// 	currentBlock := blocks[len(blocks)-1]
// 	currentInstr := currentBlock.instruc
// 	cFunctionDecl := "declare i8* @malloc(i32)\ndeclare void @free(i8*)\ndeclare i32 @printf(i8*, ...)\ndeclare i32 @scanf(i8*, ...)\n"
// 	currentInstr = currentInstr+cFunctionDecl
// 	currentBlock.instruc = currentInstr 
	

// 	return blocks
// }
