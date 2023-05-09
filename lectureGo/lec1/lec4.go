package main

import "fmt"

func tempFunc(i ...int) {
	fmt.Println(i)
}

func main() {
	tempFunc(1)
	tempFunc(1, 2)
	tempFunc(1, 2, 3)
}
