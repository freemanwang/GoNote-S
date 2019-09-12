package main

import "fmt"

type Human struct {
	name string
	age int
}

type Student struct {
	Human   //匿名字段，会把 Human 结构体所有字段引入 Student 结构体
	grade, class int
}

func (s Student) String() string {
	return fmt.Sprintf( "My name is %s, I am %d years old. I'm in grade %d class %d.\n", s.name, s.age, s.grade, s.class)
	//return fmt.Printf( "My name is %s, I am %d years old.", s.name, s.age, s.grade, s.class)
}

func main()  {
	a := Student{Human{"Alice", 18}, 2, 1}
	fmt.Println(a)
	fmt.Printf("name: %s \n", a.name)   // 可见 Human 内的 name 字段， 被当作了 Student 内部的字段，可直接用
	fmt.Printf("grade: %d \n", a.grade)  // name 可以和 Student 内定义的 grade 字段一样方式使用
}
