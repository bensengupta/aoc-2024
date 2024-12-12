package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	num  uint64
	iter int
}

func dfs(num uint64, i int, memo map[Stone]int) int {
	key := Stone{num, i}
	if memo[key] != 0 {
		return memo[key]
	}
	if i == 0 {
		return 1
	}
	if num == 0 {
		res := dfs(1, i-1, memo)
		memo[key] = res
		return res
	}

	numDigits := int(math.Log10(float64(num))) + 1

	if numDigits%2 == 0 {
		pow := uint64(math.Pow10(numDigits / 2))
		left := dfs(num/pow, i-1, memo)
		right := dfs(num%pow, i-1, memo)
		res := left + right
		memo[key] = res
		return res
	}

	res := dfs(num*2024, i-1, memo)
	memo[key] = res
	return res
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

	total := 0
	numBlinks := 75
	memo := make(map[Stone]int)
	for i, num := range nums {
		println("stone", i+1, "/", len(nums))
		res := dfs(num, numBlinks, memo)
		total += res
	}

	println("Total is", total)
}
