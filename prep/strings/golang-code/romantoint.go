package main

import (
	"fmt"
)

// Convert Roman to Integer
// Ex: input : V output: 5
// logic :
// 1. make roman to int map
// 2. iterate over each char from string and check
//    2.1. if i smaller than n-1
//    2.2. if currenCharValue is smaller than nextCharValue then total will be total minus currenCharValue
//    2.3 else total will be total plus currenCharValue

func romanToInt(str string) int {

	romantoIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0

	for i := 0; i < len(str); i++ { // iterate over every char of str
		if i < len(str)-1 && romantoIntMap[str[i]] < romantoIntMap[str[i+1]] { // check till i is less than last element and value of roman char from str at ith pos is less than i+1th pos
			total = total - romantoIntMap[str[i]] // if above true then subtract value of roman char from total
		} else {
			total = total + romantoIntMap[str[i]] // else add value of roman char to total
		}
	}

	return total
}

func main() {
	fmt.Println(romanToInt("XXV"))
}
