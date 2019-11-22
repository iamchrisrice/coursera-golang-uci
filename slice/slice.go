package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)

	for {
		var input string

		fmt.Print("Enter an Integer (or X to eXit): ")
		fmt.Scanln(&input)

		if input == "X" {
			break
		}

		number, _ := strconv.Atoi(input)
		sli = append(sli, number)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
