package main

import "fmt"

func solution(num_list []int) []int {
	answer := []int{}

	for i := len(num_list); i > 0; i-- {
		answer = append(answer, num_list[i-1])
	}
	fmt.Println(answer)
	return answer
}

func main() {
	// num_list := []int{1, 2, 3, 4, 5}
	num_list := []int{1, 1, 1, 1, 1, 2}
	solution(num_list)
}
