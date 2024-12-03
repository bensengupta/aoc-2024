package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)

	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr1 []int64

	counts := make(map[int64]int64)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		if len(parts) != 2 {
			log.Fatal("expected two parts, got ", len(parts))
		}

		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal("invalid integer: ", num1)
		}
		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal("invalid integer: ", num2)
		}

		arr1 = append(arr1, num1)
		counts[num2] += 1
	}

	slices.Sort(arr1)

	var total int64
	for _, n := range arr1 {
		total += n * counts[n]
	}

	log.Println("Total is", total)
}

