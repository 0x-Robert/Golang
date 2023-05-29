package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n)
	var a string
	for i := 0; i < n; i++ {
		a += "*"
		fmt.Println(a)
	}
}
