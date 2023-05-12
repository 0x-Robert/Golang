package main

import "fmt"

func solution(n int) [][]int {

	//결과로 쓸 배열을 정의한다.
	//이중 배열 선언
	res := make([][]int, n)

	// 	res [[] [] [] []]
	// [0 0 0 0]
	// [0 0 0 0]
	// [0 0 0 0]
	// [0 0 0 0]
	//이중 배열 크기만큼 순회해서 res 배열에 원소 설정
	fmt.Println("res", res)
	for i := range res {
		res[i] = make([]int, n)
		fmt.Println(res[i])
	}

	//움질일 방향을 설정하기 위해 다음 변수를 설정하자.
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	//현재 위치 및 방향 초기화
	//행, 열, 나선방향, 위,오른쪽,왼쪽,아래
	row, col, dirIdx := 0, 0, 0

	//배열에 채울 숫자 초기화
	num := 1

	//배열을 모두 채울때까지 반복
	for num <= n*n {
		//현재 위치에 숫자 채우기
		//1로 채우기
		res[row][col] = num
		fmt.Println("res", res)
		nextRow, nextCol := row+dirs[dirIdx][0], col+dirs[dirIdx][1]
		fmt.Println("nextRow", nextRow, "nextCol", nextCol)
		//다음 위치가 배열의 범위를 벗어나거나 이미 숫자가 채워져 있는 경우 방향을 바꿈
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || res[nextRow][nextCol] != 0 {
			dirIdx = (dirIdx + 1) % 4
			//0 + 1 % 4 = 1
			//1 + 1 % 4 = 2
			//2 + 1 % 4 = 3
			//3 + 1 % 4 = 0
		}

		//다음 위치로 이동
		fmt.Println("row", row)
		row += dirs[dirIdx][0]
		col += dirs[dirIdx][1]
		fmt.Println("row", row)
		//채널 숫자 증가
		num++
	}

	return res
}

func main() {

	n := 4
	solution(n)
}
