package main

import (
	"fmt"
	"os"
)

func main() {
	var argumentat1, argumentat2 string

	argumentwithpath := os.Args // return all arguments including path
	argumentat1 = os.Args[1]    // return argumant after path
	argumentat2 = os.Args[2]    // return specified argument only

	fmt.Println(argumentwithpath)
	fmt.Println(argumentat1)
	fmt.Println(argumentat2)
}
