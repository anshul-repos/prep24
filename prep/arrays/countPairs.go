package main

import "fmt"

func countPairs(nums []int, target int) int {

	count := 0

	// naive/bruteforce approach

	//for i := 0; i < len(nums); i++ {
	//	for j := 0; j < len(nums); j++ {
	//		sum := nums[i] + nums[j]
	//		if i < j {
	//			if sum < target {
	//				count++
	//			}
	//		}
	//	}
	//}

	// two pointer approach
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]+nums[right] < target {
			count += right - left
			left++
		} else {
			right--
		}
	}

	return count

}

func main() {

	ums := []int{9, -5, -5, 5, -5, -4, -6, 6, -6}
	target := 27

	res := countPairs(ums, target)

	fmt.Println(res)

}
