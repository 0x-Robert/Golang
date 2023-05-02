package main

func solution(x int) bool {

	temp := x
	final := 0
	for temp > 0 {
		final += int(temp % 10)
		temp /= 10
	}

	return x%final == 0

}

func main() {
	x := 13
	solution(x)
}
