package main

import (
	"fmt"
	"sync"
)

func main() {

	result := make([]string, 0, 3)

	/**
	 Two goroutines appending to same slice.
	 Order of slice at end of program execution, and therefore printed output, depends on order of goroutine execution.
	**/
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		result = append(result, "Max Verstappen")
	}()
	go func() {
		defer wg.Done()
		result = append(result, "Lewis Hamilton")
	}()
	wg.Wait()

	fmt.Println("The winner is: " + result[0])

}
