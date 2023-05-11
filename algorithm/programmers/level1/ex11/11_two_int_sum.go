package main

func solution11(a int, b int) int64 {
	sum := 0

	if a > b {
		a, b = b, a
	}

	for i := a; i <= b; i++ {
		sum += i
	}

	return int64(sum)
}

func main() {
	a := 3
	b := 3
	solution11(a, b)
}
