package main

import "fmt"

func main() {
	var number = 10

	var squarenum = func() int {
		number = number * number
		return number
	}

	fmt.Println(squarenum())
	fmt.Println(squarenum())
	fmt.Println(squarenum())

}
