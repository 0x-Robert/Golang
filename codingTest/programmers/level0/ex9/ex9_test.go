package main

import "testing"

func TestSolution(t *testing.T) {
	a1 := 2
	b1 := 2
	c1 := 2
	d1 := 2

	a2 := 4
	b2 := 1
	c2 := 4
	d2 := 4

	a3 := 6
	b3 := 3
	c3 := 3
	d3 := 6

	a4 := 2
	b4 := 5
	c4 := 2
	d4 := 6

	a5 := 6
	b5 := 4
	c5 := 2
	d5 := 5
	Solution(a1, b1, c1, d1)
	Solution(a2, b2, c2, d2)
	Solution(a3, b3, c3, d3)
	Solution(a4, b4, c4, d4)
	Solution(a5, b5, c5, d5)
}
