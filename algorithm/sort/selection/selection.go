package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// https://ko.wikipedia.org/wiki/%EC%84%A0%ED%83%9D_%EC%A0%95%EB%A0%AC

// 특징
// 주어진 리스트에서 최소값을 찾는다.
// 그 값을 앞의 값과 교체한다.
// 앞의 위치를 뺀 나머지 리스트를 위와 같은 방식으로 계속 정렬한다.
// 장점 : 버블정렬보다 2배 빠르다. , 합병정렬에 비해서 작은 배열 요소 10개에서 20개 미만일 때 빠르다.
// 단점 : 시간복잡도가 O(n^2)
// 시간복잡도 : O(n²)

// 기타 다른 정렬과의 비교
// 거품 정렬(bubble sort) : 시간 복잡도 Θ ( n 2 )인 정렬 알고리즘 중에서
// 선택 정렬은 버블 정렬보다 항상 우수하다.

// 삽입 정렬(insertion sort) : 삽입 정렬은 k번째 반복 이후, 첫번째 k 요소가 정렬된 순서로 온다는 점에서 유사하다.
// 하지만 선택 정렬은 k+1 번째 요소를 찾기 위해 나머지 모든 요소들을 탐색하지만 삽입 정렬은 k+1 번째 요소를 배치하는 데 필요한 만큼의 요소만 탐색하기 때문에 훨씬 효율적으로 실행된다는 차이가 있다.
// 즉 삽입정렬이 선택정렬보다 뛰어나다.

// 합병 정렬(merge sort) : 선택 정렬은 합병 정렬과 같은 분할 정복 알고리즘을 사용하지만
// 일반적으로 큰 배열보다 작은 배열(요소 10~20개 미만)에서 더 빠르다.
// 따라서 충분히 작은 하위 목록에만 삽입 정렬 혹은 선택 정렬을 이용해서 최적화하는 것이 좋다.
// 즉 선택정렬은 합병정렬에 비해 작은 배열 요소 10개에서 20개 미만일 때 빠르다.

func Selection[T constraints.Ordered](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				fmt.Println("j:", j)
				min = j
			}
		}
		fmt.Println("min :", min, "i :", i, "arr :", arr)
		arr[i], arr[min] = arr[min], arr[i]
		fmt.Println("min :", min, "i :", i, "arr :", arr)
	}
	return arr
}

func main() {
	arr := []int{5, 20, 45, 1, 7, 100, 3}
	Selection(arr)
}
