package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname, lname string
}

func stringToName(s string) name {
	w := strings.Split(s, " ")
	n := name{fname: w[0], lname: w[1]}
	return n
}

func nameToString(n name) string {
	return "First Name: " + n.fname + ", Last Name: " + n.lname
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("Filename: ")
	stdin.Scan()

	file, err := os.Open(stdin.Text())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	names := make([]name, 0, 1)

	for scanner.Scan() {
		names = append(names, stringToName(scanner.Text()))
	}

	for _, n := range names {
		fmt.Println(nameToString(n))
	}
}
