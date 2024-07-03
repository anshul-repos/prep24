package main

import (
	"fmt"
)

//A string is called a palindrome if the reverse of the string is the same as the original one.
//Example: “madam”, “racecar”, “12321”.

// my logic
// start from right and left i.e i:= 0 and i++,j:=len(string)-1 and j--
//	if i and j = mid break else continue string(i) = string(j)

func isPalindrome(str string) bool {

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}

func main() {
	str := "maninam"
	fmt.Println(isPalindrome(str))
}
