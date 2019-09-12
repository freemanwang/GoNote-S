package main

import (
	"log"
	"time"
)

func foo1(n int) int {
	log.Println("n1=", n)
	defer log.Println("n=", n)
	n += 100
	log.Println("n2=", n)
	return n
}

func foo2(n int) int {
	defer log.Println("1111")
	time.Sleep(1 * time.Second)
	defer log.Println("2222")
	time.Sleep(1 * time.Second)
	defer log.Println("3333")
	time.Sleep(1 * time.Second)
	return n
}

func main() {
	var i int = 100
	foo1(i)
	foo2(i)

}