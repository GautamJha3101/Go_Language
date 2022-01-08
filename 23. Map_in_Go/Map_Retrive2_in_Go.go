package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 10
	m["key2"] = 20
	m["Key3"] = 30
	fmt.Println(m)

	v, ok := m["key2"]
	fmt.Println("The value: ", v, "Present? ", ok)
	i, j := m["Key4"]
	fmt.Println("he Value: ", i, "Present? ", j)
}
