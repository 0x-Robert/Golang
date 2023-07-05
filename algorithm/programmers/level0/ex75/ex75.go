// 가까운 수
// 문제 설명
// 정수 배열 array와 정수 n이 매개변수로 주어질 때, array에 들어있는 정수 중 n과 가장 가까운 수를 return 하도록 solution 함수를 완성해주세요.

// 제한사항
// 1 ≤ array의 길이 ≤ 100
// 1 ≤ array의 원소 ≤ 100
// 1 ≤ n ≤ 100
// 가장 가까운 수가 여러 개일 경우 더 작은 수를 return 합니다.
// 입출력 예
// array	n	result
// [3, 10, 28]	20	28
// [10, 11, 12]	13	12
// 입출력 예 설명
// 입출력 예 #1

// 3, 10, 28 중 20과 가장 가까운 수는 28입니다.
// 입출력 예 #2

// 10, 11, 12 중 13과 가장 가까운 수는 12입니다.
// ※ 공지 - 2023년 3월 29일 테스트 케이스가 추가되었습니다. 기존에 제출한 코드가 통과하지 못할 수도 있습니다.

package main

import (
	"math"
	"sort"
)

func solution(array []int, n int) int {
	mp := make(map[int]int, len(array))
	sk := make([]int, 0, len(array))
	if len(array) == 1 {
		return array[0]
	}
	for _, v := range array {
		c := n - v
		mp[c] = v
		sk = append(sk, c)
	}
	sort.Ints(sk)
	min := sk[0]
	minIndex := 0
	for i, v := range sk {
		if math.Abs(float64(v)) < math.Abs(float64(min)) {
			min = v
			minIndex = i
		} else if math.Abs(float64(v)) == math.Abs(float64(min)) {
			if v > min {
				min = v
				minIndex = i
			}
		}
	}
	// fmt.Println(mp[sk[minIndex]])
	return mp[sk[minIndex]]
}


func solution1(array []int, n int) int {
    answer, min := 0, int(math.MaxInt32)
    for _, num := range array {
        t := int(math.Abs(float64(num-n)))
        if min > t {
            answer = num
            min = t
        } else if min == t && answer > num {
            answer = num
        }
    }

    return answer
}








func main() {
	// arr := []int{3, 10, 28}
	arr := []int{10, 11, 12, 14}
	// arr := []int{3}
	// arr := []int{1, 4, 2, 1}
	// arr := []int{1, 2, 1, 1}
	// arr := []int{2, 3, 0, 4}
	solution(arr, 13)
}
