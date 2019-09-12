package main

import "fmt"

func main(){
	i := 99
	var p *int
	p  = &i
	//var p  = &i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 200
	fmt.Println(p)
	fmt.Println(*p)
}
