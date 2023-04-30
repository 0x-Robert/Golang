package main

import "fmt"

func countIslands(grid [][]string) int {
	//dfs solution
	HEIGHT := len(grid)
	WIDTH := 0

	if HEIGHT > 0 {
		WIDTH = len(grid[0])
	}
	count := 0

	for row := 0; row < HEIGHT; row++ {
		for col := 0; col < WIDTH; col++ {
			if grid[row][col] == "0" {
				continue
			}
			count++
			searchIsland(row, col, grid, HEIGHT, WIDTH)
		}

	}
	return count
}

func searchIsland(row int, col int, grid [][]string, HEIGHT int, WIDTH int) {
	if row < 0 || col < 0 || row >= HEIGHT || col >= WIDTH {
		return
	}

	if grid[row][col] == "0" {
		return
	}

	grid[row][col] = "0"
	searchIsland(row-1, col, grid, HEIGHT, WIDTH)
	searchIsland(row+1, col, grid, HEIGHT, WIDTH)
	searchIsland(row, col-1, grid, HEIGHT, WIDTH)
	searchIsland(row, col+1, grid, HEIGHT, WIDTH)
}

func main() {
	grid := [][]string{
		{"0", "1", "1", "1"},
		{"0", "1", "1", "1"},
		{"1", "1", "0", "0"},
	}
	result := countIslands(grid)
	fmt.Println(result) // --> 1

	grid = [][]string{
		{"0", "1", "1", "1", "0"},
		{"0", "1", "0", "0", "0"},
		{"0", "0", "0", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "1", "0"},
	}
	result = countIslands(grid)
	fmt.Println(result) // --> 3
}
