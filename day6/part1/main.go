package main

import (
	"bufio"
	"log"
	"os"
)

type Direction int

const (
	up Direction = iota
	right
	down
	left
)

func getDirectionCoords(dir Direction) (int, int) {
	if dir == up {
		return -1, 0
	}
	if dir == right {
		return 0, 1
	}
	if dir == down {
		return 1, 0
	}
	if dir == left {
		return 0, -1
	}
	panic("unexpected direction")
}

func incrementDirection(dir Direction) Direction {
	return (dir + 1) % 4
}

func findStart(grid [][]rune) (int, int) {
	for r, row := range grid {
		for c, cell := range row {
			if cell == '^' {
				return r, c
			}
		}
	}

	panic("expected to find start state, but failed")
}

func main() {
	log.SetFlags(0)

	filename := "day6/input.txt"
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
		grid = append(grid, []rune(scanner.Text()))
	}

	total := 1

	row, col := findStart(grid)
	dir := up
outer:
	for {
	inner:
		for {
			dr, dc := getDirectionCoords(dir)
			nextrow, nextcol := row+dr, col+dc
			// println("nextrow ", nextrow, " nextcol ", nextcol, " dr ", dr, " dc ", dc)

			if nextrow < 0 || nextrow >= len(grid) || nextcol < 0 || nextcol >= len(grid[0]) {
				break outer
			}

			if grid[nextrow][nextcol] == '#' {
				break inner
			}

			row, col = nextrow, nextcol

			if grid[row][col] != 'X' {
				total += 1
			}

			grid[row][col] = 'X'
		}

		dir = incrementDirection(dir)
	}

	log.Println("Total is", total)
}
