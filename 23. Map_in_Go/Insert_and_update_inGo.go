package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)
	m["Key1"] = 10
	m["key2"] = 20
	m["key3"] = 30
	fmt.Println(m)
	m["Key2"] = 5555

	fmt.Println(m)
}
