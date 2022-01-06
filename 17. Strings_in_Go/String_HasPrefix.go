package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "INDIA"
	fmt.Println(strings.HasPrefix(s, "IN"))
	fmt.Println(strings.HasPrefix(s, "DI"))

}
