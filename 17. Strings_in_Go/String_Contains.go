package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi...there"
	fmt.Println(strings.Contains(str, "th"))
	fmt.Println(strings.Contains(str, "abc"))
}
