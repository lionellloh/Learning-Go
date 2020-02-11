package main

import (
	"fmt"
	"sort"
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
			u, err := strconv.Atoi(u)
			if err == nil {
				arr = append(arr, u)
				sort.Ints(arr)
				fmt.Println(arr)
			} else {
				fmt.Println(err)
			}
		}
	}

}