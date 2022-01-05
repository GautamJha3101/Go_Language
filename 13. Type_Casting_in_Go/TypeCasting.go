package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "10.123"

	fmt.Println(float64(i)) // converted int to float
	fmt.Println(int(f))     // converted float to int

	newInt, _ := strconv.ParseInt(str1, 0, 64) // string converted to int
	fmt.Println(newInt)

	newFloat, _ := strconv.ParseFloat(str2, 64) // string converted to float

	fmt.Println(newFloat)
}
