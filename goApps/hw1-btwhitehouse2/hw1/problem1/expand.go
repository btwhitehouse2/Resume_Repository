//websites and outside sources referenced and utilized
//Previously submitted work
//Assistance from https://www.geeksforgeeks.org/golang-program-that-removes-duplicate-elements-from-the-array/
//https://stackoverflow.com/questions/39013158/how-do-i-convert-string-of-int-slice-to-int-slice
//https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
// Golang Program that Removes Duplicate
// Elements From the Array

package problem1

import (
	"strconv"
	"strings"
)

//expand is the main function meant to be run from the command line
//calls other functions as filters in order to take a string and return
//a splice of ints
func Expand(intList string) []int {
	var slice = []int{}
	var list []int
	var newInt int
	var oldInt string
	unique_items := StrExpStr(intList)
    
	intL2 := strings.Trim(unique_items," ")
   splitList := strings.Split(intL2, " ")
   for i := range(splitList) {

    oldInt = splitList[i]
    if len(oldInt)> 0{
        newInt, _ = strconv.Atoi(oldInt)
       list = append(list, newInt)
       noDupList := removeDuplicates(list)
       list = noDupList
    }else {
        return slice
    }
   }
   slice = list
   return slice
}
	
//taken from source listed above
//removes duplicates values from a list of ints
func removeDuplicates(intS []int) []int {
    Keys := make(map[int]bool)
    list := []int{}
    for _, i := range intS {
        if _, v := Keys[i]; !v {
            Keys[i] = true
            list = append(list, i)
        }
    }
    return list
}

//this serves as the other main portion of the code
//it takes a string splits it by commas
//deletes whitespace, then it calls other functions
//to expand the range of values
//uses the formatMissing function to determine what if any 
//strings need to be voided
//finally it uses check for words to eliminate any non int values
func StrExpStr(input string) string{
    var  finlString string
    
    splitInput := strings.Split(input,",")
    for _, s := range splitInput{
        split2 := MultRange(s)
        s2 := strings.Trim(split2," ")
        s3 := formatMissing(s2)
        words := checkForWords(s3)
        if words == false {
            s3 = ""
        }
        if len(s3) != 0 {
            if len(finlString) > 0 {
                finlString = finlString +" "+ s3
            }else{
                finlString = s3
            }
        }
    }
    return finlString
}

//this inspects for any of the possible formatting errors ie. -5 will
//return nothing instead of being expanded
func formatMissing(s2 string) string{
    var intEnd, intStart, curInt int
    var expandVar []int
    var  newString, complString, s4 string

    subStr := strings.Index(s2,"-")   
    if subStr > -1 {
            s3 := strings.Trim(s2," ")
            splitSubStr := strings.Split(s3,"-")
            startV := splitSubStr[0]
            endV := splitSubStr[1]
            if startV == ""{
                s4 = ""
                return s4
            }
            if endV == ""{
                s4 = ""
                return s4  
            }
            intStart, _ = strconv.Atoi(startV)
            intEnd, _ = strconv.Atoi(endV)
            checkIntFormat := LargeIntLast(intStart,intEnd)
            if (checkIntFormat == false){
                return ""
            }
            for j := intStart; j <= intEnd; j++{
                expandVar = append(expandVar,j)
            }
            for k := range expandVar{ 
                curInt = expandVar[k]
                newString = strconv.Itoa(curInt)
                if len(complString) > 0 {
                complString = complString +" "+ newString
                }else{
                    complString = newString
                }
            }
            s4 = complString    
    }else {
        s3 := strings.Trim(s2," ")
        singleStr := strings.Index(s3," ")
        if singleStr <= -1 {
            return s3
        }
        }     
    return s4
}

//this ensures proper formatting from largest to smallest
func LargeIntLast (int1, int2 int) bool{
    if int1 < int2 {
        return true
    }else{
        return false
    }
}

//expands the range of ints
func MultRange (input string) string {
    num := strings.Index(input,"-")
    if num > -1 {
        splitnum := strings.Split(input,"-")
        if len(splitnum) > 2 {
            return ""
        }
        return input
    }else{
        return input
    }
}

//ensures only ints are used
func checkForWords (input string) bool{
    splitnum := strings.Split(input," ")
    for i := range(splitnum){
        num := splitnum[i]
        _,err := strconv.Atoi(num)
        if err != nil {
            return false
        }
    }
    return true
}
