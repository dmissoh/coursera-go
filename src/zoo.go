package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow
type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Printf("grass")
}

func (c Cow) Move() {
	fmt.Printf("walk")
}

func (c Cow) Speak() {
	fmt.Printf("moo")
}

// Bird
type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Printf("worms")
}

func (b Bird) Move() {
	fmt.Printf("fly")
}

func (b Bird) Speak() {
	fmt.Printf("peep")
}

// Snake
type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Printf("mice")
}

func (s Snake) Move() {
	fmt.Printf("slither")
}

func (s Snake) Speak() {
	fmt.Printf("hsss")
}

func main() {

	mapOfAnimals := make(map[string]Animal)

	for {
		var command, param1, param2 string

		fmt.Print(">")
		fmt.Scanf("%s %s %s", &command, &param1, &param2)

		if command == "newanimal" {
			animal := createNewAnimal(param1, param2)
			mapOfAnimals[param1] = animal
			fmt.Println("Created it!")
		} else if command == "query" {
			queryAnimal(mapOfAnimals, param1, param2)
			fmt.Println("")
		}
	}
}

func createNewAnimal(animalName string, animalType string) Animal {
	if animalType == "cow" {
		return Cow{animalName}
	} else if animalType == "bird" {
		return Bird{animalName}
	}
	return Snake{animalName}
}

func queryAnimal(animals map[string]Animal, animalName string, informationType string) {
	selectedAnimal := animals[animalName]
	if informationType == "eat" {
		selectedAnimal.Eat()
	} else if informationType == "move" {
		selectedAnimal.Move()
	} else if informationType == "speak" {
		selectedAnimal.Speak()
	}
}
