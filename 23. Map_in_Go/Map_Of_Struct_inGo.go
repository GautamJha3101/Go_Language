package main

import "fmt"

type Vertex struct {
	Latitude, Longitude float64
}

var m = map[string]Vertex{
	"XYZ": Vertex{40.68433, -74.39967},
	"ABC": Vertex{37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
