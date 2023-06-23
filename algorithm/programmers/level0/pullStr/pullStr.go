package main

import "fmt"

func solution(A string, B string) int {
	for _, v := range A {
		fmt.Println(string(v))
		// fmt.Println(string(A[0]))
	}

	for _, v := range B {
		fmt.Println(string(v))
		// fmt.Println(B[0])
	}
	return 0
}

func main() {
	A := "hello"
	B := "ohell"
	solution(A, B)
}
