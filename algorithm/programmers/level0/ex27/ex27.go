package main

import "fmt"

func solution(n int) int {

	result := 0
	// 6명이 주문한 피자를 남기지 않고 모두 같은 수의 피자조각을 먹게한다면 최소 몇판을 먹어야하는지 계산
	for i := 1; i <= n; i++ {
		if 6*i%n == 0 {
			result = i
			break
		}
	}
	fmt.Println(result)
	return result

}

func main() {
	n := 4
	solution(n)
}
