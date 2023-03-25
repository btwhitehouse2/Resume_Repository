package main

import (
	"fmt"
	"golite/lexer"
	"golite/parser"
	"golite/sa"
	"os"
	//"golite/ir"
)

func main() {
	inputSourcePath := os.Args[1]
	lexer := lexer.NewLexer(inputSourcePath)
	fmt.Println(inputSourcePath)
	parser := parser.NewParser(lexer)
	if ast := parser.Parse(); ast != nil {
		fmt.Println("-------")
		fmt.Println("ast",ast)
		fmt.Println("-------")
		if tables := sa.Execute(ast); tables != nil {
			fmt.Println(&tables)
			fmt.Println(tables)
			fmt.Println("passed")
			//tables are good now make intermediate language
			// llCode := ir.TranslateToLLVM(tables,ast,inputSourcePath)
			// fmt.Println(llCode)
		} else {
			//fmt.Println(&tables)
			//fmt.Println(*tables)
			fmt.Println("failed")
		}
		
	}else{
		print("AST error")
	}
	  

}
