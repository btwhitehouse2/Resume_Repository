package ir

// import (
// 	"golite/ast"
// 	st "golite/symboltable"
// 	"strconv"
// 	"bytes"
// )

// type Blocks interface {
// 	StartBlock
// 	LoopBlock
// 	IfBlock
// }

// //make a look up table?


// type BasicBlock struct {
// 	label		string  		// represents the label for current block
// 	blkgroup   	string			//what type of block group
// 	// nested		[]string //block names
// 	instruc		[]Instruction // represents the instructions for this block  
// 	pred        []*BasicBlock //would be block0
// 	succ        []*BasicBlock //would be block1, block2
// 	}

// func NewBasicBlock(label string, blkgroup string, ir []Instruction,  pred []*BasicBlock, suc []*BasicBlock, ) *BasicBlock{
// 	label := NewVarLabel()
// 	return &BasicBlock{label,blkgroup,ir,pred,suc}
// }

// type StartBlock struct { //used for begin + end
// 	label	string
// 	exp 	*BasicBlock
// 	term 	*BasicBlock
// }
// func NewStartBlock(lab string, exp *BasicBlock, term *BasicBlock) *StartBlock{
// 	return &StartBlock{lab,exp,term}
// }

// type LoopBlock struct {
// 	label	string
// 	// nest	[]*BasicBlock //block names included
// 	exp 	*BasicBlock //will always go to iter and term
// 	iter 	*BasicBlock //need to see if it is nested
// 	term	*BasicBlock 
// }
// type IfBlock struct {
// 	label	string
// 	// nest	[]*BasicBlock //block names included
// 	exp		*BasicBlock
// 	then	*BasicBlock
// 	elsedo	*BasicBlock
// 	term	*BasicBlock //need to see if it is nested if in for go to for iter
// }



// // func NewBlocks() *Blocks{
// // 	return &Blocks{blcks}
// // }

// func (b *BasicBlock) String() string {
// 	var out bytes.Buffer
// 	out.WriteString(b.label)
// 	out.WriteString(b.blkgroup)
// 	for _, instr := range b.instruc{
// 		out.WriteString(instr.String())
// 	}
// 	for _, pred := range b.pred{
// 		out.WriteString(pred.String())
// 	}
// 	for _, succ := range b.succ{
// 		out.WriteString(succ.String())
// 	}
// 	return out.String()
// }

// //always declared at the beginnig of an empty set of blocks
// // func (b *Blocks) NewStartBlock(){
	
// // 	sc := strconv.Itoa(startCounter)
// // 	newstart:= &StartBlock{"start"+sc,,}
	
// // 	b.list.append(newstart)
// // 	startCounter += 1
// // }



// //when we insert give it a list of values and the parent c
// func (b *[]BasicBlock) Insert(blks []*BasicBlock, blk *BasicBlock) []*BasicBlock{

// 	return blks
// }
