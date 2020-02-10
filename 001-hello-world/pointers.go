package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip
	fmt.Println(ip)
	fmt.Println(y)

	ptr := new(int)
	*ptr = 3

	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)
	fmt.Println(*&ptr)

	f()
	g()
}

var x = "3x lol"

func f(){
	var x = 5
	fmt.Println("%d", x)
}

func g(){
	fmt.Println(x)
}