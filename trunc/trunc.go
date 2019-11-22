package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Enter a floating point number: ")
	fmt.Scanln(&input)

	floatToTrunc, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Error: Invalid floating point number")
	} else {
		var output int = int(floatToTrunc)
		fmt.Println("Truncated integer:", output)
	}
}
