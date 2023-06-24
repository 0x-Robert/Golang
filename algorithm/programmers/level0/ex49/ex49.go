package main

func solution(n int, k int) int {
	k = (k - (n / 10)) * 2000
	n = n * 12000
	return n + k
}

func solution1(n int, k int) int {
	return n*12000 + (k-n/10)*2000
}

func main() {
	n := 10
	k := 3
	solution(n, k)
}
