package main

import "fmt"

func solution(num int) int {

	count := 0
	if num == 1 {
		return 0
	}
	for num != 1 {
		if num%2 == 0 {

			num /= 2
			count += 1
		} else {

			num = num*3 + 1
			count += 1
		}
		if count == 500 {
			fmt.Println("-1")
			return -1
		}
	}
	fmt.Println("count", count)
	return count
}

func solution2(num int) int {
	var count int

	for count = 0; count < 500; count++ {

		if num == 1 {
			break
		}

		if num%2 == 0 {
			num = num / 2
		} else {
			num = num*3 + 1
		}
	}
	if count == 500 {
		count = -1
	}

	return count
}

func main() {
	//n := 6
	//n := 16
	n := 626331
	solution(n)
}
