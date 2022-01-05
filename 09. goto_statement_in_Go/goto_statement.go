package main

import "fmt"

func main() {
	var age int
	fmt.Print("Please enter your age ")
ENTER:
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("You are eligible to Vote")
	} else {
		fmt.Println("You are not eligible to vote")
		fmt.Println("Please Re-enter Your age")
		goto ENTER
	}
}
