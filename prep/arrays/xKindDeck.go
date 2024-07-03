package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {

	xKindsMap := make(map[int][]int)

	for _, n := range deck {
		if _, matched := xKindsMap[n]; matched {
			xKindsMap[n] = append(xKindsMap[n], n)
		} else {
			xKindsMap[n] = []int{n}
		}
	}

	pairSize := 0
	for _, pair := range xKindsMap {
		if pairSize == 0 {
			pairSize = len(pair)
		} else if len(pair) != pairSize {
			return false
		}
	}

	return pairSize >= 2

}

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(hasGroupsSizeX(arr))

}
