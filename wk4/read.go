package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type person struct{
	fname, lname string
}

func main(){
	var filename string
	fmt.Printf("Filename: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	defer file.Close()

	output := []person{}
	if err == nil {
		reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')
		str_arr := strings.Split(line, " ")
		if len(str_arr) == 2 {
			p := person{str_arr[0], str_arr[1]}
			output = append(output, p)
		} else {
			break
		}

	}
	fmt.Println("Printing from struct <person> now")
	for _, e := range output {

		fmt.Print(e.fname, " ", e.lname)
	}

	} else {
		fmt.Println("Error: ", err)
	}
}
