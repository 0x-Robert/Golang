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
	//오른쪽으로 이동, 아래로 이동, 왼쪽으로 이동, 위로 이동
	//하단의 if 문 안에 dirIdx 변수를 통해 방향을 확인할 수 있다.
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
		//경계를 넘을 때 dirIdx가 업데이트 된다.
		//처음에는 오른쪽으로 이동하다가 경계를 넘으면 아래로 이동하게끔 업데이트 된다. 따라서 dirIdx는0에서 1로 변경된다.
		if nextRow < 0 || nextRow >= n || nextCol < 0 || nextCol >= n || res[nextRow][nextCol] != 0 {
			dirIdx = (dirIdx + 1) % 4
			//0 + 1 % 4 = 1 오른쪽 이동에서 아래로 이동
			//1 + 1 % 4 = 2 아래 이동에서 왼쪽으로 이동
			//2 + 1 % 4 = 3 왼쪽 이동에서 위쪽으로 이동
			//3 + 1 % 4 = 0 위로 이동에서 오른쪽으로 이동
			fmt.Println("check dirIdx", dirIdx)
		}

		//다음 위치로 이동
		fmt.Println("row1", row)
		fmt.Println("col1", col)
		row += dirs[dirIdx][0]
		col += dirs[dirIdx][1]
		fmt.Println("row2", row)
		fmt.Println("col2", col)
		//채널 숫자 증가
		num++
	}

	return res
}

type Field struct {
	X int
	Y int
}

func solution2(n int) [][]int {
	answer := make([][]int, n)
	for i := 0; i < n; i++ {
		answer[i] = make([]int, n)
	}

	field := &Field{-1, 0}
	way := 0
	for i := 1; i <= n*n; i++ {
		switch way {
		case 0:
			if field.X+1 < n && answer[field.Y][field.X+1] == 0 {
				field.X++
			} else {
				way = 1
				i--
				continue
			}
		case 1:
			if field.Y+1 < n && answer[field.Y+1][field.X] == 0 {
				field.Y++
			} else {
				way = 2
				i--
				continue
			}
		case 2:
			if field.X-1 >= 0 && answer[field.Y][field.X-1] == 0 {
				field.X--
			} else {
				way = 3
				i--
				continue
			}
		case 3:
			if field.Y-1 >= 0 && answer[field.Y-1][field.X] == 0 {
				field.Y--
			} else {
				way = 0
				i--
				continue
			}
		}
		answer[field.Y][field.X] = i
	}
	return answer
}

func main() {

	n := 4
	solution(n)
}
