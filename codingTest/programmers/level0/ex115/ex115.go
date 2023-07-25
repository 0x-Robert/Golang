// k의 개수
// 문제 설명
// 1부터 13까지의 수에서, 1은 1, 10, 11, 12, 13 이렇게 총 6번 등장합니다. 정수 i, j, k가 매개변수로 주어질 때, i부터 j까지 k가 몇 번 등장하는지 return 하도록 solution 함수를 완성해주세요.

// 제한사항
// 1 ≤ i < j ≤ 100,000
// 0 ≤ k ≤ 9
// 입출력 예
// i	j	k	result
// 1	13	1	6
// 10	50	5	5
// 3	10	2	0
// 입출력 예 설명
// 입출력 예 #1

// 본문과 동일합니다.
// 입출력 예 #2

// 10부터 50까지 5는 15, 25, 35, 45, 50 총 5번 등장합니다. 따라서 5를 return 합니다.
// 입출력 예 #3

// 3부터 10까지 2는 한 번도 등장하지 않으므로 0을 return 합니다.

package main

import (
	"strconv"
	"strings"
)

func solution(i int, j int, k int) int {
	t := ""
	k2 := strconv.Itoa(k)
	answer := 0
	for a := i; a <= j; a++ {
		t = strconv.Itoa(a)
		// if strings.Contains(t, k2) {
		// 	answer += 1
		// }
		answer += strings.Count(t, k2)
	}
	//	fmt.Println(answer)
	return answer
}

func main() {
	i := 1
	j := 13
	k := 1
	// i := 3
	// j := 10
	// k := 2
	solution(i, j, k)
}
