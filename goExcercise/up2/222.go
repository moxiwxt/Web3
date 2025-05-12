package main

import "fmt"

func operSlice(intSlice *[]int) {

	for i := range *intSlice {
		(*intSlice)[i] *= 2
	}

}

func main() {

	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	operSlice(&intSlice)
	fmt.Println(intSlice)

}
