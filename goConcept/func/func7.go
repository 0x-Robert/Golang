package main

import "fmt"

func Multiple(a, b int) int {
	return a * b
}

func main() {
	result := Multiple(3, 4)
	fmt.Println(result)
}
