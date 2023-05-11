package main

func solution(common []int) int {
	answer := 0
	
	if common[2]-common[1] == common[1]-common[0] {
		answer = common[len(common)-1] + (common[1] - common[0])
	} else {
		answer = common[len(common)-1] * (common[1] / common[0])
	}
	return answer
}

func main() {
	common := []int{1, 2, 3, 4}
	//common := []int{2, 4, 8}
	solution(common)
}
