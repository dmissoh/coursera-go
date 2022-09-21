package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

func main() {
	var name string
	var address string
	fmt.Scan(&name)
	fmt.Scan(&address)
	newPerson := Person{name, address}
	bytes, _ := json.Marshal(newPerson)
	fmt.Println(string(bytes))
}
