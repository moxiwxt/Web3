package main

import "fmt"

func increase(pointer *int) {
	*pointer += 10
}

func main() {

	i := 103
	poin := &i
	for i < 200 {
		increase(poin)
		fmt.Println(i)
	}

}
