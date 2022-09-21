package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	channel := make(chan []int)

	fmt.Printf("Please enter integers speratared by space: ")

	in := bufio.NewReader(os.Stdin)
	userInput, _ := in.ReadString('\n')

	userInput = strings.TrimSpace(userInput)
	userInput = strings.Trim(userInput, "\n")

	digits := strings.Split(userInput, " ")

	var slice []int = convertStringsToInts(digits)

	chunkSize := (len(slice) / 4)
	chunks := chunkSlice(slice, chunkSize)

	go sortSlice(chunks[0], channel)
	go sortSlice(chunks[1], channel)
	go sortSlice(chunks[2], channel)
	go sortSlice(chunks[3], channel)

	sortedChunk1 := <-channel
	sortedChunk2 := <-channel
	sortedChunk3 := <-channel
	sortedChunk4 := <-channel

	fmt.Println("sorted subarrary #1: ", sortedChunk1)
	fmt.Println("sorted subarrary #2: ", sortedChunk2)
	fmt.Println("sorted subarrary #3: ", sortedChunk3)
	fmt.Println("sorted subarrary #4: ", sortedChunk4)

	result1 := merge(sortedChunk1, sortedChunk2)
	result2 := merge(sortedChunk3, sortedChunk4)

	result := merge(result1, result2)

	fmt.Println("sorted array       : ", result)
}

func convertStringsToInts(digitsAsString []string) []int {
	var slice []int
	for _, digit := range digitsAsString {
		number, _ := strconv.Atoi(digit)
		slice = append(slice, number)
	}
	return slice
}

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func sortSlice(items []int, channel chan []int) {
	sort.Slice(items, func(i, j int) bool {
		return items[j] > items[i]
	})
	channel <- items
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
