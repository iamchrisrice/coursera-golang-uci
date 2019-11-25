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

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hiss")
}

func getCommandsFromInput(input string) (string, string, string) {
	w := strings.Split(input, " ")
	return strings.ToLower(w[0]), strings.ToLower(w[1]), strings.ToLower(w[2])
}

func createAnimal(name string) Animal {
	switch name {
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

		cmd1, cmd2, cmd3 := getCommandsFromInput(input)

		switch cmd1 {
		case "newanimal":
			animals[cmd2] = createAnimal(cmd3)
		case "query":
			query(animals[cmd2], cmd3)
		}
	}
}
