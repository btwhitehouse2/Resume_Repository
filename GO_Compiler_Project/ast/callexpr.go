package ast

import (
	"bytes"
	st "golite/symboltable"
	"golite/token"
	ty "golite/types"
)

type CallExpr struct {
	*token.Token            // The token information
	ID		string 
	args 	*Arguments
	ty    	ty.Type //could map return type but not needed
}

//need to implement type checking at some point
//calling symbol table to confirm var has been declared
//the declared var type matches in the input ie. var y = bool ; y = 2+1


//call this function from selector Term
func (c *CallExpr) String() string {
	var out bytes.Buffer
	out.WriteString(c.ID)
	out.WriteString("\n")
	out.WriteString(c.args.String())
	out.WriteString("\n")
	return out.String()
}


func NewCallExpr(argId string,  args *Arguments,t *token.Token) *CallExpr{
	unk := ty.StringToType("unknown")
	return &CallExpr{t, argId, args, unk}
}
	
//expressions may have been declared or a parameters for a variable func or struct
func (c *CallExpr) BuildSymbolTable(tables *st.SymbolTables) {}

func (c *CallExpr) GetType() ty.Type{
	return c.ty
}

func (c *CallExpr) TypeCheck(errors []string, tables *st.SymbolTables, funcId string) []string {
	errors = c.args.TypeCheck(errors,tables,funcId)

	if funcEntry := tables.Funcs.Contains(funcId); funcEntry != nil{ 
		for _,ent := range funcEntry.Parameters{ //check parameters
			if ent.Name == c.ID{
				c.ty = ent.Ty
				return errors
			}
		}
		if varEnt:= funcEntry.Variables.Contains(c.ID); varEnt != nil {
			c.ty = varEnt.Ty
			return errors
		}else{//have checked parameters and now have checked variables have not found make error
			errorMessage := ("When doing call expression for"+c.ID+" found function: "+funcId+"but ST did not have value in param or vars")
			errors = append(errors, errorMessage)
			c.ty = ty.StringToType("unknown")
			return errors
		}
	}else {//function not declared make error
		errorMessage := ("When doing call expression for"+c.ID+" could not find function: "+funcId)
		errors = append(errors, errorMessage)
		c.ty = ty.StringToType("unknown")
		return errors
	}
}



type CallExpression interface {
	String() string
	TypeCheck([]string, *st.SymbolTables) []string 
	GetType()ty.Type 
}