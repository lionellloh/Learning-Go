package main

import (
	"fmt"
)

//func main() {
//
//	arr := make([]int, 3)
//
//	for  {
//		var u string
//		fmt.Print("Please enter an integer: ")
//		fmt.Scan(&u)
//		//fmt.Println(n, err)
//		if u == "X" {
//
//			return
//		} else{
//			u, err := strconv.Atoi(u)
//			if err == nil {
//				arr = append(arr, u)
//				sort.Ints(arr)
//				fmt.Println(arr)
//			} else {
//				fmt.Println(err)
//			}
//		}
//	}
//
//}

//func main() {
//	x := [...]int {4, 8, 5}
//	y := x[0:2]
//	z := x[1:3]
//	y[0] = 1
//	z[1] = 3
//	fmt.Print(x)
//}

//func main() {
//	x := [...]int {1, 2, 3, 4, 5}
//	y := x[0:2]
//	z := x[1:4]
//	fmt.Print(len(y), cap(y), len(z), cap(z))
//}

func main() {
	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}
