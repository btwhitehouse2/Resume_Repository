 package ir

// import (
// 	"fmt"
// 	"golite/ast"
// 	st "golite/symboltable"
// 	"os"
// 	"strings"
// )

// //could place inside of AST Program
// func TranslateToLLVM(tables *st.SymbolTables, tree *ast.Program,sourcePath string) []*BasicBlock{
// 	//need to write a file to .ll
// 	path := strings.Split(sourcePath,"/")
// 	fileName := path[len(path)-1]
// 	fileName = strings.Replace(fileName,".golite","",1)
// 	llFileName := "/llFiles/"+fileName+".ll"


// 	llFile, err:= os.Create(llFileName)
// 	if err != nil{
// 		fmt.Println("Creating .ll file had error",err)
// 	}
// 	fmt.Fprintf(llFile,"source_filename = \"%s\" \n",llFileName)
// 	fmt.Fprintf(llFile,"target triple = \"x86_64-linux-gnu\" \n")
	

// 	// return a slice of basic block full cfg for various functions
// 	var bblock []*BasicBlock
	
// 	bblock = Program.BuildBlocks(tables,bblock)

// 	//call each block and write each line to the file
// 	//complete with blocks
// 	//

// 	fmt.Fprintf(llFile,"declare i8* @malloc(i32)\n")
// 	fmt.Fprintf(llFile,"declare void @free(i8*)\n")
// 	fmt.Fprintf(llFile,"declare i32 @prinft(i8*, ...)\n")
// 	fmt.Fprintf(llFile,"declare i32 @scanf(i8*, ...)\n")
// 	return bblock	
// }

// // func NextNode(tree *ast.Program, currentNode ,seenNodes []) {

// // }