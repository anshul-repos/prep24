package main

import (
	"fmt"
	"slices"
)

func lastStoneWeight(stones []int) int {

	n := len(stones)

	if n == 1 {
		return stones[0]
	}

	slices.Sort(stones)

	// difference of last and second last element
	diff := stones[n-1] - stones[n-2]

	// trim last two
	stones = stones[:n-2]

	// append diff into stones array
	stones = append(stones, diff)

	return lastStoneWeight(stones)
}

func main() {
	arr := []int{2, 2}
	fmt.Println(lastStoneWeight(arr))

}
