package main

import "fmt"

func findFactorial(n int) int {

	// base condition

	if n == 1 {
		return 1
	}

	return n * findFactorial(n-1)
}

func main() {
	res := findFactorial(5)

	fmt.Println(res)
}
