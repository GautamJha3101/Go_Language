package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Value of x before using changeX function ", x)
	changeX(&x)
	fmt.Println("Value of x after using changeX function ", x)
}
func changeX(x *int) {
	*x = 0
}
