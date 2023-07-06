package main

import (
	"fmt"
)

//flag에따라 다른 값 반환하기

func solution(a int, b int, flag bool) int {

	if flag == true {
		fmt.Println(a + b)
		return a + b
	} else {
		fmt.Println(a - b)
		return a - b
	}

	return 0
}

func main() {
	a := -4
	b := 7
	flag := true
	solution(a, b, flag)
}
