package main

import (
	"fmt"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	evenSlice, oddSlice := cleanData()

	sort.Ints(evenSlice)
	sort.Ints(oddSlice)

	sum := 0
	for i := 0; i < len(evenSlice); i++ {
		diff := evenSlice[i] - oddSlice[i]
		sum = sum + abs(diff)
	}

	fmt.Println(sum)               // part 1
	fmt.Println(SimilarityScore()) // part 2
}
