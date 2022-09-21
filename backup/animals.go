package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() string {
	return animal.food
}

func (animal Animal) Move() string {
	return animal.locomotion
}

func (animal Animal) Speak() string {
	return animal.noise
}

func main() {

	// init
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	mapOfAnimals := make(map[string]Animal)
	mapOfAnimals["cow"] = cow
	mapOfAnimals["bird"] = bird
	mapOfAnimals["snake"] = snake

	for {
		// acceleration, initial velocity, and initial displacement.
		var animal, verb string

		fmt.Print(">")
		fmt.Scanf("%s %s", &animal, &verb)

		selectedAnimal := mapOfAnimals[animal]

		result := "not available"

		if verb == "eat" {
			result = selectedAnimal.Eat()
		} else if verb == "move" {
			result = selectedAnimal.Move()
		} else if verb == "speak" {
			result = selectedAnimal.Speak()
		}

		fmt.Printf("%s(%s) = %s \n", verb, animal, result)
	}

}
