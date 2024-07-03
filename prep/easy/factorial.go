package main

import "fmt"

//  factorial of number
//  5! = 5 * 4 * 3 * 2 * 1
// = 120

// recursion approach
// n x n-1 x n-2 x n-3

func factorial(currentNum int) int {

	if currentNum == 1 {
		return 1
	}

	fact := currentNum * factorial(currentNum-1)

	return fact
}

func main() {
	fmt.Println(factorial(5))
}
