package main

import "fmt"

//피자조각 슬라이스와 먹는사람 n이 주어질 때 최소 1명당 1조각씩이라도 먹으려면 피자를 몇판 주문해야하는지 계산하시오

func solution(slice int, n int) int {
	answer := 0
	for i := 1; i <= n; i++ {
		if slice*i >= n {
			answer = i
			break
		}
	}
	fmt.Println(answer)
	return answer
}

func solution2(slice int, n int) int {
	return (n + slice - 1) / slice
}
func main() {
	// slice := 7
	// n := 10
	slice := 4
	n := 12
	solution(slice, n)
}
