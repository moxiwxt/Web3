package main

import (
	"fmt"
)

type Cat struct {
	name string
	age  int
}

func fetchChannel(ch chan Cat) {
	value := <-ch
	fmt.Printf("type: %T, value: %v\n", value, value)
}

func main() {

	array := []int{1, 1, 4, 4, 5, 5, 6, 7, 7, 10, 10}
	collectNum := map[int]int{}
	fmt.Println(len(array))
	for i := 0; i < len(array); i++ {
		value, isExist := collectNum[array[i]]
		if isExist {
			collectNum[array[i]] = value + 1
		} else {
			collectNum[array[i]] = 1
		}
	}

	fmt.Println(collectNum)
}
