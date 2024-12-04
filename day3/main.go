package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	//part 1
	input1, err := os.ReadFile("day3/input1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(MullItOver(string(input1)))

	//part 2
	input2, err := os.ReadFile("day3/input2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(DoAndDontParser(string(input2)))
}

func MullItOver(data string) int {
	mullPattern := `mul\(\d{1,3},\d{1,3}\)`
	numPattern := `\d{1,3}`

	mullRegex, err := regexp.Compile(mullPattern)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}

	numRegex, err := regexp.Compile(numPattern)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}

	mullMatch := mullRegex.FindAllString(data, -1)

	var sum int
	for _, item := range mullMatch {
		nums := numRegex.FindAllString(item, 2)

		x, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		y, _ := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("Failed to convert string to int: %v", err)
		}

		mul := x * y
		sum = sum + mul
	}
	return sum
}

func DoAndDontParser(data string) int {
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`
	doRegex, err := regexp.Compile(doPattern)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}
	dontRegex, err := regexp.Compile(dontPattern)
	if err != nil {
		log.Fatalf("Failed to compile regex: %v", err)
	}

	var sum int

	for {
		// Keep counting until the first don't() is found
		matchDont := dontRegex.FindStringIndex(data)
		if matchDont == nil {
			break
		}

		before := data[:matchDont[0]]
		sum = sum + MullItOver(before)
		// parse data after the don't()
		newData := data[matchDont[1]:]
		// Start counting only after a do() in encountered
		matchDo := doRegex.FindStringIndex(newData)
		if matchDo == nil {
			break
		}
		data = newData[matchDo[1]:]

	}

	return sum

}
