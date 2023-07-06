// 진료 순서 정하기
// 문제 설명
// 외과의사 머쓱이는 응급실에 온 환자의 응급도를 기준으로 진료 순서를 정하려고 합니다. 정수 배열 emergency가 매개변수로 주어질 때 응급도가 높은 순서대로 진료 순서를 정한 배열을 return하도록 solution 함수를 완성해주세요.

// 제한사항
// 중복된 원소는 없습니다.
// 1 ≤ emergency의 길이 ≤ 10
// 1 ≤ emergency의 원소 ≤ 100
// 입출력 예
// emergency	result
// [3, 76, 24]	[3, 1, 2]
// [1, 2, 3, 4, 5, 6, 7]	[7, 6, 5, 4, 3, 2, 1]
// [30, 10, 23, 6, 100]	[2, 4, 3, 5, 1]
// 입출력 예 설명
// 입출력 예 #1

// emergency가 [3, 76, 24]이므로 응급도의 크기 순서대로 번호를 매긴 [3, 1, 2]를 return합니다.
// 입출력 예 #2

// emergency가 [1, 2, 3, 4, 5, 6, 7]이므로 응급도의 크기 순서대로 번호를 매긴 [7, 6, 5, 4, 3, 2, 1]를 return합니다.
// 입출력 예 #3

// emergency가 [30, 10, 23, 6, 100]이므로 응급도의 크기 순서대로 번호를 매긴 [2, 4, 3, 5, 1]를 return합니다.

package main

import (
	"sort"
)

func solution(emergency []int) []int {
	// 배열 길이만큼 배열을 만든다. 원소가 3개면 1,2,3인 배열을 만든다.
	// 입력받은 배열의 크기별로 비교한 순서를 리턴한다.
	// 예를 들어 [3,76,24]인 배열을 입력받으면 아래와 같이 result로 리턴한다.
	// result := [3, 1, 2]
	emergency2 := make([]int, len(emergency))
	// 원본 슬라이스 복사
	copy(emergency2, emergency)
	// 내림차순 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(emergency)))

	// 순서를 기록할 맵 정의
	orderMap := map[int]int{}

	// map 생성
	for i, v := range emergency {
		orderMap[i+1] = v
	}

	// orderMap := map[1:76 2:24 3:3]
	// emergency2 :=  [3, 76, 24]
	// 1 3, 2 76, 3 24
	arr := []int{}
	for _, v := range emergency2 {
		for j, v2 := range orderMap {
			if v == v2 {
				arr = append(arr, j)
			}
		}
	}
	return arr
}

func solution0(emergency []int) []int {
	sorted := make([]int, len(emergency))
	copy(sorted, emergency)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

	result := make([]int, len(emergency))
	for k, v := range sorted {
		for k1, v1 := range emergency {
			if v == v1 {
				result[k1] = k + 1
			}
		}
	}
	return result
}

func main() {
	emergency := []int{3, 76, 24}
	solution(emergency)
}
