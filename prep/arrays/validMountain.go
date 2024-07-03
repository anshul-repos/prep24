package main

import "fmt"

func validMountainArray(arr []int) bool {

	k := 0
	max := arr[0]
	for i, n := range arr {

		if n > max {
			max = n
			k = i
		}
	}

	if k == 0 || k == len(arr)-1 {
		return false
	}

	for i := 0; i < k; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	for i := k; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}

	return true

}

func main() {

	arr := []int{0, 3, 2, 1}
	res := validMountainArray(arr)
	fmt.Println(res)

}
