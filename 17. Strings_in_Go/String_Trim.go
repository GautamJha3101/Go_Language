package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "\t\n I love my country \n\t\r\n"
	fmt.Println(s)
	fmt.Println()
	fmt.Println("After using trim  Function")
	fmt.Println(strings.TrimSpace(s))
}
