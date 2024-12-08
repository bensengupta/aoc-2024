package main

import (
	"bufio"
	"log"
	"os"
)

type Pos struct {
	r int
	c int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func markAntinodes(grid [][]rune, pos1 Pos, pos2 Pos) {
	k := 0
	for {
		r := pos1.r + k*(pos1.r-pos2.r)
		c := pos1.c + k*(pos1.c-pos2.c)

		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			break
		}
		grid[r][c] = '#'
		k++
	}

	k = 0
	for {
		r := pos2.r - k*(pos1.r-pos2.r)
		c := pos2.c - k*(pos1.c-pos2.c)

		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			break
		}
		grid[r][c] = '#'
		k++
	}
}

func main() {
	log.SetFlags(0)

	filename := "day8/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	antennas := make(map[rune][]Pos)

	for r, gridRow := range grid {
		for c, cell := range gridRow {
			if cell != '.' {
				antennas[cell] = append(antennas[cell], Pos{r, c})
			}
		}
	}

	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				ant1 := positions[i]
				ant2 := positions[j]
				markAntinodes(grid, ant1, ant2)
			}
		}
	}

	total := 0
	for _, gridRow := range grid {
		for _, cell := range gridRow {
			if cell == '#' {
				total++
			}
			// print(string(cell))
		}
		// print("\n")
	}

	log.Println("Total is", total)
}

