package main

import "fmt"

func yourage(age int) int {
	return age
}
func main() {
	var age int
	fmt.Println("Enter your age")
	fmt.Scan(&age)

	var myage = yourage(age)
	fmt.Println("Your age is ", myage)

}
