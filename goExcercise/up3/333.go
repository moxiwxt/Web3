package main

import (
	"fmt"
	"sync"
)

func printOdd(wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go printOdd(&wg)
	go printEven(&wg)

	wg.Wait()

}
