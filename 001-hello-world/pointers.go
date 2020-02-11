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
	h()
	conditional()
	for_loop()
	scan_func()
}

var x = "3x lol"

func f(){
	var x = 5
	fmt.Println("%d", x)
}

func g(){
	fmt.Println(x)
}

func h(){
	//type conversion
	var x int32 = 1
	var y int16 = 2
	x = int32(y)
	fmt.Println(x)

	var a float64 = 123.45
	var b float64 = 1.23e3
	fmt.Println(a, b)

	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		F

	)

	fmt.Println(A)
	fmt.Println(B)

//	UTF-8 is variable length. (7) 8-bit UTF codes are same as ASCII
// A rune is a UTF 8 code point
}

func conditional(){
	var x int = 3
	if x > 2 {
		fmt.Println("x is more than 2")
		fmt.Println(x)
	}
}

func for_loop(){
	for i:=0; i < 10; i++ {
		fmt.Println(i);
	}

	var j = 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
}

func switch_func(){
	x:=1
	switch {
	case x > 1:
		fmt.Printf("case 1")

	case x < 1:
		fmt.Printf("case2")


	default:
		fmt.Printf("default")
	}
}

func scan_func(){
	var appleNum int
	fmt.Printf("How mnany apples?")
	fmt.Scan(&appleNum)
	fmt.Println(appleNum)
}