package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println((*p).X)
	fmt.Println(p.Y)
	p.Y = 31415926
	fmt.Println(p.Y)
}

