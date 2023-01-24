package parser

import (
	"fmt"
	"os"
	ct "hw3/token"
)


/***
Here is the grammar for this specific parser
Program = Block 'eof'                       ; Pro = Blo 'eof'
Block = {Statement}                         ; Blo = Sta
Statement = Let | Assignment | Print        ; if Sta = Let// elseif Sta = Ass// elseif Sta = Pri 
Let = 'let' 'id' '=' Expression ';'         ; Let = 'let' 'id' '=' Exp ';'
Assignment = 'id' '=' Expression ';'        ; Ass = 'id' '=' Exp ';'
Print = 'print' Expression ';'              ; Pri = 'print' Exp ';'
Expression = Term ExpressionPrime           ; Exp = Ter Epr
ExpressionPrime = '+' Term ExpressionPrime  ; if Epr = '+' Ter Epr// else if Epr = '-' Ter Epr// else if Emp
ExpressionPrime = '-' Term ExpressionPrime  ;
ExpressionPrime = 'ε'                       ;
Term = Factor TermPrime                     ; Ter = Fac Tpr
TermPrime = '*' Factor TermPrime            ; if Tpr = '*' Fac Tpr// else if Tpr = '/' Fac Tpr// else if Tpr = Emp
TermPrime = '/' Factor TermPrime            ; 
TermPrime = 'ε'                             ;
Factor = 'INT' | 'id'                       ;Fac = 'int' // Fac = 'id'

***/

// Parser includes all fields necessary to perform recursive decent parsing
type Parser struct {
	tokens  []ct.Token
	currIdx int
}

// New creates and initializes a new parser
func New(tokens []ct.Token) *Parser {
	return &Parser{tokens, -1}
}

func (p *Parser) currToken() ct.Token {
	return p.tokens[p.currIdx]
}

func (p *Parser) nextToken() ct.Token {
	var token ct.Token
	//if the current value being seen is equal to the len do the following
	//first look will never be equal because -1 cannot equal empty or 0
	if p.currIdx == len(p.tokens)-1 {
		token = p.tokens[p.currIdx]
	} else {
	//every time it is not equal add 1 so not it is 0 or the first char
	//will return that token
		p.currIdx += 1
		token = p.tokens[p.currIdx]
	}
	return token
}

//calls the error on the token that did not match the grammar.
//returns the line and number of occurences
func (p *Parser) parseError() {
	err := fmt.Errorf("syntax error(line:%v): occurence %v of %q not allowed",
	p.currToken().Line,p.currToken().Count,p.currToken().Literal)
	fmt.Println(err)
	os.Exit(0)
}

//determines if the current token matches the expected token from the grammar
func (p *Parser) match(token ct.TokenType) bool {
	if token == p.currToken().Type {
		p.nextToken()
		return true
	}
	return false
}

//Will return a bool, which the driver will print false if no errors
//Or will return an error message
func (p *Parser) Parse() bool {
	p.nextToken()
	return Pro(p)
}


func Pro(p *Parser) bool { //Program in Cal Grammar
	return Blo(p) && Eof(p)
}

func Eof(p *Parser) bool { //End of Function token in Cal Grammar
	if p.match(ct.EOF) {
		return true}
	p.parseError()
	return false
}

func Blo(p *Parser) bool { //Block in Cal Grammar
	return Sta(p)
}

func Sta(p *Parser) bool { //Statement in Cal Grammar can have 0 or multiple statements
		if Let(p){
			return true
		}else if Pri(p){
			return true
		} else if Ass(p){
			return true
		}else {
			return true
		}
}

func Let(p *Parser) bool { //Let Statement in Cal attempts to match Let token 
	if p.match(ct.LET){ //If Let Token is matched then must be Let statement or error
		if p.match(ct.ID){
			if p.match(ct.EQUALS){
				if Exp(p){
					if p.match(ct.SCOLON){
						return true && Sta(p)
					}else{p.parseError()}	
				}else{p.parseError()}
			}else{p.parseError()}
		}else{p.parseError()}
	}
	return false
}

func Ass(p *Parser) bool { //Assignment Statement in Cal attempts to match Assignment Token
	if p.match(ct.ID){  //If Assignment Token is matched then must be Assignment statement or error
		if p.match(ct.EQUALS){
			if Exp(p) {
				if p.match(ct.SCOLON){
					return true && Sta(p)
				}else{p.parseError()}
			}else{p.parseError()}
		}else{p.parseError()}
	}
	return false
}

func Pri(p *Parser) bool { //Print Statement in Cal attempts to match Print Token
	if p.match(ct.PRINT) { //If Print Token is matched then must be Print statement or error
		if Exp(p) {
			if p.match(ct.SCOLON) {
				return true && Sta(p)
			}else{p.parseError()}
		}else{p.parseError()}
	}
	return false
}

func Exp(p *Parser) bool { //Expression in Cal 
	return Ter(p) && Epr(p)
}

func Epr(p *Parser) bool { //Expression Prime in Cal
	if p.match(ct.PLUS){
		return Ter(p) && Epr(p)
	}else if p.match(ct.MINUS){
		return Ter(p) && Epr(p)
	}else{
		return true
	}
}

func Ter(p *Parser) bool { //Term in Cal
	return Fac(p) && Tpr(p)
}

func Tpr(p *Parser) bool { //Term Prime in Cal
	if p.match(ct.ASTERISK){
		return Fac(p) && Tpr(p)
	}else if p.match(ct.DIVIDE){
		return Fac(p) && Tpr(p)
	}else{
		return true
	}
}

func Fac(p *Parser) bool { // Factor in Cal
	if p.match(ct.INT){
		return true
	}else if p.match(ct.ID){
		return true
	}
	return false
}
