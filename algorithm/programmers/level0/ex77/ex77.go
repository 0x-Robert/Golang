// 취급합니다.
// 입출력 예
// cipher	code	result
// "dfjardstddetckdaccccdegk"	4	"attack"
// "pfqallllabwaoclk"	2	"fallback"
// 입출력 예 설명
// 입출력 예 #1

// "dfjardstddetckdaccccdegk" 의 4번째, 8번째, 12번째, 16번째, 20번째, 24번째 글자를 합친 "attack"을 return합니다.
// 입출력 예 #2

// "pfqallllabwaoclk" 의 2번째, 4번째, 6번째, 8번째, 10번째, 12번째, 14번째, 16번째 글자를 합친 "fallback"을 return합니다.

package main

func solution(cipher string, code int) string {
	answer := ""
	for i := code - 1; i < len(cipher); i = i + (code) {
		answer += string(cipher[i])
	}

	return answer
}

func main() {
	cipher := "dfjardstddetckdaccccdegk"
	code := 4
	solution(cipher, code)
}
