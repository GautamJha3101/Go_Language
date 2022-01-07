package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	age       int
}

func main() {
	x := person{
		age:       30,
		FirstName: "Gautam",
		LastName:  "Jha",
	}

	fmt.Println(x)
	fmt.Println(x.FirstName)
}
