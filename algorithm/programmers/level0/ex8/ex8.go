package main

import (
	"fmt"
	"math/big"
)

func solution(numer1 int, denom1 int, numer2 int, denom2 int) []int {
	// 분수를 big.Rat 타입으로 생성
	rat1 := big.NewRat(int64(numer1), int64(denom1))
	rat2 := big.NewRat(int64(numer2), int64(denom2))

	// 분수 덧셈
	sum := new(big.Rat).Add(rat1, rat2)

	// 분자와 분모 추출
	num := sum.Num().Int64()
	denom := sum.Denom().Int64()

	// 분수를 기약 분수로 만들기 위해 최대공약수를 구하고 분자와 분모를 나눔
	gcd := gcd(num, denom)
	num /= gcd
	denom /= gcd

	return []int{int(num), int(denom)}
}

// 최대공약수를 구하는 함수
func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	result := solution(1, 2, 3, 4)
	fmt.Println(result)

	result = solution(9, 2, 1, 3)
	fmt.Println(result)
}
