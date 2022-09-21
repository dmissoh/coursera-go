package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("Please enter integers speratared by space: ")

	in := bufio.NewReader(os.Stdin)
	userInput, _ := in.ReadString('\n')

	userInput = strings.TrimSpace(userInput)
	userInput = strings.Trim(userInput, "\n")

	digits := strings.Split(userInput, " ")

	var slice []int = convertStringsToInts(digits)

	BubleSort(slice)
	fmt.Println("Sorted list of integer numbers using bubble sort:")
	fmt.Println(slice)
}

func convertStringsToInts(digitsAsString []string) []int {
	var slice []int
	for _, digit := range digitsAsString {
		number, _ := strconv.Atoi(digit)
		slice = append(slice, number)
	}
	return slice
}

func BubleSort(ints []int) {
	for i := 0; i < len(ints)-1; i++ {
		for j := 0; j < len(ints)-i-1; j++ {
			if ints[j] > ints[j+1] {
				Swap(ints, j)
			}
		}
	}
}

func Swap(ints []int, index int) {
	nextIndex := index + 1
	if index < len(ints) && (nextIndex) < len(ints) {
		current := ints[index]
		next := ints[nextIndex]
		ints[index] = next
		ints[nextIndex] = current
	}
}
