package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x string = "Hello Gautam"
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
}
