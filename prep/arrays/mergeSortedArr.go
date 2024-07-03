package main

import "fmt"

func mergeSorted(n1, n2 []int, m, n int) {

	i := m
	for _, num2 := range n2 {
		if i < m+n {
			fmt.Println(i, ":", num2)
			n1[i] = num2
			i++
		}
	}

	fmt.Println(n1)

}

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	mergeSorted(nums1, nums2, m, n)

}
