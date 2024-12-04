package main

import (
	"bufio"
	"log"
	"os"
)

func search(lines []string, searchstring *string, r int, c int, i int, dr int, dc int) int {
	if i == len(*searchstring) {
		return 1
	}

	if r < 0 || r >= len(lines) || c < 0 || c >= len(lines[0]) {
		return 0
	}

	if lines[r][c] != (*searchstring)[i] {
		return 0
	}

	return search(lines, searchstring, r+dr, c+dc, i+1, dr, dc)
}

func main() {
	log.SetFlags(0)

	filename := "day4/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	searchstring := "XMAS"

	total := 0

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[0]); c++ {
			total += search(lines, &searchstring, r, c, 0, -1, -1)
			total += search(lines, &searchstring, r, c, 0, -1, 0)
			total += search(lines, &searchstring, r, c, 0, -1, 1)
			total += search(lines, &searchstring, r, c, 0, 0, -1)
			total += search(lines, &searchstring, r, c, 0, 0, 1)
			total += search(lines, &searchstring, r, c, 0, 1, -1)
			total += search(lines, &searchstring, r, c, 0, 1, 0)
			total += search(lines, &searchstring, r, c, 0, 1, 1)
		}
	}

	log.Println("Total is", total)
}
