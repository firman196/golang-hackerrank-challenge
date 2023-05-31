package main

import (
	"fmt"
	"sort"
)

// sorting by compare old sum value with new sum value to search max and min values
func miniMaxSum1(arr []int32) {
	var max, min int64
	max, min = 0, 0
	for idxVal, _ := range arr {
		var sum int64 = 0
		for idxNesVal, nesVal := range arr {
			if idxVal != idxNesVal {
				sum += int64(nesVal)
			}
		}
		if max <= sum {
			max = sum
		}
		if min >= sum || min == 0 {
			min = sum
		}
	}
	fmt.Println(min, max)
}

// using array and sorting sum value
func miniMaxSum2(arr []int32) {
	var arrSum []int
	for idxVal, _ := range arr {
		var sum int
		for idxNesVal, nesVal := range arr {
			if idxVal != idxNesVal {
				sum += int(nesVal)
			}
		}
		arrSum = append(arrSum, sum)
	}
	sort.Ints(arrSum)
	fmt.Println(arrSum[0], arrSum[len(arr)-1])
}

/*func main() {
	miniMaxSum1([]int32{5, 5, 5, 5, 5})
	miniMaxSum2([]int32{5, 5, 5, 5, 5})
}*/
