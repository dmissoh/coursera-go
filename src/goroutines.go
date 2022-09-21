/*
In the attached go code, there is a shared variable declared at line #8 (var sharedInteger int = 0)

This variable is read by two goroutines (at line#18 and line#20), increased and written (assigned back) to sharedInteger. There is a race condition since both go routines are sharing read and write access to the same variable sharedInteger.

To check that the race condition occurs, execute the code with the following command:

`go run -race goroutines.go`

*/

package main

import (
	"fmt"
	"time"
)

var sharedInteger int = 0

func f(from string) {
	for i := 0; i < 3; i++ {
		sharedInteger = sharedInteger + 1
		fmt.Println(from, ":", sharedInteger)
	}
}

func main() {
	go f("goroutine#1")

	go f("goroutine#2")

	time.Sleep(time.Second)

	fmt.Println("done")
}
