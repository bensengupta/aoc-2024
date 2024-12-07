package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func concat(num1 int64, num2 int64) int64 {
	exp := int(math.Log10(float64(num2))) + 1
	total := num1*int64(math.Pow10(exp)) + num2
	// println("concat ", num1, " ", num2, " -> ", total)
	return total
}

func hasCombination(sum int64, cur int64, nums []int64, i int) bool {
	if cur > sum {
		return false
	}
	if i == len(nums) {
		return cur == sum
	}
	return (hasCombination(sum, cur+nums[i], nums, i+1) ||
		hasCombination(sum, cur*nums[i], nums, i+1) ||
		hasCombination(sum, concat(cur, nums[i]), nums, i+1))
}

func main() {
	log.SetFlags(0)

	filename := "day7/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int64

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			log.Fatal("unexpected length of parts")
		}
		sum, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		var nums []int64
		for _, str := range strings.Split(parts[1], " ") {
			num, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		if hasCombination(sum, nums[0], nums, 1) {
			total += sum
		}
	}

	log.Println("Total is", total)
}
