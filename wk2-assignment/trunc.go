package main
import "fmt"

func main() {
	var userInput float64
	fmt.Printf("Please enter a floating point number")
	fmt.Scan(&userInput)
	fmt.Println(int(userInput))
}
