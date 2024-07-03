package main

import (
	"fmt"
)

func commonPrefix(prefix, str string) string {

	var minLen int

	if len(prefix) > len(str) {
		minLen = len(str)
	} else {
		minLen = len(prefix)
	}

	for i := 0; i < minLen; i++ {
		if prefix[i] != str[i] {
			return prefix[:i]
		}
	}

	return prefix[:minLen]
}

func longestCommonPrefix(strs []string) {

	//sort.Strings(strs)
	//fmt.Println(strs)

	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
		if prefix == "" {
			break
		}
	}

	fmt.Println(prefix)
}

func main() {
	longestCommonPrefix([]string{"top", "topper", "topandtown"})
}
