package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
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

func getAnimalNameAndVerbFromInput(input string) (string, string, error) {
	var animal, verb string
	err := errors.New("invalid input")
	w := strings.Split(input, " ")
	if len(w) == 2 {
		animal = strings.ToLower(w[0])
		verb = strings.ToLower(w[1])
		err = nil
	}
	return animal, verb, err
}

func animalVerb(animal *Animal, verb string) error {
	switch verb {
	case "eat":
		animal.Eat()
		return nil
	case "move":
		animal.Move()
		return nil
	case "speak":
		animal.Speak()
		return nil
	default:
		return errors.New("unkown verb: " + verb)
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
		var err error = nil
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input := stdin.Text()

		animalName, command, err := getAnimalNameAndVerbFromInput(input)
		if err == nil {
			if animal, exists := animals[animalName]; exists {
				err = animalVerb(&animal, command)
			} else {
				err = errors.New("unknown species: " + animalName)
			}
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
