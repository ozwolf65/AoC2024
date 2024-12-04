package main

import (
	"os"
	"slices"
	"strings"
)

func p1(grid [][]rune, i, j int) int {
	matches := 0

	// Right
	if j < len(grid[i])-3 && grid[i][j] == 'X' && grid[i][j+1] == 'M' && grid[i][j+2] == 'A' && grid[i][j+3] == 'S' {
		matches += 1
	}

	// Left
	if j > 2 && grid[i][j] == 'X' && grid[i][j-1] == 'M' && grid[i][j-2] == 'A' && grid[i][j-3] == 'S' {
		matches += 1
	}

	// Down
	if i < len(grid)-3 && grid[i][j] == 'X' && grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' {
		matches += 1
	}

	// Up
	if i > 2 && grid[i][j] == 'X' && grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S' {
		matches += 1
	}

	// Down-Right
	if j < len(grid[i])-3 && i < len(grid)-3 && grid[i][j] == 'X' && grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' {
		matches += 1
	}

	// Down-Left
	if j > 2 && i < len(grid)-3 && grid[i][j] == 'X' && grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' {
		matches += 1
	}

	// Up-Right
	if i > 2 && j < len(grid[i])-3 && grid[i][j] == 'X' && grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S' {
		matches += 1
	}

	// Up-Left
	if i > 2 && j > 2 && grid[i][j] == 'X' && grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S' {
		matches += 1
	}
	return matches

}

func p2(grid [][]rune, i, j int) int {
	if i < 1 || j < 1 || i >= len(grid)-1 || j >= len(grid[i])-1 {
		return 0
	}
	centre := grid[i][j]
	top_left := grid[i-1][j-1]
	top_right := grid[i-1][j+1]
	bottom_left := grid[i+1][j-1]
	bottom_right := grid[i+1][j+1]
	good := []rune{'M', 'S'}
	if centre != 'A' {
		return 0
	}
	// if not all M and S
	if !slices.Contains(good, top_right) || !slices.Contains(good, top_left) || !slices.Contains(good, bottom_right) || !slices.Contains(good, bottom_left) {
		return 0
	}
	// If not spelling out MAS
	if top_left == bottom_right || top_right == bottom_left {
		return 0
	}

	return 1
}

func main() {
	f, err := os.ReadFile("./day4/input.txt")
	if err != nil {
		panic(err)
	}
	grid := make([][]rune, 0)
	rows := strings.Split(string(f), "\n")
	for _, row := range rows {
		grid = append(grid, []rune(row))
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count += p1(grid, i, j)
		}
	}
	println(count)
	count = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			count += p2(grid, i, j)
		}
	}
	println(count)

}
