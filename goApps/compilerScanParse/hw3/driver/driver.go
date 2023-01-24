package main

import ("os"
		"bufio"
		"hw3/scanner"
		"hw3/token"
		"hw3/parser"
		"strconv"
		"strings"
		"fmt")



func main(){
	var eofType token.TokenType
	tokenarray := []token.Token{}
	
	stringSplit := []string{} //initialize array to compile read in lines
	lineNumArray := []string{} //initialize array that tracks line number for strings being read
	//ensure proper number of inputs or end program
	if len(os.Args) != 2 { 
		fmt.Println("Incorrect input:$go run scanner.go 'filename.cal'")
		os.Exit(0)
		}	
	//confirm file that we are attempting to read is real/ can be found
	pathToFile := os.Args[1]
	_, error := os.Stat(pathToFile)
	if os.IsNotExist(error) {
		fmt.Println("Could not reach path for file:$go run scanner.go 'filename.cal'")
		os.Exit(0)
	}
	//read in inputs and write to array for the program
	f,_ := os.Open(pathToFile)
	defer f.Close()	
	reader := bufio.NewScanner(f)
	curL := 1
	for reader.Scan(){
		lineText := reader.Text()
		lineSplit := strings.Split(lineText," ")
		for _,str := range lineSplit{
			curString := str
			stringSplit = append(stringSplit,curString)
			curS := strconv.Itoa(curL)
			lineNumArray = append(lineNumArray,curS)
			//fmt.Println(stringSplit,"___after string append")
		}
		curL += 1
		}	
	//initialize *Scanner Type
	returnScan := scanner.New(stringSplit,lineNumArray)

	//the Driver will continue to return next scan until the Scanner Lexem
	//has no more values to scan	
	for len(returnScan.Lexem) > 0 {
		nToken := returnScan.NextToken()
		tokenarray = append(tokenarray, nToken)
	}

	//At the end of the Token stream need to append EOF
	eofType = token.EOF
	lToken:= tokenarray[len(tokenarray)-1]
	lastLine := lToken.Line
	tokenarray = append(tokenarray,token.Token{eofType,1,"",lastLine})
	//Taking the Token stream and returning the Parse variation
	
	returnParse := parser.New(tokenarray)
	
	//Parses the input to confirm if it follows the grammar will return true if it does
	//and will print false or an error message will populate from rdp.go and program ends 
	//with no return to driver.

	curToken := returnParse.Parse()
	if curToken{
		fmt.Println(false)
	}
	
}
