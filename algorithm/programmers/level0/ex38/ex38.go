package main

import (
	"strconv"
)

func solution(a int, b int) int {
	c, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	d := 2 * a * b
	if c > d {
		return c
	} else if c < d {
		return d
	} else {
		return c
	}

}
func main() {
	a := 9
	b := 91
	solution(a, b)
}
