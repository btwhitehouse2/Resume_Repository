package types

import (
	"sync"
)

//types that I need
//int
//bool
//variable "string"
//void "a function that does not return a value"
//struct
//base
//unknown "help handle errors"(does not need to be compared)

// Type includes information about working on Types
type Type interface {
	String() string
}

//
type StructType interface {
	Type
}

type StructTy struct{
	Id			string
}

func (structTy *StructTy) String() string {
	// for _, i := range structTy.fields {
	// 	fields[i] = t

	// }
	return "struct"
}

//need to define comparison for struct type

//global variables
var structLock = &sync.Mutex{}
var singleStruct *StructTy
func getInstanceStruct() *StructTy{
	if singleStruct == nil {
		structLock.Lock()
		defer structLock.Unlock()
		//creating single instance
		if singleStruct == nil {
			singleStruct = &StructTy{}
		}
		//else single instance already exists
	}
	return singleStruct
}


type IntTy interface {
	Type	
}
type intTy struct{
	i int 
}

func (i *intTy) String() string {
	return "int"
}

//Int Singleton
//global variables
var intLock = &sync.Mutex{}
var singleInt *intTy
func getInstanceInt() *intTy{
	if singleInt == nil {
		intLock.Lock()
		defer intLock.Unlock()
		//creating single instance
		if singleInt == nil {
			singleInt = &intTy{}
		}
		//else single instance already exists
	}
	return singleInt
}

type BoolTy interface {
	Type
}


type boolTy struct{
	b 		bool //point sinn
	// eq		*boolTy//singletong
	}	

func (b *boolTy) String() string {
	return "bool"
}

//Bool Singleton
//global variables
var boolLock = &sync.Mutex{}
var singleBool *boolTy
func getInstanceBool() *boolTy{
	if singleBool == nil {
		boolLock.Lock()
		defer boolLock.Unlock()
		//creating single instance
		if singleBool == nil {
			singleBool = &boolTy{}
		}
		//else single instance already exists
	}
	return singleBool
}

type VarTy interface {
	Type
}


type varTy struct{
	v  string
}

func (v *varTy) String() string {
	return "variable"
}

//Variable Singleton
//global variables
var varLock = &sync.Mutex{}
var singleVar *varTy
func getInstanceVar() *varTy{
	if singleVar == nil {
		varLock.Lock()
		defer varLock.Unlock()
		//creating single instance
		if singleVar == nil {
			singleVar = &varTy{}
		}
		//else single instance already exists
	}
	return singleVar
}

type UnkTy interface {
	Type
}


type unkTy struct{
//	u string
}

func (uTy *unkTy) String() string{
	return "unknown"
}

//Variable Singleton
//global variables
var unkLock = &sync.Mutex{}
var singleUnk *unkTy
func getInstanceUnk() *unkTy{
	if singleUnk == nil {
		unkLock.Lock()
		defer unkLock.Unlock()
		//creating single instance
		if singleUnk == nil {
			singleUnk = &unkTy{}
		}
		//else single instance already exists
	}
	return singleUnk
}

type VoidTy interface {
	Type
}
type voidTy struct {
	v    string
}
func (vTy *voidTy) String() string{
	return "void"
}
//Variable Singleton
//global variables
var vLock = &sync.Mutex{}
var singleV *voidTy
func getInstanceVoid() *voidTy{
	if singleV == nil {
		vLock.Lock()
		defer vLock.Unlock()
		//creating single instance
		if singleV == nil {
			singleV = &voidTy{}
		}
		//else single instance already exists
	}
	return singleV
}

func StringToType(ty string) Type {
	if ty == "int" {
		singletonInt := getInstanceInt()
		return singletonInt
	} else if ty == "bool" {
		singletonBool := getInstanceBool()
		return singletonBool
	} else if ty == "variable" {
		singletonVariable := getInstanceVar()
		return singletonVariable
	}else if ty == "unknown" {
		singletonVariable := getInstanceUnk()
		return singletonVariable
	}else if ty == "void" {
		singletonVariable := getInstanceVoid()
		return singletonVariable
	}else if ty == "struct" {
		singletonVariable := getInstanceStruct()
		return singletonVariable
	}
	panic("Not found type")
}
