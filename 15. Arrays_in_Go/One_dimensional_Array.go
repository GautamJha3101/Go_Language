package main

import "fmt"

func main() {
	var arr [5]int
	var i, j int

	for i = 0; i < 5; i++ {
		arr[i] = i + 50
	}

	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, arr[j])
	}
}
