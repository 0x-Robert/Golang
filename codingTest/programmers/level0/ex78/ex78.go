// 대문자와 소문자
// 문제 설명
// 문자열 my_string이 매개변수로 주어질 때, 대문자는 소문자로 소문자는 대문자로 변환한 문자열을 return하도록 solution 함수를 완성해주세요.

// 제한사항
// 1 ≤ my_string의 길이 ≤ 1,000
// my_string은 영어 대문자와 소문자로만 구성되어 있습니다.
// 입출력 예
// my_string	result
// "cccCCC"	"CCCccc"
// "abCdEfghIJ"	"ABcDeFGHij"
// 입출력 예 설명
// 입출력 예 #1

// 소문자는 대문자로 대문자는 소문자로 바꾼 "CCCccc"를 return합니다.
// 입출력 예 #2

// 소문자는 대문자로 대문자는 소문자로 바꾼 "ABcDeFGHij"를 return합니다.
package main

import (
	"strings"
)

func solution(my_string string) string {
	answer := ""
	// 아스키 코드에서 65번부터 90번까지가 대문자 A~Z와 같다.
	for _, v := range my_string {
		if v >= 65 && v <= 90 {
			answer += strings.ToLower(string(v))
			// 아스키 코드에서 97번부터 122번까지가 소문자 a~z와 같다.
		} else if v >= 97 && v <= 122 {
			answer += strings.ToUpper(string(v))
		}
	}

	return answer
}

func main() {
	m := "cccCCC"
	solution(m)
}
