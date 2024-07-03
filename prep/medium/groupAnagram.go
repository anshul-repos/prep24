package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortStr(str)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	return anagramMap
}

func sortStr(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {

	strs := []string{"eat", "tea", "ups", "sup", "run", "pus"}

	res := groupAnagrams(strs)

	for _, grp := range res {

		fmt.Println(grp)
	}

}
