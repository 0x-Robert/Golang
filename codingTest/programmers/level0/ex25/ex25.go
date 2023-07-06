package main

import "fmt"

// func solution(arr []string) string {

// 	answer := ""
// 	for _, v := range arr {
// 		answer += string(v)
// 	}

// 	fmt.Println(answer)
// 	return answer
// }

func solution(my_string string, k int) string {

	answer := ""
	for i := 0; i < k; i++ {
		answer += my_string
	}
	fmt.Println(answer)
	return answer
}

func main() {
	//arr := []string{"a", "b", "c"}
	my_string := "string"
	k := 3
	solution(my_string, k)
}
