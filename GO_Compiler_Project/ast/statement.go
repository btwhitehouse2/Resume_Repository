package ast

import(
		"golite/types"
	 	st "golite/symboltable"
)

type Statement interface {
	String() string
	TypeCheck([]string, *st.SymbolTables, string) []string 
	GetType()types.Type 
	//BuildSymbolTable(tables *st.SymbolTables, funcId string) 
	BuildSymbolTable(*st.SymbolTables, string) 
	//BuildBlocks
}
