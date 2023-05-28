package main

import (
	"strconv"
)

func solution(a int, b int) int {

	c := strconv.Itoa(a) + strconv.Itoa(b)
	d := strconv.Itoa(b) + strconv.Itoa(a)

	c1, _ := strconv.Atoi(c)
	d1, _ := strconv.Atoi(d)
	if c1 > d1 {
		return c1
	} else if c1 < d1 {
		return d1
	} else {
		return c1
	}
}

func solution0(a int, b int) int {
	ab, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	ba, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))

	if ab > ba {
		return ab
	} else {
		return ba
	}
}
func main() {
	a := 9
	b := 91
	solution(a, b)
}
