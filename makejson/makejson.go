package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := make(map[string]string)
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Print("Name: ")
	stdin.Scan()
	person["name"] = stdin.Text()

	fmt.Print("Address: ")
	stdin.Scan()
	person["address"] = stdin.Text()

	personJSON, _ := json.Marshal(person)

	fmt.Println(string(personJSON))
}
