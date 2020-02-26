package main

import (
	"fmt"
	"time"
)

func main(){
	//fmt.Printf("%d", time.Millisecond)
	go spinner()
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}


func fib(n int) int{

	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
func spinner(){
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100000000)
		}
	}
}