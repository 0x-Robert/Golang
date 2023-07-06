package main

func solution(num int, n int) int {

	if num%n == 0 {
		return 1
	} else {
		return 0
	}

}

func main() {
	num := 98
	n := 2
	solution(num, n)
}
