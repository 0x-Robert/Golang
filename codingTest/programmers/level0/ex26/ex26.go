package main

import (
	"fmt"
	"math"
)

func solution(n int) int {

	a := n % 7
	b := n / 7

	if a == 0 {
		fmt.Println(b)
		return b
	} else {
		fmt.Println(b + 1)
		return b + 1
	}

	// 7 이하는 무조건 1판이 필요함
	// 7초과 14 이하 사이는 2판이 필요함
	// 15 초과 21 이하 사이는 3판이 필요함

}

func solution2(n int) int {
	return int(math.Ceil(float64(n) / 7.0))
}

func main() {
	n := 15
	solution(n)
}
