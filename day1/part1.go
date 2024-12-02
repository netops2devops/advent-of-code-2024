package main

import (
	"os"
	"strconv"
	"strings"
)

func readInput() []byte {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return bytes
}

func cleanData() ([]int, []int) {
	var evenSlice, oddSlice []int
	bytes := readInput()
	data := string(bytes)
	newSlice := strings.Fields(data)

	for i, item := range newSlice {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			evenSlice = append(evenSlice, num)
		} else {
			oddSlice = append(oddSlice, num)
		}
	}
	return evenSlice, oddSlice
}
