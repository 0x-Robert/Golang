package main

import "fmt"

func solution2(a int, b int, n int) int {
	var result int
	var left int = 0
	var quo int = 0
	// var quotient int = 0

	for left < n {
		// n = 20 / a = 2 / b = 1
		// 답 = result
		// 몫 = quo
		// 나머지 = left
		left = n % a
		quo = n / a
		n = n / a
		if left == 0 {
			n = quo
			result += n

		} else {
			n = quo + left
			result += quo

		}
		fmt.Println("left,quo,n,a,result", left, quo, n, a, result)

	}

	fmt.Println("result2", result)
	return result
}

func solution(a int, b int, n int) int {
	emptyBottles := n // 처음에 가지고 있는 빈 병의 수
	totalColas := 0   // 받은 전체 콜라 병의 수

	for emptyBottles >= a { // 보유 중인 빈 병이 a개 이상인 동안 반복
		colaBottles := emptyBottles / a // 현재 가지고 있는 빈 병으로 받을 수 있는 콜라 병의 수
		totalColas += colaBottles * b   // 받은 콜라 병의 수를 누적

		emptyBottles = (emptyBottles % a) + (colaBottles * b) // 마신 콜라 병의 수를 빈 병으로 교환하여 보유 중인 빈 병의 수 업데이트
	}
	fmt.Println(totalColas)
	return totalColas
}

func main() {
	// a := 2
	// b := 1
	// n := 20

	a := 3
	b := 1
	n := 20
	solution(a, b, n)
}
