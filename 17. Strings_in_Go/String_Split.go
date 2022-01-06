package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I,love,my,country"
	var arr []string = strings.Split(str, ",")
	var len = len(arr)

	for a := 0; a < len; a++ {
		fmt.Println("Indes : ", a, " Value : ", arr[a])
	}
}
