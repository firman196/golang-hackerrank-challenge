package main

func breakingTheRecords(scores []int32) []int32 {
	result := []int32{0, 0}
	max, min := scores[0], scores[0]
	for i := 1; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
			result[0] += 1
		}

		if scores[i] < min {
			min = scores[i]
			result[1] += 1

		}

	}
	return result
}

/*
func main() {
	fmt.Println(breakingTheRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
	fmt.Println(breakingTheRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
}
*/
