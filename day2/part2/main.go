package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func failsConditions(isAscending bool, num1 int64, num2 int64) bool {
	if !isAscending {
		return failsConditions(true, num2, num1)
	}
	if num1 >= num2 {
		return true
	}
	if num2-num1 > 3 {
		return true
	}
	return false
}

func isValidReport(levels []int64) bool {
	first := levels[0]
	second := levels[1]

	isAscending := first < second
	last := first

	for _, next := range levels[1:] {
		if failsConditions(isAscending, last, next) {
			return false
		}

		last = next
	}
	return true
}

func printSlice(levels []int64) {
	for _, n := range levels {
		print(n, " ")
	}
	print("\n")
}

func isValidReportAllowOneProblem(levels []int64) bool {
	if isValidReport(levels) {
		return true
	}

	for i := range levels {
		var levelsCopy []int64
		levelsCopy = append(levelsCopy, levels[:i]...)
		levelsCopy = append(levelsCopy, levels[i+1:]...)
		if isValidReport(levelsCopy) {
			return true
		}
	}

	return false
}

func main() {
	log.SetFlags(0)

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numSafeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var levels []int64
		for _, str := range parts {
			level, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, level)
		}

		if isValidReportAllowOneProblem(levels) {
			numSafeReports += 1
		}
	}

	log.Println("Total number of safe reports:", numSafeReports)
}

