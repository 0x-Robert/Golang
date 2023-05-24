package main

import "fmt"

func solution(str1 string, str2 string) string {
	answer := ""
	for i := 0; i < len(str1); i++ {
		answer += string(str1[i])
		answer += string(str2[i])
	}
	fmt.Println(answer)
	return answer
}

func main() {
	str1 := "aaaaa"
	str2 := "bbbbb"
	solution(str1, str2)
}
