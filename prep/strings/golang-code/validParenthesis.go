package main

import "fmt"

// check if string is Valid parenthesis i.e
// Ex: 1. [] > valid 2. {{}}] > invalid

func isValidParenthesis(str string) bool {

	stack := []rune{}

	bracketsMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range str {

		switch char {
		case '(', '{', '[':
			// Push bracket into stack
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bracketsMap[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	testcases := []string{"[]{}((", "[](){}", "([)]"}
	for _, str := range testcases {
		fmt.Println()
		fmt.Printf("%s IsValid Parenthesis : %v ", str, isValidParenthesis(str))
		fmt.Println()
	}
}
