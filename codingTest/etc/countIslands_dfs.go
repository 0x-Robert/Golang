package main

import (
	"fmt"
)

func countIslands(grid [][]string) int {
	// 방문 여부를 체크하기 위한 2차원 배열 생성
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := 0

	// 상하좌우 이동을 위한 배열
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	// DFS를 위한 함수 정의
	var dfs func(int, int)
	dfs = func(x, y int) {
		visited[x][y] = true

		// 상하좌우 인접 노드 방문
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			// 범위를 벗어나거나, 물이라면 방문하지 않음
			if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] == "0" {
				continue
			}

			// 방문하지 않은 인접 노드 방문
			if !visited[nx][ny] {
				dfs(nx, ny)
			}
		}
	}

	// 모든 노드에 대해서 DFS 실행
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == "1" {
				count++
				dfs(i, j)
			}
		}
	}

	return count
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
