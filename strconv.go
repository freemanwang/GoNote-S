package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){

	i := 95
	ch := strconv.Itoa(i)   // int 转 string 一定成功，所以返回值仅一个 string
	fmt.Printf("ch 的类型是 %T，打印ch: %s\n", ch, ch)

	ch = "32"
	i, _ = strconv.Atoi(ch) // string 转 int 有可能不成功，返回值有2个： value, error

	fmt.Println(i) // 本次转换会成功

	ch = "a"
	i, err := strconv.Atoi(ch)   //转换会失败，因此用 err 做成功与否的判断

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	i2, _ := strconv.ParseInt("-42", 10, 64)   // -42
	fmt.Println(i2)

	// 声明一个slice
	b10 := []byte("int (base 10):")

	// 将转换为10进制的string，追加到slice中
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

	ctime := time.Now().Unix()
	fmt.Println(ctime)

}
