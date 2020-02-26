package main

import "fmt"

/*
var clock chan int

*/

func f(n int, search []int, key int){
	for i:=0; i<len(search); i++{
		fmt.Println(n, ":", i, "is searching")
		if search[i] == key {
			fmt.Println(i, " search is successful")
		}
	}
}

