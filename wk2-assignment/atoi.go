package main

//func main() {
//	i, _ := strconv.Atoi("10")
//	fmt.Println()
//	y := i*2
//	fmt.Println(y)
//}

//func main() {
//	x:=7
//	switch {
//	case x>3:
//		fmt.Printf("1")
//	case x>5:
//		fmt.Printf("2")
//	case x==7:
//		fmt.Printf("3")
//	default:
//		fmt.Printf("4")
//	}
//}

//func main() {
//	var xtemp int
//	x1 := 0
//	x2 := 1
//	for x:=0; x<5; x++ {
//		xtemp = x2
//		x2 = x2 + x1
//		x1 = xtemp
//	}
//	fmt.Println(x2)
//}

func main() {
	var x int
	var y *int
	z := 3
	y = &z
	x = &y
}