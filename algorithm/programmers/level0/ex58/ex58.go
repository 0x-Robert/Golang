// 구슬을 나누는 경우의 수
// 문제 설명
// 머쓱이는 구슬을 친구들에게 나누어주려고 합니다. 구슬은 모두 다르게 생겼습니다. 머쓱이가 갖고 있는 구슬의 개수 balls와 친구들에게 나누어 줄
// 구슬 개수 share이 매개변수로 주어질 때, balls개의 구슬 중 share개의 구슬을 고르는 가능한 모든 경우의 수를 return 하는 solution 함수를 완성해주세요.

// 제한사항
// 1 ≤ balls ≤ 30
// 1 ≤ share ≤ 30
// 구슬을 고르는 순서는 고려하지 않습니다.
// share ≤ balls
// 입출력 예
// balls	share	result
// 3	2	3
// 5	3	10
// 입출력 예 설명
// 입출력 예 #1

// 서로 다른 구슬 3개 중 2개를 고르는 경우의 수는 3입니다. 스크린샷 2022-08-01 오후 4.15.55.png
// 입출력 예 #2

// 서로 다른 구슬 5개 중 3개를 고르는 경우의 수는 10입니다.
// Hint
// 서로 다른 n개 중 m개를 뽑는 경우의 수 공식은 다음과 같습니다. 스크린샷 2022-08-01 오후 4.37.53.png

package main

import "fmt"

func p2(n int, m int) uint64 {
	var count uint64
	count = 1
	for i := n; i >= m; i-- {
		count = count * uint64(i)
	}
	return count
}

func pactorial(n int) uint64 {
	var count uint64
	count = 1
	for i := n; i > 0; i-- {
		count = count * uint64(i)
	}
	return count
}

func solution(balls int, share int) int {
	if balls == share || share == 0 {
		return 1
	}
	// balls = n, share = m
	// n! / (n-m)! * m!
	// n := int(pactorial(balls))
	// m := (int(pactorial(balls-share)) * int(pactorial(share)))

	n := p2(balls, balls-share+1)
	m := pactorial(share)
	answer := int(n / m)
	// fmt.Println("answer", answer)
	return answer
}

func solution0(balls, share int) int {
	dp := make([][]int, balls+1)
	fmt.Println("dp", dp)
	for i := 0; i <= balls; i++ {
		dp[i] = make([]int, share+1)
		fmt.Println("dp", dp[i])
	}
	fmt.Println("dp", dp)

	for i := 0; i <= balls; i++ {
		dp[i][0] = 1
	}

	fmt.Println("dp", dp)

	for i := 1; i <= balls; i++ {
		for j := 1; j <= share; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}

	return dp[balls][share]
}

func solution1(balls int, share int) int {
	answer := 1
	for i := 0; i < share; i++ {
		answer *= balls - i
		answer /= i + 1
	}
	return answer
}

func main() {
	// balls := 3
	// share := 2
	// balls := 5
	// share := 3
	balls := 30
	share := 15

	solution0(balls, share)
}
