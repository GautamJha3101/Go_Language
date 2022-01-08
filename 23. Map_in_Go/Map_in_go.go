package main

import "fmt"

func main() {
	x := map[string]int{
		"Kate": 28,
		"Jhon": 37,
		"Raj":  20,
	}
	fmt.Println(x)
	fmt.Println("\n", x["Raj"])
}
