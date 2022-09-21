package main

import (
	"fmt"
	"math"
)

func main() {
	// acceleration, initial velocity, and initial displacement.
	var acceleration, velocity, displacement, time float64

	fmt.Print("Enter acceleration: ")
	fmt.Scanf("%v", &acceleration)

	fmt.Print("Enter the initial velocity: ")
	fmt.Scanf("%v", &velocity)

	fmt.Print("Enter the initial displacement: ")
	fmt.Scanf("%v", &displacement)

	fmt.Print("Enter the time: ")
	fmt.Scanf("%v", &time)

	fn := GenDisplaceFn(acceleration, velocity, displacement)

	result := fn(time)

	fmt.Println("The displacement is:", result)
}

// acceleration a, initial velocity vo, and initial displacement so
func GenDisplaceFn(a, vo, so float64) func(time float64) float64 {
	return func(time float64) float64 {
		//s = ½ a t^2 + vo*t + so
		return 0.5*a*(math.Pow(time, 2)) + (vo * time) + so
	}
}
