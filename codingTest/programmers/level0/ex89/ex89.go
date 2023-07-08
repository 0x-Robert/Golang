// 자릿수 더하기
// 문제 설명
// 정수 n이 매개변수로 주어질 때 n의 각 자리 숫자의 합을 return하도록 solution 함수를 완성해주세요

// 제한사항
// 0 ≤ n ≤ 1,000,000
// 입출력 예
// n	result
// 1234	10
// 930211	16
// 입출력 예 설명
// 입출력 예 #1

// 1 + 2 + 3 + 4 = 10을 return합니다.
// 입출력 예 #2

// 9 + 3 + 0 + 2 + 1 + 1 = 16을 return합니다.

package main

import (
	"strconv"
)

func solution(n int) int {
	str := strconv.Itoa(n)
	answer := 0
	for _, v := range str {
		answer += int((v - 48))
	}

	return answer
}

func solution0(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}

func main() {
	n := 1234
	solution(n)
}
