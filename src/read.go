package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {

	var fname string
	var namesSlice = make([]name, 0)

	fmt.Printf("Please enter the file name: ")
	fmt.Scanln(&fname)

	readFile, err := os.Open(fname)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var line = fileScanner.Text()
		split := strings.Split(line, " ")
		var firstName = split[0]
		var lastName = split[1]
		var nameStructure = name{fname: firstName, lname: lastName}
		namesSlice = append(namesSlice, nameStructure)
	}

	readFile.Close()

	fmt.Println()
	for _, value := range namesSlice {
		fmt.Printf("First name is '%s' and last name is '%s'", value.fname, value.lname)
		fmt.Println()
	}
	//fmt.Println(person{fname: "Claudia", lname: "Westhus"})
}
