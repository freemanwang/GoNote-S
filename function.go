package main

import "fmt"

//func add(x, y, z int, str string ) (int, string, string){
//	return  x + y + z, "     ", str
//}

func forLoop(x int) int {
	sum := 0
	for i:= 0; i <= x; i++ {
		sum += i
	}
	return sum
}

func Count100(){
	count := 100
	sum := 0
	for ; count > 0; {
		count --
		sum ++
	}
	fmt.Printf(" 100 个 1 相加的和为：%d", sum)
}
func sayHi( device string) string {
	if device == "SamSung Fold" {
		return "awesome!这是我独享的moment！！"
	}
	return "Hi there~"
}

func add(x, y int) int {
	return x + y
}
func compute(fn func(x, y int) int) int{
	return fn(3, 4)
}


func main() {
	ch := 'a'
	fmt.Println(ch)
	//Count100()
	fmt.Println(sayHi("SamSung Fold"))

	addFunc := add

	fmt.Println(addFunc(5,10))

	fmt.Println(compute(addFunc))
}

