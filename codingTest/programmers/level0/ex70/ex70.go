// 소인수분해
// 문제 설명
// 소인수분해란 어떤 수를 소수들의 곱으로 표현하는 것입니다. 예를 들어 12를 소인수 분해하면 2 * 2 * 3 으로 나타낼 수 있습니다. 따라서 12의 소인수는 2와 3입니다. 자연수 n이 매개변수로 주어질 때 n의 소인수를 오름차순으로 담은 배열을 return하도록 solution 함수를 완성해주세요.

// 제한사항
// 2 ≤ n ≤ 10,000
// 입출력 예
// n	result
// 12	[2, 3]
// 17	[17]
// 420	[2, 3, 5, 7]
// 입출력 예 설명
// 입출력 예 #1

// 12를 소인수분해하면 2 * 2 * 3 입니다. 따라서 [2, 3]을 return합니다.
// 입출력 예 #2

// 17은 소수입니다. 따라서 [17]을 return 해야 합니다.
// 입출력 예 #3

// 420을 소인수분해하면 2 * 2 * 3 * 5 * 7 입니다. 따라서 [2, 3, 5, 7]을 return합니다.
// 테스트케이스 중 입력값 n이 252일 때 패스 되도록 설정

package main

import "fmt"

func check(pr int) bool {
	for j := 2; j < pr; j++ {
		if pr%j == 0 {
			return false
		}
	}
	return true
}

func solution(n int) []int {
	arr := []int{}
	for i := 2; i <= n; i++ {
		if n%i == 0 && check(i) {
			arr = append(arr, i)
			n = n / i
			fmt.Println(arr, n, i)
		}
	}
	return arr
}

func main() {
	// n := 12
	// n := 17
	// n := 420
	n := 252
	solution(n)
}
