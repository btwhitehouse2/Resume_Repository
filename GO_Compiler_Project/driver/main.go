package main

import "fmt"

func isSqrt(v int) int {
	return v
}

func main() {
	//Calculates distance between two integer points
	var pt1X, pt2X int
	var pt1Y, pt2Y int
	var xSquared, ySquared int

	xSquared = (pt2X - pt1X) * (pt2X - pt1X)
	ySquared = (pt2Y - pt1Y) * (pt2Y - pt1Y)

	//Assume we have an integer square root function
	fmt.Printf("distance=%d", isSqrt(xSquared+ySquared))

}
