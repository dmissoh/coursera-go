package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {

	var userString string
	var userInt int

	var userSlice = make([]int, 0)

	for {

		// ask for user input
		fmt.Printf("Please enter an integer: ")
		_, err :=
			fmt.Scan(&userString)

		// handle error
		if err != nil {
			log.Fatal(err)
		}

		if userString == "X" || userString == "x" {
			fmt.Println("Quit the loop!")
			break
		} else {
			userInt, err = strconv.Atoi(userString)
		}

		userSlice = append(userSlice, userInt)

		sort.Ints(userSlice)
		

		fmt.Printf("Content of the sorted sliceContent of the sorted slice is: %v", userSlice)
		fmt.Println()

	}

}
