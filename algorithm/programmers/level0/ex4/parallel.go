package main

import (
	"reflect"
	"sort"
)

// func solution(dots [][]int) int {

// 	sort.Slice(dots, func(i, j int) bool {
// 		return dots[i][0] < dots[j][0]
// 	})
// 	//	fmt.Println("dots", dots[0], dots[1])
// 	if reflect.DeepEqual(dots[0], dots[1]) {
// 		// fmt.Println("1")
// 		return 1
// 	} else {
// 		if (dots[1][1]-dots[0][1])/(dots[1][0]-dots[0][0]) == (dots[3][1]-dots[2][1])/(dots[3][0]-dots[2][0]) {
// 			// fmt.Println("1")
// 			return 1
// 		} else {
// 			// fmt.Println("0")
// 			return 0
// 		}
// 	}
// }

func solution(dots [][]int) int {
	// 첫번째 점과 두번째 점을 잇는 선분과
	// 세번째 점과 네번째 점을 잇는 선분의 기울기가 같으면
	// 두 선분은 평행이다.

	x1, y1 := dots[0][0], dots[0][1]
	x2, y2 := dots[1][0], dots[1][1]
	x3, y3 := dots[2][0], dots[2][1]
	x4, y4 := dots[3][0], dots[3][1]

	sort.Slice(dots, func(i, j int) bool {
		return dots[i][0] < dots[j][0]
	})
	if reflect.DeepEqual(dots[0], dots[1]) || reflect.DeepEqual(dots[2], dots[3]) {
		return 1
	} else {
		if (y2-y1)*(x4-x3) == (y4-y3)*(x2-x1) {
			return 1
		}
		return 0
	}

}

func solution0(dots [][]int) int {


	x1, y1 := dots[0][0], dots[0][1]
	x2, y2 := dots[1][0], dots[1][1]
	x3, y3 := dots[2][0], dots[2][1]
	x4, y4 := dots[3][0], dots[3][1]

	if (x1-x2)*(y3-y4) == (y1-y2)*(x3-x4) { //1234
		return 1
	}
	if (x1-x3)*(y2-y4) == (y1-y3)*(x2-x4) { //1324 
		return 1
	}
	if (x1-x4)*(y2-y3) == (y1-y4)*(x2-x3) { //1423
		return 1
	}
	return 0
}

func main() {

	//dots := [][]int{{1, 4}, {9, 2}, {3, 8}, {11, 6}}
	dots := [][]int{{1, 4}, {9, 2}, {1, 4}, {9, 2}}
	//dots := [][]int{{3, 5}, {4, 1}, {2, 4}, {5, 10}}
	solution(dots)

}
