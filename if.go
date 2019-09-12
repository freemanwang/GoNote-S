package main

import (
	"add/myMath"
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return v - lim
	}
	//fmt.Println(v)   这里不注销会报错，undefined: v
	return lim
}

func scores(x int){
	if x < 0  || x > 100 {
		fmt.Println("输入非法")
	} else if x < 60 {
		fmt.Println("不及格")
	} else if x >= 60 && x <= 80 {
		fmt.Println("及格")
	} else {
		fmt.Println("优秀")
	}
}

func scoresSwitch (x int){
	switch x/10 {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("不及格")
	case 6:
		fallthrough
	case 7:
		fallthrough
	case 8:
		fmt.Println("及格")
	case 9:
		fallthrough
	case 10:
		fmt.Println("优秀")
	default:
		fmt.Println("input error")
	}
}

func switch2(x int){
	switch  {
	case x >= 0 && x < 60 :
		fmt.Println("不及格")
	case x >= 60 && x < 80:
		fmt.Println("及格")
	case x >= 80 && x <= 100:
		fmt.Println("优秀")
	default:
		fmt.Println("输入非法")
	}
}
func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(5, 2, 20),
	)
	scores(-60)
	scores(60)
	scores(80)
	scores(100)

	scoresSwitch(50)
	scoresSwitch(90)
	scoresSwitch(70)
	scoresSwitch(120)

	switch2(50)
	switch2(70)
	switch2(90)
	switch2(190)

	fmt.Println(myMath.Add(1,2))
}

