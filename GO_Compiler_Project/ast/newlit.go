package ast

import (
	"fmt"
	"golite/token"
	st "golite/symboltable"
	ty "golite/types"
)
// Type double struct {
//     X int;
//     Y int;
// }
// Func main(){
//     Var d double;
//     New d;
// }

type NewLiteral struct {
	*token.Token       // The token for the new literal
	id		string
	ty		ty.Type
}

func NewNewLit(new string, token *token.Token) *NewLiteral {
	return &NewLiteral{token, new,nil}
}
func (n *NewLiteral) String() string {
	return fmt.Sprintf("%d", n.id)
}



func (n *NewLiteral) BuildSymbolTable(tables *st.SymbolTables) {}

//need to build
//search local make sure id is declare
//search symbol table for type 
func (n *NewLiteral) TypeCheck(errors []string, tables *st.SymbolTables,funcId string)[]string{
	if funcEntry := tables.Funcs.Contains(funcId); funcEntry != nil {
		//check variables in funct to confirm that variable has been defined
		if varEntry := funcEntry.Variables.Contains(n.id); varEntry != nil{
			n.ty = varEntry.Ty
			return errors
		}else{//function exists variable was not declared
			errorMessage := ("Function: "+funcId+"has a new declaration for"+n.id+" but variable was not locally declared first")
			errors = append(errors, errorMessage)
			funcEntry.RetTy = ty.StringToType("unknown")
			return errors
		}
	}
	return errors
}