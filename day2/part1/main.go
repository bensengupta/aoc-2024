package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numSafeReports := 0

outer:
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		first, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		second, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		isAscending := first < second
		last := first

		for _, str := range parts[1:] {
			next, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			if isAscending {
				// check 1: numbers are strictly ascending
				if last >= next {
					continue outer
				}
				// check 2: differ by at most 3
				if next-last > 3 {
					continue outer
				}
			} else {
				// check 1: numbers are strictly descending
				if next >= last {
					continue outer
				}
				// check 2: differ by at most 3
				if last-next > 3 {
					continue outer
				}
			}

			last = next
		}

		numSafeReports += 1
	}
	log.Println("Total number of safe reports:", numSafeReports)
}
