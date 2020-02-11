package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := make([]int, 3)

	for  {
		var u string
		fmt.Print("Please enter an integer: ")
		fmt.Scan(&u)
		//fmt.Println(n, err)
		if u == "X" {

			return
		} else{
			u, _ := strconv.Atoi(u)
			arr = append(arr, u)
			fmt.Println(arr)
		}

	}



}