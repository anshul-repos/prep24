package main

import (
	"fmt"
	"time"
)

func printNumsA(n int) {
	// base condition > largest invalid input
	if n > 100 {
		return
	}

	fmt.Println(n)
	printNumsA(n + 1)
}

func printNumsB(n int) {

	// base condition > smallest valid input
	if n == 0 {
		fmt.Println(n)
		return
	}

	fmt.Println(n)

	printNumsB(n - 1)

}

func main() {

	printNumsA(51)
	time.Sleep(time.Second * 3)
	printNumsB(50)
}
