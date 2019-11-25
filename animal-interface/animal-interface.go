package main

import (
	"bufio"
	"errors"
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

func parseInput(input string) (string, string, string, error) {
	var c0, c1, c2 string
	err := errors.New("invalid input")
	w := strings.Split(input, " ")
	if len(w) == 3 {
		c0 = strings.ToLower(w[0])
		c1 = strings.ToLower(w[1])
		c2 = strings.ToLower(w[2])
		err = nil
	}
	return c0, c1, c2, err
}

func animalFromSpecies(species string) (Animal, error) {
	switch species {
	case "cow":
		return Cow{}, nil
	case "bird":
		return Bird{}, nil
	case "snake":
		return Snake{}, nil
	default:
		return nil, errors.New("unknown species " + species)
	}
}

func query(animal Animal, verb string) error {
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
		return errors.New("unkonwn verb: " + verb)
	}
}

func main() {
	animals := make(map[string]Animal)

	for {
		var err error = nil

		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		input := stdin.Text()

		c0, c1, c2, err := parseInput(input)

		if err == nil {
			switch c0 {
			case "newanimal":
				animals[c1], err = animalFromSpecies(c2)
			case "query":
				if animal, exists := animals[c1]; exists {
					err = query(animal, c2)
				} else {
					err = errors.New("unknown animal: " + c1)
				}
			default:
				err = errors.New("invalid input")
			}
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
