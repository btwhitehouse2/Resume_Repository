package ast

import (
	"bytes"
	//"golite/ir"
	st "golite/symboltable"
	"golite/token"
	"golite/types"

)


type CondLit struct {
	*token.Token      
	exp 		Expression
	ifTrue		*Block
	elseFalse	*Block
	ty        	types.Type 
}

func NewCondLit(exp Expression, i *Block, e *Block, token *token.Token) *CondLit {
	return &CondLit{token, exp, i, e, nil}
}
func (c *CondLit) String() string {
	var out bytes.Buffer

	out.WriteString((c.exp).String())
	out.WriteString("\n")
	out.WriteString(c.ifTrue.String())
	out.WriteString("\n")
	out.WriteString(c.elseFalse.String())
	return out.String()
}

func (c *CondLit) GetType() types.Type {
	return c.ty
}

//blocks or expressions may have been declared or a parameters for a variable func or struct
func (c *CondLit) BuildSymbolTable(tables *st.SymbolTables, fId string) {}


func (c *CondLit) TypeCheck(errors []string, tables *st.SymbolTables, funcId string)[]string{
	errors = (c.exp).TypeCheck(errors,tables,funcId)
	
	//look up the type of the expression and set the condlit to that type
	if (c.exp).GetType() != nil && c.ty == nil {
		c.ty = (c.exp).GetType()
	}else { // c.exp is nil then we have an error
		c.ty = types.StringToType("unknown")
	}
	return errors
}

// //single block
// func (c *CondLit) BuildBlocks(tables *st.SymbolTables,b *ir.BasicBlock){
// 	//append to last block
// 	//all local declarations will occur at the beginning of a function

// 	if len c.nested != 0 {//function is nested
		

// 	}
// 	expBlock := NewBasicBlock(nil,b,)
// }