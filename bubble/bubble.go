package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Swap takes a slice of integers and swaps the values at position and position + 1
func Swap(sli []int, position int) {
	sli[position], sli[position+1] = sli[position+1], sli[position]
}

// BubbleSort sorts a slice of integers into ascending order
func BubbleSort(sli []int) {
	for {
		swapped := false
		for pos := 0; pos < len(sli)-1; pos++ {
			if shouldSwap(sli, pos) {
				Swap(sli, pos)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func shouldSwap(sli []int, position int) bool {
	return sli[position] > sli[position+1]
}

func toString(sli []int) string {
	output := ""
	for _, integer := range sli {
		output = output + strconv.Itoa(integer) + " "
	}
	return strings.TrimSpace(output)
}

func main() {
	fmt.Print("Input:  ")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := strings.TrimSpace(stdin.Text())

	sli := make([]int, 0, 0)

	for _, number := range strings.Split(input, " ") {
		integer, _ := strconv.Atoi(number)
		sli = append(sli, integer)
	}

	BubbleSort(sli)
	fmt.Println("Sorted:", toString(sli))
}
