package main

import "fmt"

func quickSort(arr []int) []int {

	// base case : array can have minimum 1 element or 0 element i.e less than 2
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := []int{}
	greater := []int{}

	for _, val := range arr[1:] {
		if val <= pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	sortedLess := quickSort(less)
	sortedGreater := quickSort(greater)

	res := append(append(sortedLess, pivot), sortedGreater...)

	return res
}

func main() {

	arr := []int{5, 2, 8, 1, 6, 7, 3, 9, 4}
	fmt.Println(quickSort(arr))

}
