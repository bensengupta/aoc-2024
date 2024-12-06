package main

import (
	"bufio"
	"log"
	"os"
)

type Direction int

const (
	up    Direction = 1
	right Direction = 2
	down  Direction = 4
	left  Direction = 8
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
	// 0001
	// 0010
	// 0100
	// 1000
	// 0001
	// ...
	return (dir << 1) % 15
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

func simulate(grid [][]rune, row int, col int, dir Direction) int {
	visit := make([][]int, len(grid))
	for i := range visit {
		visit[i] = make([]int, len(grid[0]))
	}

	for {
	inner:
		for {
			dr, dc := getDirectionCoords(dir)
			nextrow, nextcol := row+dr, col+dc

			if nextrow < 0 || nextrow >= len(grid) || nextcol < 0 || nextcol >= len(grid[0]) {
				return 0
			}

			if visit[nextrow][nextcol]&int(dir) != 0 {
				// printGrid(grid, row, col, dir, obsRow, obsCol, visit)
				// loop
				return 1
			}

			if grid[nextrow][nextcol] == '#' {
				break inner
			}

			row, col = nextrow, nextcol
			visit[row][col] = visit[row][col] | int(dir)
		}

		dir = incrementDirection(dir)
		visit[row][col] = visit[row][col] | int(dir)
	}
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

	total := 0

	startrow, startcol := findStart(grid)
	startdir := up
	grid[startrow][startcol] = '.'

	for obsRow := range grid {
		for obsCol := range grid[0] {
			if grid[obsRow][obsCol] == '.' {
				grid[obsRow][obsCol] = '#'
				total += simulate(grid, startrow, startcol, startdir)
				grid[obsRow][obsCol] = '.'
			}
		}
	}

	log.Println("Total is", total)
}
