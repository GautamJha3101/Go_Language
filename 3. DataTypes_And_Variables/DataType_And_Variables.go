package main

import "fmt"

func main() {
	var i = 36
	var f = 25.8
	var s = "Hello"

	// we can also declare all of them in one line like:-
	var a, b, c = 36, 25.8, "Hello"

	fmt.Println("i = ", i)
	fmt.Println("f = ", f)
	fmt.Println("s = ", s)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	// to know the data types of the variable we can use printf and "%T"

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("s is of type %T\n", s)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

}
