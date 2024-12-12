package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func blink(nums *[]uint64) {
	var newNums []uint64

	for _, num := range *nums {
		if num == 0 {
			newNums = append(newNums, 1)
			continue
		}

		numDigits := int(math.Log10(float64(num))) + 1

		if numDigits%2 == 0 {
			pow := uint64(math.Pow10(numDigits / 2))
			left := num / pow
			right := num % pow
			newNums = append(newNums, left)
			newNums = append(newNums, right)
			continue
		}

		newNums = append(newNums, num*2024)
	}

	*nums = newNums
}

func main() {
	log.SetFlags(0)

	filename := "day11/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		log.Fatal("failed to read line")
	}

	var nums []uint64
	for _, str := range strings.Split(scanner.Text(), " ") {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, uint64(num))
	}

	numBlinks := 25
	for i := 0; i < numBlinks; i++ {
		println("blink", i, "/", numBlinks)
		blink(&nums)
	}

	// for _, n := range nums {
	// 	print(n, " ")
	// }
	// print("\n")

	println("Total is", len(nums))
}
