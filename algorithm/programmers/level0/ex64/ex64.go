// 합성수 찾기
// 문제 설명
// 약수의 개수가 세 개 이상인 수를 합성수라고 합니다. 자연수 n이 매개변수로 주어질 때 n이하의 합성수의 개수를 return하도록 solution 함수를 완성해주세요.

// 제한사항
// 1 ≤ n ≤ 100
// 입출력 예
// n	result
// 10	5
// 15	8
// 입출력 예 설명
// 입출력 예 #1

// 10 이하 합성수는 4, 6, 8, 9, 10 로 5개입니다. 따라서 5를 return합니다.
// 입출력 예 #1

// 15 이하 합성수는 4, 6, 8, 9, 10, 12, 14, 15 로 8개입니다. 따라서 8을 return합니다.

package main

import "fmt"

func checkdiv(n int) bool {
	for j := 2; j < n; j++ {
		if n%j == 0 {
			// fmt.Println("j", j)
			return true
		}
	}
	return false
}

// 1과 자기자신은 모든 수가 약수로 가진다. 그렇다면 n 이하 중에서 1과 자기자신을 제외한 수로 나눠지면 합성수에 해당한다.
// 그래서 2부터 나눈다. 처음 테스트케이스로 실행하면 4,6,8,9,10이 나온다.두 번째 테스트케이스로 n에 15를 대입하면 4, 6, 8, 9, 10, 12, 14, 15가 나온다.
func solution(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if checkdiv(i) == true {
			fmt.Println("i", i)
			count += 1
		}
	}
	fmt.Println(count)
	return count
}

func main() {
	// n := 15
	n := 10
	solution(n)
}
