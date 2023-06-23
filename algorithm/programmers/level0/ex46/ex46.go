package main

import (
	"fmt"
	"strings"
)

func solution(my_string string, n int) string {
	var answer string

	for _, v := range my_string {
		answer += strings.Repeat(string(v), n)
	}
	fmt.Println("answer", answer)
	return answer
}

func main() {
	my_string := "hello"
	n := 3
	solution(my_string, n)
}
