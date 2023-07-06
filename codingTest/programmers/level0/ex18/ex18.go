package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string
	var a string
	var answer string
	fmt.Scan(&s1, &a)
	num, _ := strconv.Atoi(a)
	for i := 0; i < num; i++ {
		answer += s1
	}
	fmt.Println(answer)
}
