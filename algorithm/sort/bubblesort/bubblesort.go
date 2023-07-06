package main

// https://en.wikipedia.org/wiki/Bubble_sort
// https://namu.wiki/w/%EC%A0%95%EB%A0%AC%20%EC%95%8C%EA%B3%A0%EB%A6%AC%EC%A6%98
import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 특징 : 배열의 두 수를 비교해서 위치를 비교한 뒤 작으면 앞으로 보낸다.
// 장점 : 이미 정렬된 자료에서는 1번만 돌면 되서 최선의 성능을 보여준다.
// 단점 : 가장 최악의 성능을 보여주는 정렬이다.
// 시간복잡도 : O(n²)

func Bubble[T constraints.Ordered](arr []T) []T {
	swapped := true
	for swapped {
		fmt.Println(swapped)
		swapped = false
		fmt.Println(swapped)
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				fmt.Println("before arr", arr)
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
				fmt.Println("after arr", arr)
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 4, 2, 7, 3}
	Bubble(arr)
}
