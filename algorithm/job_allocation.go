package main

import (
	"fmt"
	"math"
)

func jobAllocation(jobs []int, workersNum int) int {
	//이진 탐색을 위한 최소, 최대 설정
	left := math.MinInt64
	for _, job := range jobs {
		if job > left {
			left = job
		}
	}

	right := 0
	for _, job := range jobs {
		right += job
	}

	for left <= right {
		mid := (left + right) / 2
		//현재 mid값을 가지고 분배를 해보고, 그 결과에 따라서 left, right 값을 조정합니다.
		sum := 0
		workers := 1

		for _, job := range jobs {
			sum += job
			if sum > mid {
				sum = job
				workers++
			}
		}

		if workers > workersNum {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	jobs := []int{1, 2, 3, 4, 5, 6, 7}
	workerNum := 3
	output := jobAllocation(jobs, workerNum)
	fmt.Println(output)

	jobs = []int{10, 2, 3, 4, 16, 10, 10}
	workersNum2 := 4
	output = jobAllocation(jobs, workersNum2)
	fmt.Println(output) // --> 19 (10, 2, 3, 4 / 16 / 10 / 10)
}
