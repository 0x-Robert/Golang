// 문자열 정렬하기 (2)
// 문제 설명
// 영어 대소문자로 이루어진 문자열 my_string이 매개변수로 주어질 때, my_string을 모두 소문자로 바꾸고 알파벳 순서대로 정렬한 문자열을 return 하도록 solution 함수를 완성해보세요.

// 제한사항
// 0 < my_string 길이 < 100
// 입출력 예
// my_string	result
// "Bcad"	"abcd"
// "heLLo"	"ehllo"
// "Python"	"hnopty"
// 입출력 예 설명
// 입출력 예 #1

// "Bcad"를 모두 소문자로 바꾸면 "bcad"이고 이를 알파벳 순으로 정렬하면 "abcd"입니다.
// 입출력 예 #2

// "heLLo"를 모두 소문자로 바꾸면 "hello"이고 이를 알파벳 순으로 정렬하면 "ehllo"입니다.

package main

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
}

func solution(my_string string) string {
	answer := ""
	for _, v := range my_string {
		if v >= 65 && v <= 96 {
			v += 32
			answer += string(v)
		} else {
			answer += string(v)
		}
	}

	answer = SortString(answer)
	return answer
}

func solution1(my_string string) string {
	bytes := []byte(strings.ToLower(my_string))
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func main() {
	my_string := "Bbcad"
	solution(my_string)
}
