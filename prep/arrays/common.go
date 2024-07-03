package main

import (
	"fmt"
)

func common(nums []int, k int) []int {

	nCountMap := make(map[int]int)
	result := []int{}

	for _, n := range nums {
		nCountMap[n]++
	}

	reverseCountMap := map[int][]int{}
	for n, count := range nCountMap {
		reverseCountMap[count] = append(reverseCountMap[count], n)
	}

	for i := len(nums); len(result) != k; i-- {

		for _, n := range reverseCountMap[i] {
			if len(result) != k {
				result = append(result, n)
			}
		}
	}

	return result

	//for n, count := range nCountMap {
	//	if count >= k {
	//		result = append(result, n)
	//	}
	//}

	//return result
}

func main() {
	arr := []int{1, 2} // [1,2]

	//common(arr)

	res := common(arr, 2)
	fmt.Println(res)
}
