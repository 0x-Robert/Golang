package main

import (
	"sort"
)

func solution(array []int) int {
	sort.Sort(sort.IntSlice(array))
	// fmt.Println(array[len(array)/2])
	return array[len(array)/2]
}

func main() {
	array := []int{1, 2, 7, 10, 11}
	solution(array)
}
