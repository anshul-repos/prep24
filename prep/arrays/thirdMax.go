package main

import "fmt"

func thirdMax(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	countMap := make(map[int]bool)
	count := 0

	largest, secondLargest, smallest := nums[0], nums[1], 0 // 2,2,
	skipThird := false

	if len(nums) == 2 {
		skipThird = true
	} else {
		smallest = nums[2]
	}

	for _, n := range nums {
		if countMap[n] {
			continue
		}
		countMap[n] = true
		count++

		if n > largest {
			smallest = secondLargest
			secondLargest = largest
			largest = n
		} else if n < largest && n > secondLargest {
			smallest = secondLargest
			secondLargest = n
		} else if n < smallest && !skipThird {
			smallest = n
		}
	}

	if count < 3 {
		return largest
	}

	return smallest
}

func main() {
	arr := []int{5, 2, 4, 1, 3, 6, 0}

	res := thirdMax(arr)

	fmt.Println(res)

}
