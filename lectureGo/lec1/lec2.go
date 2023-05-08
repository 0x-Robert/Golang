package main

import (
	"fmt"
)

// 문제1 두 인자의 합과 곱을 리턴하는 multiCalc 함수 부분을 채우세요
func multiCalc(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return x, y
}

func main() {
	i, j := multiCalc(10, 20)
	fmt.Printf("Sum :%d, Mul : %d", i, j)
}

//문제2 3개 이상(합산, 곱, 나누기 등) 인수를 리턴하는 함수를 자율적으로 작성
