package main

import "fmt"

func solution(array []int) int {

	//map 사용해서 int : int 의 키 밸류 선언
	count := make(map[int]int)
	fmt.Println("count", count)
	for _, num := range array {
		count[num]++
	}

	maxCount := 0
	maxNum := 0

	for num, cnt := range count {
		if cnt > maxCount {
			maxCount = cnt
			maxNum = num
		}
	}

	modeCount := 0
	for _, cnt := range count {
		if cnt == maxCount {
			modeCount++
		}
	}
	if modeCount > 1 {
		maxNum = -1
	}
	//fmt.Println("maxNum", maxNum)
	return maxNum
}

func main() {
	arr := []int{1}
	//arr := []int{1, 1, 2, 2}
	//arr := []int{1, 2, 3, 3, 3, 4}
	solution(arr)
}
