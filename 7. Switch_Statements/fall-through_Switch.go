package main

import "fmt"

func main() {
	var k = 30
	switch k {
	case 10:
		fmt.Print("<=10\n")
		fallthrough
	case 20:
		fmt.Print("<=20\n")
		fallthrough
	case 30:
		fmt.Print("<=30\n")
		fallthrough
	case 40:
		fmt.Print("<=40\n")
		fallthrough
	case 50:
		fmt.Print("<=50\n")
		fallthrough
	case 60:
		fmt.Print("<=60\n")
		fallthrough
	default:
		fmt.Print("Default Case")
	}
}
