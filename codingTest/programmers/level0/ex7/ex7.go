package main

import "fmt"

func solution(num1 int, num2 int) int {
	fmt.Println(num1 * 1000 / num2)
	return (num1 * 1000 / num2)
}
func main() {
	// 	num1 := 3
	// 	num2 := 2
	// num1 := 7
	// num2 := 3
	num1 := 1
	num2 := 16
	solution(num1, num2)
}
