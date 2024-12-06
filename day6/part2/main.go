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

func printGrid(grid [][]rune, row int, col int, dir Direction, obsRow int, obsCol int, visit [][]int) {
	guard := "^"
	if dir == right {
		guard = ">"
	} else if dir == down {
		guard = "v"
	} else if dir == left {
		guard = "<"
	}

	for r, gridRow := range grid {
		for c, cell := range gridRow {
			if r == row && c == col {
				print(guard)
				continue
			}
			if r == obsRow && c == obsCol {
				print("O")
				continue
			}
			lr := (visit[r][c]&int(left))|(visit[r][c]&int(right)) != 0
			ud := (visit[r][c]&int(up))|(visit[r][c]&int(down)) != 0
			if lr && ud {
				print("+")
				continue
			}
			if lr {
				print("-")
				continue
			}
			if ud {
				print("|")
				continue
			}
			print(string(cell))
		}
		print("\n")
	}
	print("\n\n")
}

func simulateObstacle(grid [][]rune, row int, col int, dir Direction, obsRow int, obsCol int) int {
	// early return: already determined that this obstacle pos can result in a loop
	if grid[obsRow][obsCol] == 'x' {
		return 0
	}

	visit := make([][]int, len(grid))
	for i := range visit {
		visit[i] = make([]int, len(grid[0]))
	}

	if grid[obsRow][obsCol] != '.' {
		log.Fatal("expected obstacle position to be '.' but got:", grid[obsRow][obsCol])
	}

	grid[obsRow][obsCol] = '#'

	for {
	inner:
		for {
			dr, dc := getDirectionCoords(dir)
			nextrow, nextcol := row+dr, col+dc

			if nextrow < 0 || nextrow >= len(grid) || nextcol < 0 || nextcol >= len(grid[0]) {
				grid[obsRow][obsCol] = '.'
				return 0
			}

			if visit[nextrow][nextcol]&int(dir) != 0 {
				// printGrid(grid, row, col, dir, obsRow, obsCol, visit)
				// mark the loop and return
				grid[obsRow][obsCol] = 'x'
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

	row, col := findStart(grid)
	grid[row][col] = '.'
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

			total += simulateObstacle(grid, row, col, dir, nextrow, nextcol)

			row, col = nextrow, nextcol
		}

		dir = incrementDirection(dir)
	}

	log.Println("Total is", total)
}
