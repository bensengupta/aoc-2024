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

func dfs(grid [][]int, visit map[Pos]bool, r int, c int, n int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}
	if grid[r][c] != n {
		return 0
	}
	if n == 9 {
		if visit[Pos{r, c}] {
			return 0
		}
		visit[Pos{r, c}] = true
		return 1
	}
	return dfs(grid, visit, r-1, c, n+1) + dfs(grid, visit, r+1, c, n+1) + dfs(grid, visit, r, c-1, n+1) + dfs(grid, visit, r, c+1, n+1)
}

func main() {
	log.SetFlags(0)

	filename := "day10/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, c := range line {
			num := int(c - '0')
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	total := 0
	for r, gridRow := range grid {
		for c, cell := range gridRow {
			if cell == 0 {
				score := dfs(grid, make(map[Pos]bool), r, c, 0)
				total += score
			}
		}
	}

	log.Println("Total is", total)
}
