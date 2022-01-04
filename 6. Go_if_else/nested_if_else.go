package main

import "fmt"

func main() {

	var x int = 10
	var y int = 20

	if x >= 10 {
		fmt.Print("In outer if condtion statement\n")
		if y >= 10 {
			fmt.Print("In nested if statement\n")
		}
	}

	fmt.Println("Value of x is ", x)
	fmt.Println("Value of y is ", y)
}
