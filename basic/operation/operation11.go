package main

import "fmt"

func main() {
	var x int8 = 1
	x <<= 7
	x >>= 7
	fmt.Printf("%d\n", x)
}
