package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)

	bytes, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := string(bytes)

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	var total int64

	for _, match := range r.FindAllString(contents, -1) {
		parts := strings.Split(match[4:len(match)-1], ",")
		if len(parts) != 2 {
			log.Fatal("unexpected multiplication parts, expected 2, got ", len(parts))
		}
		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		total += num1 * num2
	}

	log.Println("Total is", total)
}
