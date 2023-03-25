package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	ty "golite/types"
	"strconv"

	//"honnef.co/go/tools/go/ir"
)

type TypeDecl struct {
	*token.Token
	id			string
	flds		*Fields
	ty			ty.Type
}

func NewTypeDecl(id string, flds *Fields, token *token.Token) *TypeDecl{
	structT := ty.StringToType("struct")
	return &TypeDecl{token, id, flds, structT}
}
func (p *TypeDecl) String() string {

	var out bytes.Buffer
	out.WriteString(p.ty.String())
	out.WriteString("\n")
	out.WriteString(p.id)
	out.WriteString("\n")
	out.WriteString(p.flds.String())
	out.WriteString("\n")
	return out.String()
}
//want to record a struct literal as a struct type in structs table

//need to pass the struct name
func (t *TypeDecl) BuildSymbolTable(tables *st.SymbolTables,errors []string)[]string{
	var par []*st.VarEntry
	//confirmed that two structs with the same name do not exist when building symbol table and if they did then I added
	//struct with unknown type 
	if structEnt := tables.Structs.Contains(t.id); structEnt != nil {
		errorMessage := ("The following struct:"+t.id+"has been declared more than once")
		errors = append(errors,errorMessage)
	}else if structEnt := tables.Globals.Contains(t.id); structEnt != nil { //confirm that a global variable does not exist
		//the same name
		errorMessage := ("The following struct:"+t.id+"has been declared more than once")
		errors = append(errors,errorMessage)
		//struct entry does not exist and we need to declare
	}else{	tables.Structs.Insert(t.id, &st.StructEntry{t.id,ty.StringToType("struct"),par})
	}

	t.flds.BuildSymbolTable(tables, t.id)		
	return errors
}

func (t *TypeDecl) GetType() ty.Type {
	return t.ty
}

func (t *TypeDecl) TypeCheck(errors []string, tables *st.SymbolTables) []string {
	sEntry := tables.Structs.Contains(t.id)
	//check the fields to confirm that there are not duplicate field names
	//iterate through all of the fields and increment if you have seen a field multiple time
	entryMap := make(map[string]int)
	for _,fld := range sEntry.Fields{
		if entryMap[fld.String()] >= 1{
			entryMap[fld.String()] += 1
		}else{
			entryMap[fld.String()] = 1}		
	}
	for entry,use := range entryMap{
		if entryMap[entry] > 1 {
			timesUsed := strconv.Itoa(use)
			errorMessage := ("Struct"+t.id+"had duplicate fields with the following id:"+entry+"was used"+timesUsed+"times")
			errors = append(errors, errorMessage)
		}
	}
	return errors 
}

	//%TYPE_NAME = type {<ty1>,<ty2>,<ty3>,...,<tyN>}
	//%struct.Point2D = type{i64,i64}
	//%struct.Node = type{i64,%struct.node*}
// func (t *TypeDecl) BuildBlocks(tables *st.SymbolTables, blocks []ir.BasicBlock)[]ir.BasicBlock{
// 	currentBlock := blocks[len(blocks)-1]
// 	currentInstr := currentBlock.instruc
// 	strFields := t.flds.BuildBlocks() //list of strings representing the fields in ir format
// 	inputStruct := "%struct."+t.id+" = type"+strFields+"\n"
// 	structCall := "@.nil"+t.id+" = common global %struct."+t.id+"* null \n"
// 	currentInstr = currentInstr+inputStruct+structCall
// 	currentBlock.instruc = currentInstr
// 	return blocks
// }