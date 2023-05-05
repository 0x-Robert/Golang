package main

import (
	"fmt"
	"sort"
)

func solution(s string) string {

	arr := []string{}
	for _, v := range s {
		arr = append(arr, string(v))
	}
	sort.Sort(sort.Reverse(sort.StringSlice(arr)))

	result := ""
	for _, v := range arr {
		result += v
	}

	return result
}

/*

먼저 문자열 s를 []byte로 변환하여 a 변수에 저장합니다.
그리고 sort.Slice 함수를 사용하여 a 변수를 내림차순으로 정렬합니다.
정렬에 사용되는 함수는 func(i, j int) bool { return a[i] > a[j] } 형태로 익명 함수를 사용합니다.
이 함수는 a[i]가 a[j]보다 큰 경우 true를 반환하고, 작거나 같은 경우에는 false를 반환합니다.

마지막으로, 정렬된 a 변수를 다시 문자열로 변환하여 반환합니다.
이 때 string() 함수를 사용하여 []byte 타입인 a를 문자열로 변환합니다.

*/

func solution2(s string) string {
	fmt.Println("s", s)
	a := []byte(s)
	fmt.Println("a", a)
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Println("a", a)
	fmt.Println("a", string(a))
	return string(a)
}

func main() {

	a := "Zbcdefg"
	solution2(a)
}
