package main

func solution(numbers []int) float64 {
	var answer float64
	var sum float64
	for _, v := range numbers {
		sum += float64(v)
	}
	// fmt.Println(sum / float64(len(numbers)))
	answer = sum / float64(len(numbers))
	// fmt.Println(answer)
	return answer
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	solution(numbers)
}
