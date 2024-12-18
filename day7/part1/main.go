package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func hasCombination(sum int64, cur int64, nums []int64, i int) bool {
	if cur == sum {
		return true
	}
	if cur > sum || i == len(nums) {
		return false
	}
	return hasCombination(sum, cur+nums[i], nums, i+1) || hasCombination(sum, cur*nums[i], nums, i+1)
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
