package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// ask for user input
	fmt.Printf("Please enter a text to analyse: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userString := scanner.Text()

	var userInputLowerCase = strings.ToLower(userString)

	// validate user input
	var startsWithI bool = strings.HasPrefix(userInputLowerCase, "i")
	var endsWithN bool = strings.HasSuffix(userInputLowerCase, "n")
	var containsA bool = strings.Contains(userInputLowerCase, "a")

	var found bool = startsWithI && endsWithN && containsA

	if found {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
