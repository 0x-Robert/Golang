package main

// func solution(arr1 [][]int, arr2 [][]int) [][]int {
// 	rows := len(arr1)    //배열 길이로 행 설정
// 	cols := len(arr1[0]) //첫 번째 행렬의 열길이를 설정

// 	answer := make([][]int, rows) //이중 배열을 배열 길이만큼 생성

// 	for i := range answer {
// 		answer[i] = make([]int, cols)
// 	}

// 	for i := 0; i < rows; i++ {
// 		for j := 0; j < cols; j++ {
// 			answer[i][j] = arr1[i][j] + arr2[i][j]
// 		}
// 	}
// 	fmt.Println("answer", answer)
// 	return answer
// }

func solution20(arr1 [][]int, arr2 [][]int) [][]int {
	for i := range arr1 {
		for j := range arr1[i] {
			arr1[i][j] += arr2[i][j]
		}
	}

	return arr1
}

func main() {
	arr1 := [][]int{{1, 2}, {2, 3}}
	arr2 := [][]int{{3, 4}, {5, 6}}
	solution20(arr1, arr2)
}
