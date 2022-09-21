package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var address string

	var nameAddress = make(map[string]string)

	fmt.Printf("Please enter your name: ")
	fmt.Scanln(&name)

	fmt.Printf("Please enter your address: ")
	fmt.Scanln(&address)

	nameAddress["name"] = name
	nameAddress["address"] = address
	marshaledJson, _ := json.MarshalIndent(nameAddress, "", "   ")

	fmt.Println("JSON object is: ")
	fmt.Println(string(marshaledJson))
}
