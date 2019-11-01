package main

import (
	"fmt"
)

var (
	myRes =make(map[int]int,20)
)

func factorial(n int,ch chan map[int]int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myRes[n] = res
	ch<-myRes
}

func main() {
	a :=make(chan map[int]int,20)
	for i := 1; i <= 20; i++ {
		go factorial(i,a)
	}
	b:=<-a
	for i, v := range b {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}