// 한 번만 등장한 문자
// 문제 설명
// 문자열 s가 매개변수로 주어집니다. s에서 한 번만 등장하는 문자를 사전 순으로 정렬한 문자열을 return 하도록 solution 함수를 완성해보세요. 한 번만 등장하는 문자가 없을 경우 빈 문자열을 return 합니다.

// 제한사항
// 0 < s의 길이 < 1,000
// s는 소문자로만 이루어져 있습니다.
// 입출력 예
// s	result
// "abcabcadc"	"d"
// "abdc"	"abcd"
// "hello"	"eho"
// 입출력 예 설명
// 입출력 예 #1

// "abcabcadc"에서 하나만 등장하는 문자는 "d"입니다.
// 입출력 예 #2

// "abdc"에서 모든 문자가 한 번씩 등장하므로 사전 순으로 정렬한 "abcd"를 return 합니다.
// 입출력 예 #3

// "hello"에서 한 번씩 등장한 문자는 "heo"이고 이를 사전 순으로 정렬한 "eho"를 return 합니다.

package main

import (
	"sort"
	"strings"
)

func solution(s string) string {
	mp := make(map[string]int)
	answer := ""

	for _, v := range s {
		mp[string(v)] += 1
	}

	for k, v := range mp {
		if v == 1 {
			answer += k
		}
	}

	answer2 := strings.Split(answer, "")
	sort.Strings(answer2)
	answer3 := strings.Join(answer2, "")

	return answer3
}

func solution1(my_string string) string {
	m := map[string]int{}
	for _, s := range my_string {
		m[string(s)]++
	}

	keys := make([]string, 0, len(m))
	for k, v := range m {
		if v == 1 {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	return strings.Join(keys, "")
}

func main() {
	// s := "abcabcadc"
	s := "abdc"
	// s := "abcabcadc"
	solution(s)
}
