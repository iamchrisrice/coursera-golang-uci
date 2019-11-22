package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startsWithI(input string) bool {
	return strings.HasPrefix(input, "i")
}

func endsWithN(input string) bool {
	return strings.HasSuffix(input, "n")
}

func containsA(input string) bool {
	return strings.Contains(input, "a")
}

func main() {
	fmt.Print("Input: ")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := strings.ToLower(stdin.Text())

	if startsWithI(input) && containsA(input) && endsWithN(input) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
