package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// https://namu.wiki/w/%EC%A0%95%EB%A0%AC%20%EC%95%8C%EA%B3%A0%EB%A6%AC%EC%A6%98

// 특징 : k번째원소를 1부터 k-1까지 비교해 적절한 위치에 넣고 그 뒤의 자료를 한칸씩 밀어내는 방식
// 장점 : 평균적으로 O(n^2) 중 빠른 편이다. / 이미 정렬되어있는 자료구조에 자료를 하나씩 삽입/제거하는 경우에는 훌륭한 알고리즘이된다.
// 그 이유는 탐색을 제외한 오버헤드가 적기 때문이다. 그리고 배열이 작을 경우 상당히 효율적이다. 구현이 매우 쉽다.
// 단점 : 경우에 따라서는 뒤로 밀어내는 시간이 크다.
// 시간복잡도 : O(n²)
// 대개 계산 시간이 정렬할 자료의 수의 제곱에 비례해서 늘어난다. 즉, 1만 개를 1초에 정렬하면
// 10만 개를 정렬하는 데에는 100초 정도가 필요하다.

func Insertion[T constraints.Ordered](arr []T) []T {
	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {

		temporary := arr[currentIndex]
		iterator := currentIndex
		fmt.Println("arr1", arr)
		for ; iterator > 0 && arr[iterator-1] > temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}

		arr[iterator] = temporary
		fmt.Println("arr2", arr)
	}
	return arr
}

func main() {
	arr := []int{1, 4, 5, 8, 2, 3, 50, 17, 6}
	Insertion(arr)
}
