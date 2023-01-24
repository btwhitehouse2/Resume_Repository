package problem2

import (
	"math"
	"strings"
	"strconv"
)

// ProcessLog returns the total number of users from the log entries that are greater than the maxDuration time
// Note: Please make sure to read the assignment description for more details.
type pLog struct {
	//visitor_num string
	time_stamp float64
	action string
}

func ProcessLog(log []string, maxDuration int) int {
	c:= make(map[string]pLog)
	//value that we will return at the end with the number of instances
	//less than or greater than the allowable time
	var counter int
	for _, s := range log {
		individual_log := strings.Replace(s,","," " , -1)
		individual_log2 := strings.Fields(individual_log)
		vNum := individual_log2[0]
		intNew,_ := strconv.ParseFloat(individual_log2[1],64)
		acT := individual_log2[2]
		val, ok := c[vNum]
		//checking to see if the value already exists in c
		//if the value does exist need to confirm the difference in time
		if ok{
			a := val.time_stamp
			time_diff := math.Abs(a - intNew)
			if time_diff <= float64(maxDuration){
				counter += 1
			}
			//if the value exceeds the time we are not concerned
			//no need to re-write the map for the different value because we know the person will only come once
			//c[vNum] = pLog{intNew, acT}
		}
		if ! ok{
			c[vNum] = pLog{intNew, acT}
		}	
	}
	return(counter)
}
