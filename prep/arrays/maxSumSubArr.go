package main

import (
	"container/list"
	"fmt"
)

//func maxSumOfSubArr(arr []int, k int) {
//
//	i, j := 0, 0
//
//	sum := 0
//
//	mx := 0
//
//	for i < len(arr)-k && j < len(arr) {
//
//		for j = 0; j <= k; j++ {
//			sum = sum + arr[j]
//
//			if j-i+1 < k {
//				j++
//			} else if j-i+1 == k {
//				if sum > mx {
//					mx = sum
//				}
//			}
//		}
//
//	}
//
//}
//
//func main() {
//
//	k := 3
//	arr := []int{2, 5, 1, 8, 2}
//
//	maxSumOfSubArr(arr, k)
//
//}

func slidingMaximum(v []int, k int) []int {
	var ans []int

	if k > len(v) {
		maxElement := v[0]
		for _, val := range v {
			if val > maxElement {
				maxElement = val
			}
		}
		return append(ans, maxElement)
	}

	i, j := 0, 0
	size := len(v)
	l := list.New()

	for j < size {
		// Calculations
		for l.Len() > 0 && l.Back().Value.(int) < v[j] {
			l.Remove(l.Back())
		}
		l.PushBack(v[j])

		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			// Append the max element in the current window to ans
			ans = append(ans, l.Front().Value.(int))

			// Slide the window
			if l.Front().Value.(int) == v[i] {
				l.Remove(l.Front())
			}
			i++
			j++
		}
	}

	return ans
}

func main() {
	v := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := slidingMaximum(v, k)
	fmt.Println(result) // Output: [3, 3, 5, 5, 6, 7]
}
