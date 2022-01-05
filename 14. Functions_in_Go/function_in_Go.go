package main

import "fmt"

func nmeage(name string, age int) {
	fmt.Println("Your name is ", name)
	fmt.Println("Your age is ", age)
}
func main() {
	var name string
	var age int
	fmt.Println("Enter Your name ")
	fmt.Scanln(&name)
	fmt.Println("Enter Your age")
	fmt.Scanln(&age)

	nmeage(name, age)
}
