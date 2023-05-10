package main

import (
	"strings"
)

func solution(t string, p string) int {
	var count int
	count = len(p)
	repeat := len(t) - count + 1
	var answer int
	str := strings.Split(t, "")
	for i := 0; i < repeat; i++ {
		num := strings.Join(str[i:i+count], "")
		if num <= p {
			answer += 1
		}
	}

	return answer
}

func solution2(t string, p string) int {
	result := 0
	for i := 0; i < len(t)-len(p)+1; i++ {
		str := t[i : i+len(p)]
		if str <= p {
			result++
		}
	}
	return result
}

func main() {
	t := "3141592"
	p := "271"
	solution(t, p)
}
