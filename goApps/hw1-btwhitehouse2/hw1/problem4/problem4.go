package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"hw1/set"
)

func main() {
	// check to see if both inputs are empty
	//we keep this value at 1 because os.Args does not recognize < or > effectively 
	// we will always half the path for the program 
	if len(os.Args) == 1 {
		fmt.Println("Usage: problem4  <union | intersect | diff> file")
		os.Exit(0)
		//this checks to confirm that there are at least two inputs 
		}else if len(os.Args) == 2 {
			fmt.Println("Usage: problem4  <union | intersect | diff> file")
			os.Exit(0)
		}	

	//Now that we have confirmed the inputs are not missing, 
	//now we need to confirm they are valid
	methodType := os.Args[1]
	pathToFile := os.Args[2]
	methodBool := checkMethod(methodType)
	pathBool := checkPath(pathToFile)
	if methodBool && pathBool {
		//completes the requested method
		newIntSet := doMethod(methodType, pathToFile)
		newIntSet.PrintIntSeq()
	}else{
		fmt.Println("Usage: problem4  <union | intersect | diff> file")
	}
}
func checkMethod(mT string)bool{
	listMethods := []string{"add", "union", "intersect", "diff"}
	for _,val := range(listMethods){
		if val == mT {
		return true
	}
	}
	return false
}

func checkPath(pTF string)bool{
	_, error := os.Stat(pTF)
	if os.IsNotExist(error) {
		return false
	}else{
		return true
	}	
}

func doMethod(mType string,pTFile string) set.IntSet{
	intSet1 := set.NewIntSet()
	intSet2 := set.NewIntSet()
	intSet3 := set.NewIntSet()
	//newIntSetVar := *nIntSet
	i := 0
	f,_ := os.Open(pTFile)
	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	for scanner.Scan(){
		l := scanner.Text()
		stringSplit := strings.Split(l," ")
		for _,b := range stringSplit{
			firstInt,_ := (strconv.Atoi(b))
			if firstInt != 0 {
			if i == 0 {
				intSet1.Add(firstInt) 
			}else if i==1 {
				intSet2.Add(firstInt)
			}
		}
		}

		i+=1	
	}
	if mType == "union"{
		intSet3 := (intSet1.Union(intSet2))
		//intSeq3 := intSet3.ReturnIntSeq()
		return *intSet3
	}else if mType == "intersect" {
		intSet3 := (intSet1.Intersect(intSet2))
		//intSeq3 := intSet3.ReturnIntSeq()
		return *intSet3
	}else if mType == "diff"{
		intSet3 := intSet1.Diff(intSet2)
		//intSeq3 := intSet3.ReturnIntSeq()
		return *intSet3
	}
return *intSet3
}