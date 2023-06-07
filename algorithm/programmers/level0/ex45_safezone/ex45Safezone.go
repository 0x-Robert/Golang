package main

import "fmt"

func solution(board [][]int) int {

	fmt.Println("board", board)

	for i, v := range board {
		//	fmt.Println(i, v)
		for j, v2 := range v {
			fmt.Println("v2", v2)
			if v2 == 1 {
				fmt.Println(i, j, v2)
				//fmt.Println(board[3][2])
				board[i-1][j-1] = 1
				board[i-1][j] = 1
				board[i-1][j+1] = 1
				board[i][j-1] = 1
				board[i][j+1] = 1
				board[i+1][j-1] = 1
				board[i+1][j] = 1
				board[i+1][j+1] = 1
			}
		}
	}
	fmt.Println("final", board)
	return 0
}

func main() {

	board := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
	solution(board)

}
