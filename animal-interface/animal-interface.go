package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hiss") }

func getCommandsFromInput(input string) (string, string, string) {
	var command1, command2, command3 string
	w := strings.Split(input, " ")
	if len(w) == 3 {
		command1 = strings.ToLower(w[0])
		command2 = strings.ToLower(w[1])
		command3 = strings.ToLower(w[2])
	}
	return command1, command2, command3
}

func animalFromSpecies(species string) Animal {
	switch species {
	case "cow":
		return Cow{}
	case "bird":
		return Bird{}
	case "snake":
		return Snake{}
	default:
		return nil
	}
}

func query(animal Animal, category string) {
	switch category {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input := stdin.Text()

		command1, command2, command3 := getCommandsFromInput(input)

		switch command1 {
		case "newanimal":
			animals[command2] = animalFromSpecies(command3)
		case "query":
			if animal, exists := animals[command2]; exists {
				query(animal, command3)
			}
		}
	}
}
