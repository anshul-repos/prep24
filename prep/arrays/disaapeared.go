package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {

	res := []int{}

	presentCount := make([]bool, len(nums)+1)
	for _, n := range nums {
		presentCount[n] = true
	}

	for i := 1; i < len(nums); i++ {
		if !presentCount[i] {
			res = append(res, i)
		}
	}

	return res
}

func main() {

	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(arr))

}
