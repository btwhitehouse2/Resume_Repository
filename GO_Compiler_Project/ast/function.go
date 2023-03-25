package ast

import (
	"bytes"
	"fmt"
	st "golite/symboltable"
	"golite/token"
	"golite/types"
	ty "golite/types"
	//"golang.org/x/vuln/vulncheck"
	//"honnef.co/go/tools/go/ir"
)

type Function struct {
	*token.Token
	id 			string
	params		*Parameters
	retType		ty.Type //return Type
	decls		[]*Declaration
	stmts		[]Statement
}

func NewFunction(idLit string, params *Parameters, tLit ty.Type, decls []*Declaration, stmts []Statement, token *token.Token) *Function {
	return &Function{token, idLit, params, tLit, decls, stmts}
}
func (f *Function) String() string {

	var out bytes.Buffer
	
	out.WriteString(f.id)
	out.WriteString("\n")
	out.WriteString((*f.params).String())
	out.WriteString("\n")
	out.WriteString(f.retType.String())
	out.WriteString("\n")	
	for _, dcl := range f.decls {
		out.WriteString((*dcl).String())
		out.WriteString("\n")
	}
	for _, stmt := range f.stmts {
		out.WriteString((stmt).String())
		out.WriteString("\n")
	}
	return out.String()
}
//function will be defined in func table
//variables will be defined in var table

//need to pass the function name
func (f *Function) BuildSymbolTable(tables *st.SymbolTables, errors []string)[]string{
	//local entries if they exist will come from the declarations or statements portions of function
	//first will iterate through the decls
	var fRet types.Type
	if f.retType == nil {
		fRet = ty.StringToType("void")
	}else{
		fmt.Println("-------Bld ST")
		fmt.Println(f.retType)
		fRet = f.retType
	}
	fmt.Println("Return type prior to function",fRet)
	//need to confirm that two funcs of the name do not exist
	if funcEnt := tables.Funcs.Contains(f.id); funcEnt != nil {
		//if it does exist set the table entry to unknown Type
		errorMessage := ("The following function:"+f.id+"has been declared more than once")
		errors = append(errors,errorMessage)
	}
	//create local entry
	var listVar []*st.VarEntry
	localTable := st.NewSymbolTable[*st.VarEntry](nil)
	tables.Funcs.Insert(f.id,&st.FuncEntry{f.id,fRet,listVar,localTable})
	
	newFunc := tables.Funcs.Contains(f.id)
	fmt.Println("After appending returntype",newFunc.RetTy)

	if funcEnt := tables.Funcs.Contains(f.id); funcEnt != nil {
		fmt.Print("'\n'table return type is",funcEnt.RetTy,"\n")
	}
	
	for _, dec := range f.decls {
		dec.BuildSymbolTable(tables,f.id)
	}//now go through statements
	for _, stm := range f.stmts {
		(stm).BuildSymbolTable(tables,f.id)
	}
	f.params.BuildSymbolTable(tables,f.id)
	return errors
}

func (f *Function) TypeCheck(errors []string, tables *st.SymbolTables) []string{
	errors = f.params.TypeCheck(errors,tables,f.id)
	for _,declaration := range f.decls {
		errors = declaration.TypeCheck(errors,tables,f.id)
	}
	for _,stmt := range f.stmts {
		errors = (stmt).TypeCheck(errors,tables,f.id)
	}
	//confirm that the function is defined in the correct name space by checking functions first
	//then the other names spaces first
	if funcEnt := tables.Funcs.Contains(f.id) ; funcEnt != nil {
	}else{ // funcId was not in Funcs Table
		if structEnt := tables.Structs.Contains(f.id) ; structEnt != nil{
			fmt.Print("...func not in table")
			//funcEnt.RetTy = ty.StringToType("unknown")
			errorMessage := ("Function: "+f.id+"was defined in the structs table")
			errors = append(errors,errorMessage)
		}else if varEnt := tables.Globals.Contains(f.id) ; varEnt != nil{
			fmt.Print("...var not in table")
			//funcEnt.RetTy = ty.StringToType("unknown")
			errorMessage := ("Function: "+f.id+"was defined in the globals table")
			errors = append(errors,errorMessage)
		}
	}  
	//confirm the return types is the same in the table
	funcEnt := tables.Funcs.Contains(f.id)
	
	//
	if f.retType != funcEnt.RetTy{
		fmt.Println("function return type: ",f.retType," table return type: ",funcEnt.RetTy)
		funcEnt.RetTy = ty.StringToType("unknown")
		fmt.Println("Function: "+f.id+" has a different return type then it has stored in the table: ")
		errorMessage := ("Function: "+f.id+" has a different return type then it has stored in the table")
		errors = append(errors,errorMessage)
	}

	//made it to the end then return the existing errors
	return errors
	}




	//define <ret_ty> @FUNC_NAME(<ty1>%PARM_NAME1, //Parameters will be registers assign a register for each
								//<ty2>%PARM_NAME2,
								//<ty3>%PARM_NAME3,
								//...
								//<tyN>%PARM_NAMEN)
	//{
	// 			;BASIC BLOCKS	
	//}

	//func Init(initVal int) *Point2D{} => define %struct.Point2D* @Init(i64 %initVal)
	//func main() => define i64 @main()
// func (f *Function) BuildBlocks(blocks []ir.Blocks, tables *st.SymbolTables)[]ir.BasicBlock{
// 	var result []ir.BuildBlock
// 	visited := map[string]bool{}

// 	for _,dec := range f.decls {


// 		}
	
// 	for i,stm := range f.stmts {
// 			if stm == "binop"& visited[stm.Id]==false{

// 			}else if stm == "attributecall" && confirm no visit{ //id.id.id
				
// 			}else if stm == "boollit" && confirm no visit{
				
// 			// }
// 			// else if stm == "callexpr"{//not
				
// 			}else if stm == "condlit" && confirm no visit{//if else statement
// 				statementLoop = statementLoop +"Iff"
// 				statementLoop = statementLoop +"Ife"
// 				statementLoop = statementLoop +"Els"
// 			}else if stm == "dellit" && confirm no visit{//delete
			
// 			}else if stm == "forlit"{//for loop
// 				statementLoop = statementLoop +"For"
// 			}else if stm == "intlit"{//integer
				
// 			}else if stm == "invocexpr"{//invocation
				
// 			}else if stm == "newlit"{//new instance of a struct
				
// 			}else if stm == "print"{
				
// 			}else if stm == "read"{//scan a doc
				
// 			}else if stm == "return"{// return statement
				
// 			// }else if stm == "variable declaration"{
				
// 			}else next Node
// 		return statementLoops
// 	}
// 	//now return [(for,if,for),<no loops>,for,if]
// 	for _,b in range := listFuncBlocks{
// 		if b == len(3){

// 		}
// 	}


// 		if visited[dec.ids.String()] == false{ //go for first AST in stack and see if we visited
		
// 			visited[len(AST)-i] = true
// 			//call build blocks
// 		}

// 	//decls will just be added to the instructions of the enter basic bloecks
// 	// b.NewStartBlock()
// 	//statements can have expressions

// 	//every function will return something default to zero
// 	return blocks
// }