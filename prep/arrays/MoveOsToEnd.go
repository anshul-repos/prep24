package main

import "fmt"

func moveZeros(nums []int) {

	nonZeroPositionPointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroPositionPointer] = nums[nonZeroPositionPointer], nums[i]
			nonZeroPositionPointer++
		}
	}
	fmt.Println(nums)

}

func main() {

	nums := []int{0, 1, 0, 3, 12} // output : [1,3,12,0,0]

	moveZeros(nums)

}
