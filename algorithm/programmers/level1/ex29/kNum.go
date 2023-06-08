package main

import "fmt"

func solution(array []int, commands [][]int) []int {
	fmt.Println("array", array, "commands", commands)
	return []int{}
}

func main() {
	arr := []int{1, 5, 2, 6, 3, 7, 4}
	commands := []int{{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}}
	solution(arr, commands)
}
