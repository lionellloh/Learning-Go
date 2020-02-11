package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var userString string
	fmt.Printf("Please enter a string: ")
	fmt.Scan(&userString)
	if len(userString) < 3 {
		fmt.Println("Not Found!")
		return
	}

	if ((strings.ToLower(string(userString[0])) == "i") && (strings.ToLower(string(userString[len(userString) -1])) == "n")) {
		for _, char := range userString {
			if unicode.ToLower(char) == 'a' {
				fmt.Println("Found!")
				return
			}
		}
		fmt.Println("Not Found!")
		return
	} else {
		fmt.Println("Not Found!")
	}
}
