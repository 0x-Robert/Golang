package main

import (
	"fmt"
	"sort"
)

func solution(arr []int, divisor int) []int {

	answer := []int{}
	for _, v := range arr {
		if v%divisor == 0 {
			answer = append(answer, v)
		}
	}
	if len(answer) == 0 {
		answer = append(answer, -1)
		return answer
	}
	fmt.Println("answer", answer)
	sort.Sort(sort.IntSlice(answer))
	return answer
}

func main() {

	arr := []int{5, 9, 7, 10}
	divisor := 5
	solution(arr, divisor)
}
