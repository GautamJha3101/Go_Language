package main

import "fmt"

func main() {
	// an array with 3 rows and 3 column
	var arr = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	var i, j int
	// Output of each array elements
	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
