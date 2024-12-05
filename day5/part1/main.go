package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidUpdate(update []int, before map[int]map[int]bool) bool {
	for i := 0; i < len(update); i++ {
		for j := 0; j < i; j++ {
			if before[update[i]][update[j]] {
				return false
			}
		}
	}
	return true
}

func main() {
	log.SetFlags(0)

	filename := "day5/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// rule in the format '75|47' means 75 should come before 47
	// whch is translated to
	// before[75][47] = true
	before := make(map[int]map[int]bool)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			log.Fatal("expected 2 parts, got ", len(parts))
		}
		num1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if before[int(num1)] == nil {
			before[int(num1)] = make(map[int]bool)
		}
		before[int(num1)][int(num2)] = true
	}

	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var nums []int

		parts := strings.Split(line, ",")
		for _, str := range parts {
			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, int(num))
		}

		updates = append(updates, nums)
	}

	total := 0

	for _, update := range updates {
		if isValidUpdate(update, before) {
			total += update[len(update)/2]
		}
	}

	log.Println("Total is", total)
}
