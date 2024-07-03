package main

import "fmt"

//Given an integer array nums and an integer k,
//return true if there are two distinct indices i and j in the array such that
// 1. nums[i] == nums[j]
// 2. abs(i - j) <= k.

func containsNearbyDuplicate(nums []int, k int) bool {

	track := make(map[int]int)

	for i, v := range nums {

		j, found := track[v]

		if found && abs(i-j) <= k {
			return true
		}

		track[v] = i
	}

	return false

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	arr := []int{1, 0, 1, 1}
	k := 1

	//arr := []int{1, 2, 3, 1, 2, 3}
	//k := 2

	fmt.Println(containsNearbyDuplicate(arr, k))
}
