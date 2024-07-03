package main

import "fmt"

func plusOne(arr []int) []int {

	n := len(arr)

	for i := n - 1; i >= 0; i-- {

		if arr[i] < 9 {
			arr[i]++
			return arr
		}

		arr[i] = 0

	}

	arr = append([]int{1}, arr...)

	return arr

}

func main() {

	arr := []int{1, 3, 5, 9}
	fmt.Println(plusOne(arr))

}
