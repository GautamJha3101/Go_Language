package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "INDIA"
	fmt.Println(strings.HasSuffix(s, "IA"))
	fmt.Println(strings.HasSuffix(s, "IND"))
	fmt.Println(strings.HasSuffix(s, "DIA"))
}
