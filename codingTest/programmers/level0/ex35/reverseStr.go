package main

import "fmt"

func solution(my_string string) string {
	answer := ""
	for i := len(my_string); i > 0; i-- {
		answer += string(my_string[i-1])
	}
	fmt.Println(answer)
	return answer
}

func main() {
	my_string := "jaron"
	solution(my_string)
}
