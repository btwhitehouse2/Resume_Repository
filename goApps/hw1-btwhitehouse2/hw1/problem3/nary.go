package problem3

import (
	"strings"
)

// You cannot modify the name of this struct
type Node struct {
	point string
}
//****This program tracks the change in nAry while using the assigned data structure****
//You cannot modify this data structure at all (including its types). However 
// You are not required to use all of its fields. 
type Tree struct {
	root  *Node
	nAry  int
	nodes map[string]*Node
}

//function to add parent node to treeMap
//func (t *Tree) AddNode(nSmall *Node, nLarge Node, treeMap *map[string]Tree){
	func AddNode(nSmall *Node, nLarge *Node, treeMap *map[string]Tree){
		nMap := make(map[string]*Node)
		curnAry := 1
		nMap[nSmall.point]=nLarge
		t := Tree{nSmall,curnAry,nMap}
		tMP := *treeMap
		tMP[nSmall.point] = t
	}

//func (t *Tree) AddEdge(nSmall *Node, nLarge Node, treeMap *map[string]Tree){
	func AddEdge(nSmall *Node, nLarge *Node, treeMap *map[string]Tree){
		//nMap := make(map[string]*Node)
		//is looking up the location for points already saved
		tMP := *treeMap
		curnAry := tMP[nSmall.point].nAry 
		curnAry = curnAry + 1
		curNMP := tMP[nSmall.point].nodes
		curNode:= curNMP[nSmall.point]
		str := nSmall.point
		nMap := map[string]*Node{str:nLarge,str:curNode}
	
		nMap[nSmall.point]=nLarge
		t := Tree{nSmall,curnAry,nMap}
		tMP[nSmall.point] = t
	}


// SearchTree takes in a tree and returns the number of internal nodes that have children that are greater
// than the arty of the tree, which is given by the second argument to the function.
// Note: Please make sure to read the assignment description for more details.
func SearchTree(edges []string, nAry int) int {
	//treeMap uses parent nodes to track the inputting of edges
	treeMap := make(map[string]Tree)	
	//solve the problem without using pointers
	counter := 0
	//loops through all of the edges in the input
	for _, s := range edges {
		edgeSplit := strings.Split(s," ")
		//split the edge into two points with pointers to both	
		lowPoint, highPoint := edgeSplit[0],edgeSplit[1]
		lowNode, highNode := Node{lowPoint}, Node{highPoint}
		//check to see if TreeMap already contains the low point for edge
		//If it does not exist initiate the node for treeMap
		x := treeMap[lowNode.point].nodes
		if x == nil {
			AddNode(&lowNode,&highNode,&treeMap)
		}else{
		//If the low point for the edge does exist then we need update the nodeMap to include this additional point
			AddEdge(&lowNode,&highNode,&treeMap)
		}
	}
	for _, values := range treeMap {
		x := values.nAry
		if x > nAry{
			counter += 1
		} 
	}
	return counter
	}
