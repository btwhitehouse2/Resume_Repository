package scanner

import (
	"hw3/token"
	u "unicode"
)

//define scanner type
type Scanner struct {
	Lexem []string	//array of st that we will iterate through until empty
	Count map[string]int //map of strings that tracks # of occurences for an item
	LineNum []string
}

type InputStruct struct{
	InputString	string
	LineNum     int
}		
	
//Initializes a type *Scanner for creating a token stream
func New(str []string, intL []string) *Scanner {
	
	//this will in take the current list of strings and line numbers
	//will remove all strings and line numbers that have an empty string
	//meaning a blank line
	noEmptyStr := []string{}
	noEmptyLine := []string{}
	cur:=-1
	for _,strVal := range str {
		cur += 1
		if strVal != ""{
			noEmptyStr = append(noEmptyStr, strVal)
			noEmptyLine = append(noEmptyLine, intL[cur])
		}
	}
	nCount := map[string]int{}
	return &Scanner{noEmptyStr,nCount,noEmptyLine}
}

// func (i *InputStruct) Range(){
	
// 	return []{string,int}
// }

//next token will continue to iterate until it hits a return
//the premise of this program is that it will either hit a symbol
//or if it does not then it is an int or identity 
//digit track: ensures that if during a slice of an array contains a letter
//that this is not forgotten at the end when assigning type
//range input: keeps track of the number of characters seen in a specific string
//in the event that a string in fact contains multiple tokens ie. "2+2" needs seperating
//illegalInt: tracks for an int that begins with 0 and has additional digits ie. 001 is illegal
//listS: determines if a symbol from list of symbol is being used
//the function IsSymbol from unicode was not inclusive of all symbols
//Token_Picker exists as the general return point for NexToken
//The body of this in the for loop checks length via rangeInput and if INT via digitTrack  
func (l *Scanner) NextToken() token.Token {
	listOfSmybols := []string{"/","*",";","-"}
	//iterate throug the first value realizing that the program will end upon return
	//and will only iterate through at most 1 string in array of string l.Lexem at a time
	for _,input := range l.Lexem {
		rangeInput := 0
		digitTrack := true
		firsIntZero := false
		illegalInt := false
		listS := false
		var tType token.TokenType	
		for _,char := range input {
			//check for letters first since we have a Bool track for numbers
			if u.IsLetter(char){
				//breaks the string we are reading in to confirm that the identifier 
				//did not begin with a number followed by an letter we would process
				//this as INT,ID 
				if rangeInput > 0 && digitTrack && !illegalInt {
					tType = token.INT
					goto TOKEN_PICKER
				}else if rangeInput > 0 && digitTrack && illegalInt {
					tType = token.ILLEGAL
					goto TOKEN_PICKER
				}else{
				//Will reach this as long as this is the first char being seen
				//Or all previous Chars have been letters
				digitTrack = false
				rangeInput +=1
				}
			//concerned about recieving incorrect int type ie. 001 is illegal not an int
			//however 0 is an int. Need to look at first digit after second digit
			}else if u.IsDigit(char) && rangeInput== 0{
				rangeInput +=1
				if string(char)=="0"{
					firsIntZero = true
				}
				//second digit or later
			}else if u.IsDigit(char) && digitTrack && firsIntZero{
					illegalInt = true
					rangeInput +=1
					//cotinue to iterate if the first value in this int or string was not a 0			
			}else if u.IsDigit(char) && rangeInput!= 0 {
					rangeInput +=1
			}

			
			//we hit a symbol we know that if the current input is greater than we need to break
			//off the input prior to the symbol or just break off the symbol.
			SYMBOL:
			if u.IsSymbol(char)||listS == true{
				//not the first rune to be looked at
				if rangeInput > 0{
					//no letters have been seen so far
					if digitTrack && !illegalInt {
						tType = token.INT
						goto TOKEN_PICKER
					}else if digitTrack && illegalInt{
						tType = token.ILLEGAL
						goto TOKEN_PICKER
					}else{
						//need to decide String keyword Print or Let
						curItem := l.Lexem[0]
						strItem := curItem[0:rangeInput]
						tType = pickStringToken(strItem)
						digitTrack = false
						goto TOKEN_PICKER
					}
				}
			//first character you see is a symbol
			if (u.IsSymbol(char)||listS == true) && rangeInput==0{
				rangeInput +=1
				tType = pickSymbolToken(string(char))
				goto TOKEN_PICKER
				}
		}
		//certain  symbols are not appearing to match
		for _,b := range listOfSmybols{
			if string(char) == b {
				listS = true
				goto SYMBOL
			}
		}

		//illegal character found
		if !(u.IsLetter(char)|| u.IsNumber(char) || u.IsSymbol(char)){
			//this is the first character seen
			if rangeInput == 0{
				tType = token.ILLEGAL
				rangeInput += 1
				goto TOKEN_PICKER
			}else{
				goto UNASSIGNED_TYPE
			}
		}
	}

	//cycled through all of the runes and did not hit a symbol
	//must be int or string
	UNASSIGNED_TYPE:	
	if digitTrack && !illegalInt {
		tType = token.INT
	}else if digitTrack && illegalInt{
		tType = token.ILLEGAL
	}else{
		curItem := l.Lexem[0]
		curSlice := curItem[0:rangeInput]
		tType = pickStringToken(curSlice)
	}


	TOKEN_PICKER:
	curItem := l.Lexem[0]
	//when the current string fully matches one of our types ie. string, int, symbol
	//meaning proper spacing allowed for characters to be split correctly
	if len(curItem) == rangeInput {
		l.Count[curItem] +=1
		cCount := l.Count[curItem]
		curLine := l.LineNum[0]
		fToken := token.Token{tType,cCount,curItem,curLine}
		l.Lexem = l.Lexem[1:]
		l.LineNum = l.LineNum[1:]
		return fToken
	}else {
		//when splitting a string
		curSlice := curItem[0:rangeInput]
		l.Count[curSlice] += 1
		cCount := l.Count[curSlice]
		curLine := l.LineNum[0]
		fToken := token.Token{tType,cCount,curSlice,curLine}
		l.Lexem[0] = curItem[rangeInput:]
		return fToken
	}
	}
	//in the event that an item does not get caught we return this illegal error token
	return token.Token{token.ILLEGAL,0, "0","1"}
}

func pickStringToken(s string) token.TokenType {
	if s == "print" && len(s) == 5{
		return token.PRINT
	}else if s == "let" && len(s) == 3{
		return token.LET
	}
	//if it gets to the boottom return ID because it is a string
	return token.ID
}


func pickSymbolToken(s string) token.TokenType {
	if s == "="{
		return token.EQUALS
	}else if s == "+" {
		return token.PLUS
	}else if s == "-" {
		return token.MINUS
	}else if s == "/" {
		return token.DIVIDE
	}else if s == "*" {
		return token.ASTERISK
	}else if s == ";" {
		return token.SCOLON
	}
	//if it gets to the boottom return illegal
	return token.ILLEGAL
}