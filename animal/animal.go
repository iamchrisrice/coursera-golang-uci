package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func getAnimalNameAndCommandFromInput(input string) (string, string) {
	w := strings.Split(input, " ")
	animal := strings.ToLower(w[0])
	command := strings.ToLower(w[1])
	return animal, command
}

func animalCommand(animal Animal, command string) {
	switch command {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func createAnimals() *map[string]Animal {
	animals := make(map[string]Animal)
	animals["cow"] = Animal{"grass", "walk", "moo"}
	animals["bird"] = Animal{"worms", "fly", "peep"}
	animals["snake"] = Animal{"mice", "slither", "hsss"}
	return &animals
}

func main() {
	animals := *createAnimals()

	for {
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input := stdin.Text()

		animalName, command := getAnimalNameAndCommandFromInput(input)
		if animal, exists := animals[animalName]; exists {
			animalCommand(animal, command)
		}
	}
}
