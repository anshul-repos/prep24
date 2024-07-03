package main

import "fmt"

//String : abc
//Subsequences : a, b, c, ab, bc, ac, abc
//they should be in order that's it.

// recursion approach:
//choice 1. Exclude the current char and move next
//choice 2. Include the current char and move next

func subSeqStr(str, currentStr string, index int, result *[]string) {

	if index == len(str) {
		*result = append(*result, currentStr)
		return
	}

	// exclude current character and move to next
	subSeqStr(str, currentStr, index+1, result)

	// include current character and move next
	subSeqStr(str, currentStr+string(str[index]), index+1, result)

}

func printSubsequence(str string) {
	result := []string{}

	subSeqStr(str, "", 0, &result)
	for _, subSeq := range result {
		fmt.Println(subSeq)
	}
}

func main() {
	printSubsequence("abc")
}
