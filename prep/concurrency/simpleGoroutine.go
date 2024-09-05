package main

import (
	"fmt"
	"sync"
)

func generateNums(n int, ch chan int, wg *sync.WaitGroup) {
	for i := 0; i <= n; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func main() {

	ch := make(chan int, 1)
	var wg sync.WaitGroup
	n := 50

	wg.Add(1)
	go generateNums(n, ch, &wg)
	for num := range ch {
		fmt.Println(num)
	}
	wg.Wait()

	fmt.Println("EXIT")
}
