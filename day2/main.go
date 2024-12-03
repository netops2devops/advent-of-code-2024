package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int // counter for tracking safe records

func main() {
	safeReportCount := safetyCheck("day2/input.txt")
	fmt.Println(safeReportCount)

}

func safetyCheck(inputFile string) int {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	s := string(data)

	lines := strings.Split(s, "\n")
	for _, record := range lines {
		if safeAscendingRecord(record) {
			count++
		} else if safeDescendingRecord(record) {
			count++
		} else {
			singleViolationAllowed(record)
		}

	}
	return count
}

func safeAscendingRecord(s string) bool {
	record := strings.Fields(s)
	for i := 0; i < len(record)-1; i++ {
		// when all numbers in one record are in increasing order
		num, _ := strconv.Atoi(record[i])
		adjNum, _ := strconv.Atoi(record[i+1])
		if adjNum > num && adjNum-num >= 1 && adjNum-num <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}

func safeDescendingRecord(s string) bool {
	record := strings.Fields(s)
	for i := 0; i < len(record)-1; i++ {
		// when all numbers in one record are in decreasing order
		num, _ := strconv.Atoi(record[i])
		adjNum, _ := strconv.Atoi(record[i+1])
		if num > adjNum && num-adjNum >= 1 && num-adjNum <= 3 {
			continue
		} else {
			return false
		}
	}
	return true
}

func singleViolationAllowed(s string) {
	record := strings.Fields(s)
	for i := 0; i < len(record); i++ {
		freshRecord := strings.Fields(s)
		newSlice := append(freshRecord[:i], freshRecord[i+1:]...)
		if safeAscendingRecord(strings.Join(newSlice, " ")) || safeDescendingRecord(strings.Join(newSlice, " ")) {
			count++
			break
		}
	}
}
