package ir

import (
	"strconv"
)
//generate unnique lables
//generate unique registers

var BlockCounter int = 0
var StartCounter int = 1
var LoopCounter int = 1
var IfCounter int = 1


type RegisterName struct {
	
}

	//need to generate new register names
	//decide if it is on the stack
	//keep pointer and iterate
func NewRegisterName(){
	//always start with % 
	//followed by either number or letter

}

type VarLabel struct {
	name string
}

func NewVarLabel(varName string)*VarLabel{
	newName := varName+strconv.Itoa(BlockCounter); BlockCounter +=1
	return &VarLabel{newName}
}