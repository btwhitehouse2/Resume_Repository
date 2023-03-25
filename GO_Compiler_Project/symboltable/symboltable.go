package symboltable

import (
	"bytes"
	"fmt"
	"golite/types"
	"strconv"
	//"sync"
)

type VarScope int

const (
	GLOBAL VarScope = iota
	LOCAL
)

type StructEntry struct {
	Name		string
	SType	types.Type // will be a structType or unknown for error
	Fields		[]*VarEntry //field in a function
}

type VarEntry struct {
	Name     string
	Ty       types.Type
	Scope    VarScope
}

type FuncEntry struct {
	Name       string
	RetTy      types.Type //return type or unknown for error
	Parameters []*VarEntry //inputs
	Variables  *SymbolTable[*VarEntry] //variables defined inside of function
}

func (entry *FuncEntry)String() string { 
	
	var out bytes.Buffer
	out.WriteString(entry.Name)
	out.WriteString("\n")
	out.WriteString("type")
	out.WriteString(entry.RetTy.String())
	out.WriteString("\n")
	for _, p := range entry.Parameters{
		out.WriteString(p.String())
		out.WriteString("\n")
	}
	return out.String()
}

func (entry *StructEntry)String() string { 
	
	var out bytes.Buffer
	out.WriteString(entry.Name)
	out.WriteString("\n")
	for _, p := range entry.Fields{
		out.WriteString(p.String())
		out.WriteString("\n")
	}
	return out.String()
}

//unknown entry for variables that have not been defined
//add entries for predefined declarations


func (entry *VarEntry)String() string { 
	
	var out bytes.Buffer
	out.WriteString(entry.Name)
	out.WriteString("\n")
	out.WriteString("type")
	out.WriteString(entry.Ty.String())
	out.WriteString("\n")
	s := strconv.Itoa(int(entry.Scope))
	out.WriteString(s)
	return out.String()
}

type SymbolTables struct {
	Structs	*SymbolTable[*StructEntry] // can be reached via expression or declaration
	Funcs   *SymbolTable[*FuncEntry] //parent
	Globals *SymbolTable[*VarEntry] 
}

func NewSymbolTables() *SymbolTables {
	return &SymbolTables{NewSymbolTable[*StructEntry](nil),NewSymbolTable[*FuncEntry](nil), NewSymbolTable[*VarEntry](nil)}
}

//t represents the type parameter for that table: var or func
//stringer means that the typer parameter must implement the string method
type SymbolTable[T fmt.Stringer] struct {
	parent *SymbolTable[T]
	table  map[string]T
}

//parent needs to be a pointer to the parentt for func to ensure global vs local variable is correct
func NewSymbolTable[T fmt.Stringer](parent *SymbolTable[T]) *SymbolTable[T] {
	return &SymbolTable[T]{	parent,make(map[string]T)}
}

func (st *SymbolTable[T]) Insert(id string, entry T) {
	st.table[id] = entry
}

//need to look up if not in local table need to go to parent to look up
//need to see if it is globals
//structs
//funcs
func (st *SymbolTable[T]) Contains(id string) T {
	//search locals because local will override global
	var localVar T
	for key,value := range (st.table) {
		if key == id {
			return value}
		}
	if st.parent != nil {
		return st.parent.Contains(id)
	}
	return localVar
}


type Stringer interface {
	String() string
}
