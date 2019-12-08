package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func stdinToIntegerSlice() []int {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	input := strings.TrimSpace(stdin.Text())
	var sli []int
	for _, number := range strings.Split(input, " ") {
		integer, _ := strconv.Atoi(number)
		sli = append(sli, integer)
	}
	return sli
}

func chunk(sli []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(sli); i += size {
		if len(sli) > i+size {
			chunks = append(chunks, sli[i:i+size])
		} else {
			chunks = append(chunks, sli[i:])
		}
	}
	return chunks
}

func numChunks(sli []int, num int) [][]int {
	chunkSize := int(math.Ceil(float64(len(sli)) / float64(num)))
	chunks := chunk(sli, chunkSize)
	for i := len(chunks); i < num; i++ {
		var empty []int
		chunks = append(chunks, empty)
	}
	return chunks
}

func merge(first []int, second []int) []int {
	sli, sliPos := make([]int, len(first)+len(second)), 0
	for len(first) > 0 && len(second) > 0 {
		if first[0] < second[0] {
			sli[sliPos], first = first[0], first[1:]
		} else {
			sli[sliPos], second = second[0], second[1:]
		}
		sliPos++
	}
	for firstPos := 0; firstPos < len(first); firstPos++ {
		sli[sliPos] = first[firstPos]
		sliPos++
	}
	for secondPos := 0; secondPos < len(second); secondPos++ {
		sli[sliPos] = second[secondPos]
		sliPos++
	}
	return sli
}

func mergeChunksIntoSlice(chunks [][]int) []int {
	var sli []int
	for chunkNumber := 0; chunkNumber < len(chunks); chunkNumber++ {
		sli = merge(sli, chunks[chunkNumber])
	}
	return sli
}

func toString(sli []int) string {
	output := ""
	for _, integer := range sli {
		output = output + strconv.Itoa(integer) + " "
	}
	return strings.TrimSpace(output)
}

func main() {
	const numberOfChunks = 4
	fmt.Print("Input:  ")
	sli := stdinToIntegerSlice()
	chunks := numChunks(sli, numberOfChunks)

	var wg sync.WaitGroup
	wg.Add(numberOfChunks)
	for chunkNumber := 0; chunkNumber < numberOfChunks; chunkNumber++ {

		go func(chunk int) {
			defer wg.Done()
			fmt.Println("Sorting:", chunks[chunk])
			sort.Ints(chunks[chunk])
		}(chunkNumber)

	}
	wg.Wait()

	sorted := mergeChunksIntoSlice(chunks)
	fmt.Println("Sorted:", toString(sorted))
}
