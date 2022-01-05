package main

import "fmt"

func main() {
	// nested for loop :- loop inside a loop
	for a := 0; a < 3; a++ {
		for b := 4; b > 0; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}

}
