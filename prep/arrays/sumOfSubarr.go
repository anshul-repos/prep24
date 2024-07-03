package main

import "fmt"

func sumOfSubArr(arr []int, sizeOfSubArr int) {

	res := 0

	for i := 0; i <= len(arr)-sizeOfSubArr; i++ {
		for j := i; j < i+sizeOfSubArr; j++ { // subarray movement
			res = res + arr[j]
		}
	}

	fmt.Println(res)

}

func main() {

	arr := []int{2, 3, 5}
	sumOfSubArr(arr, 2)

}

// 5 + 8
// 13
