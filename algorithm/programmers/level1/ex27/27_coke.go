package main

import "fmt"

func solution4(a int, b int, n int) int {
	answer := 0

	for n >= a {
		answer += (n / a) * b
		n = ((n / a) * b) + (n % a)
	}

	return answer
}

func solution3(a int, b int, n int) int {

	var ans int
	for n >= a {
		ans += n / a * b
		n = n/a*b + n%a
	}

	return ans
}

// func solution2(a int, b int, n int) int {
// 	var result int
// 	var left int = 0

// 	// var quotient int = 0

// 	for n > a {
// 		// n = 20 / a = 2 / b = 1
// 		// 답 = result
// 		// 몫 = quo
// 		// 나머지 = left
// 		left = n % a

// 		n = n / a
// 		if left == 0 {

// 			result += n

// 		} else {
// 			n = n + left
// 			result += n

// 		}
// 		fmt.Println("left,quo,n,a,result", left, n, a, result)

// 	}

// 	fmt.Println("result2", result)
// 	return result
// }

/*

콜라 문제에서 사용된 알고리즘은 재귀(recursion)와 수학적 귀납법(induction)입니다.

재귀(recursion)는 함수가 자신을 다시 호출하여 문제를 해결하는 기법입니다. 콜라 문제에서는 주어진 조건에 따라 재귀적으로 함수를 호출하여 문제를 해결했습니다.

수학적 귀납법(induction)은 수학적 명제를 증명하는 기법 중 하나입니다. 이 방법은 기본 단계(base case)와 귀납 단계(inductive step)로 이루어져 있습니다. 콜라 문제에서는 빈 병의 개수가 일정한 수 이상일 때 콜라를 받을 수 있는 최대 병 수를 구하는데, 이를 귀납법으로 증명했습니다. 빈 병의 개수가 a개 이상일 때, 상빈이가 받을 수 있는 최대 병 수는 a/b + f(a%b)개임을 증명했습니다. 이 때, a%b는 a를 b로 나눈 나머지이며, f()는 a%b를 입력받아 콜라를 받을 수 있는 최대 병 수를 반환하는 함수입니다. 이러한 방식으로 귀납적으로 문제를 해결했습니다.
*/

func solution1(a int, b int, n int) int {
	emptyBottles := n // 처음에 가지고 있는 빈 병의 수
	totalColas := 0   // 받은 전체 콜라 병의 수

	for emptyBottles >= a { // 보유 중인 빈 병이 a개 이상인 동안 반복

		colaBottles := emptyBottles / a // 현재 가지고 있는 빈 병으로 받을 수 있는 콜라 병의 수
		totalColas += colaBottles * b   // 받은 콜라 병의 수를 누적

		emptyBottles = (emptyBottles % a) + (colaBottles * b) // 마신 콜라 병의 수를 빈 병으로 교환하여 보유 중인 빈 병의 수 업데이트
		fmt.Println("cola, total, empty", colaBottles, totalColas, emptyBottles)
	}
	fmt.Println(totalColas)
	return totalColas
}

func solution0(a int, b int, n int) int {
	if n > b {
		fmt.Println((n - b) / (a - b) * b)
		return (n - b) / (a - b) * b
	} else {
		return 0
	}
}

func main() {
	// a := 2
	// b := 1
	// n := 20

	a := 3
	b := 1
	n := 20
	solution0(a, b, n)
}
