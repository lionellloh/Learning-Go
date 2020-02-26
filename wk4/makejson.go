package wk4

import (
	"encoding/json"
	"fmt"
)
func main() {
	var name string
	var address string
	person := map[string]string{}
	fmt.Printf("Name: ")
	fmt.Scan(&name)
	fmt.Printf("Address: ")
	fmt.Scan(&address)
	person["name"] = name
	person["address"] = address

	data, err := json.Marshal(person)

	if err == nil {
		fmt.Println(data)
	}
}
