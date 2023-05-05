package main

import "fmt"

func solution(n int) string {

	var result string
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result += "수"
		} else {
			result += "박"
		}
	}
	fmt.Println("result", result)
	return result
}

func solution2(n int) (ret string) {
	for i := 0; i < n/2; i++ {
		ret += "수박"
	}
	if n%2 == 1 {
		ret += "수"
	}
	return
}

func main() {
	n := 3
	solution(n)
}
