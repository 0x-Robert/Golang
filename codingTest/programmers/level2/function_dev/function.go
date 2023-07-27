package main

import "fmt"

func solution(progresses []int, speeds []int) []int {
	answer := []int{}
	time := 0
	count := 0

	for len(progresses) > 0 {
		if (progresses[0] + time*speeds[0]) >= 100 {
			progresses = progresses[1:]
			speeds = speeds[1:]
			count++
		} else {
			if count > 0 {
				answer = append(answer, count)
				count = 0
			}
			time++
		}
	}
	answer = append(answer, count)
	fmt.Println(answer)
	return answer
}

func main() {
	progresses := []int{90, 30, 55}
	speeds := []int{1, 30, 5}
	solution(progresses, speeds)
}
