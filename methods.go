package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

type Pet string

func (pet Pet) bark() {
	fmt.Printf("%s bark! \n", pet)
}

func (person Person) sayHi() {
	fmt.Printf("Hi~~, I'm %s ! \n", person.name)
}


type DoubleNums struct {
	X, Y float64
}

func (v DoubleNums) add() float64 {
	return v.X + v.Y
}

//func (v *DoubleNums) Scale(f float64) {
func (v DoubleNums) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}





func main() {
	a := Person{"Alice", 18}
	a.sayHi()

	p := Pet("dog")
	p.bark()

	v := DoubleNums{3, 4}
	v.Scale(10)
	fmt.Println(v.add())
}
