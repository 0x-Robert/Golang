package main

// func solution(n int) int {
// 	answer := 0
// 	for val := 1; val <= n; val++ {
// 		if n%val == 0 {
// 			answer += val

// 		}
// 	}
// 	//fmt.Println("answer ", answer)

// 	return answer
// }

//반환값을 명시한다.

func solution1(n int) (result int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result += i
		}
	}
	//fmt.Println(result)
	return
}

func main() {
	n := 12
	solution1(n)
}
