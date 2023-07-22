// 안전지대 
// 문제 설명
// 다음 그림과 같이 지뢰가 있는 지역과 지뢰에 인접한 위, 아래, 좌, 우 대각선 칸을 모두 위험지역으로 분류합니다.
// image.png
// 지뢰는 2차원 배열 board에 1로 표시되어 있고 board에는 지뢰가 매설 된 지역 1과, 지뢰가 없는 지역 0만 존재합니다.
// 지뢰가 매설된 지역의 지도 board가 매개변수로 주어질 때, 안전한 지역의 칸 수를 return하도록 solution 함수를 완성해주세요.

// 제한사항
// board는 n * n 배열입니다.
// 1 ≤ n ≤ 100
// 지뢰는 1로 표시되어 있습니다.
// board에는 지뢰가 있는 지역 1과 지뢰가 없는 지역 0만 존재합니다.

// 제한사항
// board는 n * n 배열입니다.
// 1 ≤ n ≤ 100
// 지뢰는 1로 표시되어 있습니다.
// board에는 지뢰가 있는 지역 1과 지뢰가 없는 지역 0만 존재합니다.
// 입출력 예
// board	result
// [[0, 0, 0, 0, 0], [0, 0, 0, 0, 0], [0, 0, 0, 0, 0], [0, 0, 1, 0, 0], [0, 0, 0, 0, 0]]	16
// [[0, 0, 0, 0, 0], [0, 0, 0, 0, 0], [0, 0, 0, 0, 0], [0, 0, 1, 1, 0], [0, 0, 0, 0, 0]]	13
// [[1, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 1], [1, 1, 1, 1, 1, 1]]	0
// 입출력 예 설명
// 입출력 예 #1

// (3, 2)에 지뢰가 있으므로 지뢰가 있는 지역과 지뢰와 인접한 위, 아래, 좌, 우, 대각선 총 8칸은 위험지역입니다. 따라서 16을 return합니다.
// 입출력 예 #2

// (3, 2), (3, 3)에 지뢰가 있으므로 지뢰가 있는 지역과 지뢰와 인접한 위, 아래, 좌, 우, 대각선은 위험지역입니다. 따라서 위험지역을 제외한 칸 수 13을 return합니다.
// 입출력 예 #3

// 모든 지역에 지뢰가 있으므로 안전지역은 없습니다. 따라서 0을 return합니다.

package main

// func solution(board [][]int) int {
// 	count := len(board[0]) * len(board)

// 	bt := make([][]bool, len(board))
// 	for i, v1 := range bt {
// 		bt[i] = make([]bool, len(board))
// 		for j := range v1 {
// 			bt[i][j] = false
// 		}
// 	}

// 	fmt.Println("bt", bt)
// 	fmt.Println(count)
// 	for i, v := range board {
// 		for j, v2 := range v {
// 			fmt.Println(i, j, v2, board[i][j])
// 			// 해당하는 배열의 주위를 1로 바꾸면 인덱스 문제가 발생한다. 그럼 true, false로 바꿔볼까??

// 			if board[i][j] == 1 && (i+1 < len(board[0]) && j+1 < len(board)) {
// 				count -= 1

// 				fmt.Println("i,j", i, j)
// 				if i-1 != -1 && j-1 != -1 {
// 					fmt.Println("test", board, count)
// 					if board[i-1][j-1] == 0 {
// 						count -= 1
// 						board[i-1][j-1] = 1
// 					}
// 					if board[i-1][j] == 0 {
// 						count -= 1
// 						board[i-1][j] = 1
// 					}
// 					if board[i-1][j+1] == 0 {
// 						count -= 1
// 						board[i-1][j+1] = 1
// 					}
// 					if board[i][j-1] == 0 {
// 						count -= 1
// 						board[i][j-1] = 1
// 					}
// 					if board[i][j+1] == 0 {
// 						count -= 1
// 						board[i][j+1] = 1
// 					}
// 					if board[i+1][j-1] == 0 {
// 						count -= 1
// 						board[i+1][j-1] = 1
// 					}
// 					if board[i+1][j] == 0 {
// 						count -= 1
// 						board[i+1][j] = 1
// 					}
// 					if board[i+1][j+1] == 0 {
// 						count -= 1
// 						board[i+1][j+1] = 1
// 					}
// 				}

// 			}

// 		}
// 	}
// 	fmt.Println(board, count)
// 	return 0
// }

// 챗GPT의 코드


func solution(board [][]int) int {
	n := len(board)
	count := 0

	// 상하좌우 및 대각선 방향
	//dx는 행, -1은 상, 0은 현재 행, 1은 다음 행
	//dy는 열  -1은 상, 0은 현재 열, 1은 다음 열
	//-1,-1 대각선 오른쪽
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				continue // 현재 칸이 지뢰인 경우 다음 칸으로 넘어감
			}

			// 주변 지뢰 개수 확인
			mineCount := 0
			for k := 0; k < 8; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if nx >= 0 && nx < n && ny >= 0 && ny < n && board[nx][ny] == 1 {
					mineCount++
				}
			}

			if mineCount == 0 {
				count++ // 안전한 지역인 경우 카운트 증가
			}
		}
	}

	return count
}

func main() {
	// board := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}}
	// board := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 1, 0}, {0, 0, 0, 0, 0}}
	board := [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}}
	solution(board)
}
