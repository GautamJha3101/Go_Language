package main

import "fmt"

func main() {
	var sum = 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

}
