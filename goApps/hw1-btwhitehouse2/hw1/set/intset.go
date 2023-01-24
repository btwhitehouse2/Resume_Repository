package set
import ("fmt"
		"sort")

// IntSet represents a set of integers
type IntSet struct {
	intSeq []int // Uses an []int slice to hold the collection of integers that are part of the set.
}

// NewIntSet returns a newlly allocated IntSet with its contents initilaized to zero
func NewIntSet() *IntSet {
	return &IntSet{[]int{}}
}

func (recv *IntSet) PrintIntSeq() {
	fmt.Printf("{")
	for i,val := range recv.intSeq {
		fmt.Printf("%d",val)
		if i != len(recv.intSeq)-1{
			fmt.Printf(",")
		}
	}
	fmt.Printf("}\n")
}

//Adds single int to IntSet after confirming no duplicates
func (recv *IntSet) Add(num int){
	notDuplicate := true
 	for _,val := range recv.intSeq {
		if num == val {
			notDuplicate = false
		}
	}
	if notDuplicate == true {
		recv.intSeq = append(recv.intSeq,num)
	}
}

//Utilizing add, which already confirms no duplicates it adds both sets to a new set
func (recv *IntSet) Union(other *IntSet) *IntSet{
	set3 := NewIntSet()
	for _,val := range recv.intSeq{
		set3.Add(val)
	}
	for _,val := range other.intSeq{
		set3.Add(val)
	}
	sort.Ints(set3.intSeq)
	return set3

}

//uses map int of ints to count the number of times an int has been seen
func (recv *IntSet) Intersect(other *IntSet) *IntSet{
	set3 := NewIntSet()
	mapArray := map[int]int{}
	for _,val := range recv.intSeq {
		mapArray[val] = 1
	}
	for _,val := range other.intSeq {
		if mapArray[val] == 1 {
			set3.Add(val)
		}else{
			continue
		}
	}
	return set3
}

//uses 3 for loops. 1st to build a count from the 1st set
//second loop will increase the count by. 3rd loop checks the 1st set of values now that
//the second set have been inputted
func (recv *IntSet) Diff(other *IntSet) *IntSet{
	set3 := NewIntSet()
	mapArray := map[int]int{}
	for _,val := range recv.intSeq {
		mapArray[val] = 1
	}
	for _,val := range other.intSeq {
		mapArray[val] += 1
	}
	for _,val := range recv.intSeq {
		if mapArray[val] == 1{
			set3.Add(val)}
		}
	return set3
}
