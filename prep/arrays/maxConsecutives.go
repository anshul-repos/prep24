package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {

	countOccurance := 0
	maxCount := 0

	for _, n := range nums {

		if n != 1 {
			if countOccurance > maxCount {
				maxCount = countOccurance
			}
			countOccurance = 0
			continue
		}
		countOccurance++
	}

	if countOccurance > maxCount {
		maxCount = countOccurance
	}

	return maxCount
}

func main() {
	arr := []int{1, 0, 1, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(arr))
}
