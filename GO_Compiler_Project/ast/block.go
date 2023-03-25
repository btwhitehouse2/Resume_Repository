package ast

import (
	"golite/token"
	st "golite/symboltable"
	"bytes"
)

// func foo() {
// 	{
// 		{
// 		  printf("hello world");
// 		}
// 	}
//   }

//just: (), (expr), (expr,expr,...)
//maybe do only 1 expressions force later term to select additional expression?
type Block struct {
	*token.Token
	Stmts    []Statement 
}

func NewBlockStm(s []Statement, token *token.Token) *Block {
	return &Block{token, s}
}
func (b *Block) String() string {

	var out bytes.Buffer
	for _,stm := range b.Stmts{
		out.WriteString((stm).String())
		out.WriteString("\n")
	}
	out.WriteString(";")

	return out.String()

}


//what account for parenthesis when are those removed

func (bl *Block) BuildSymbolTable(tables *st.SymbolTables, fId string) {
	for _, stm := range bl.Stmts{
		(stm).BuildSymbolTable(tables, fId)
	}

}

func (bl *Block) GetType() {}

func (bl *Block) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string { 
	for _,Stm := range bl.Stmts{
		errors = (Stm).TypeCheck(errors,tables,funcId)
	}
	return errors
}

