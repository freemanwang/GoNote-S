package main

import "fmt"

type Person struct {
	name string
	age int
}
type Pet string


func (p *Person) sayHi() string {
	return "Hi~ I'm " + p.name + ". Nice to meet you!"
}
func (p Pet) bark() string {
	return string(p + " bark!")
}


type Hi interface {
	sayHi() string
}
type Bark interface {
	bark() string
}


func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main(){
	var hi Hi
	var bark Bark
	person := Person{"Alice", 18}
	hi = &person   // 注意， person 是 *Person类的指针，因此 hi=&person   ;

	pet := Pet("Dog")
	bark = pet  // pet 是对象，因此直接 =    注意和  hi = &person 的区别

	fmt.Println(hi.sayHi())  //通过 interface 调用方法
	fmt.Println(bark.bark())  //通过 interface 调用方法

	a := Person{"Alice", 18}
	b := Person{"Bob", 20}
	c := Person{"Carel", 16}
	fmt.Println(a, b, c)
}