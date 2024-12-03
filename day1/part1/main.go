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
	var arr2 []int64

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
		arr2 = append(arr2, num2)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	var total int64
	for i := 0; i < len(arr1); i++ {
		diff := arr1[i] - arr2[i]
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}

	log.Println("Total is", total)
}

