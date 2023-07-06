package main

import "fmt"

// 점 3개를 찍어서 가변인수를 표현함
func sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func main() {
	fmt.Printf("1부터 5까지 합은 %d입니다.\n", sum(1, 2, 3, 4, 5))
}
