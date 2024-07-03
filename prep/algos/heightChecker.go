package main

import (
	"fmt"
	"slices"
	"sort"
)

func heightChecker(heights []int) int {

	expectedHeights := []int{}

	for _, v := range heights {
		expectedHeights = append(expectedHeights, v)
	}

	slices.Sort(expectedHeights)
	sort.Ints(expectedHeights)

	count := 0

	for i := 0; i < len(expectedHeights); i++ {

		eh := expectedHeights[i]
		h := heights[i]

		if eh != h {
			count++
		}

	}

	return count

}

func main() {

	arr := []int{1, 1, 4, 2, 1, 3}
	//arr := []int{5,1,2,3,4}

	fmt.Println(heightChecker(arr))

}
