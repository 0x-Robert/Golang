package main

import "fmt"

func solution(numbers []int) []int {

	for i, v := range numbers {
		numbers[i] = 2 * v
	}
	fmt.Println(numbers)
	return numbers
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	solution(numbers)
}
