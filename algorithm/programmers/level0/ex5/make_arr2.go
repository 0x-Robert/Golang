package main

import (
	"fmt"
	"strconv"
)

func check(num int) bool {
	str := strconv.Itoa(num)
	for _, r := range str {
		if r != '5' && r != '0' {
			return false
		}
	}
	return true
}

func solution(l int, r int) []int {
	result := []int{}
	for i := l / 5 * 5; i <= r; i += 5 {
		if i >= l && i <= r && check(i) {
			result = append(result, i)
		}
	}
	if len(result) == 0 {
		return []int{-1}
	}
	fmt.Println("result", result)
	return result
}

func main() {
	l := 5
	r := 555
	solution(l, r)
}
