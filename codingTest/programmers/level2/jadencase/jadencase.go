// JadenCase 문자열 만들기
// 문제 설명
// JadenCase란 모든 단어의 첫 문자가 대문자이고, 그 외의 알파벳은 소문자인 문자열입니다. 단, 첫 문자가 알파벳이 아닐 때에는 이어지는 알파벳은 소문자로 쓰면 됩니다. (첫 번째 입출력 예 참고)
// 문자열 s가 주어졌을 때, s를 JadenCase로 바꾼 문자열을 리턴하는 함수, solution을 완성해주세요.

// 제한 조건
// s는 길이 1 이상 200 이하인 문자열입니다.
// s는 알파벳과 숫자, 공백문자(" ")로 이루어져 있습니다.
// 숫자는 단어의 첫 문자로만 나옵니다.
// 숫자로만 이루어진 단어는 없습니다.
// 공백문자가 연속해서 나올 수 있습니다.
// 입출력 예
// s	return
// "3people unFollowed me"	"3people Unfollowed Me"
// "for the last week"	"For The Last Week"

// 다음 문서와 코드 참고..
// https://pkg.go.dev/strings#Title
// func
// Title
// DEPRECATED
// func Title(s string) string
// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.

// Deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// Compare this example to the ToTitle example.
// 	fmt.Println(strings.Title("her royal highness"))
// 	fmt.Println(strings.Title("loud noises"))
// 	fmt.Println(strings.Title("хлеб"))
// }

package main

import (
	"fmt"
	"strings"
)

func solution0(s string) string {
	// upper가 true일 때 대문자, false일 때는 소문자화
	// 문자 첫번째가 숫자일 경우, 나머지 문자는 공백이 나올 때까지 소문자로 변환
	upcheck := 0
	answer := ""

	for i, v := range s {
		if i == 0 && v != 32 {
			//&& v >= 65 && v <= 90
			answer = strings.ToUpper(string(v))
			upcheck = 0
		}
		// fmt.Println("v, answer, upper", v, answer, upcheck)
		// 앞이 숫자일 때
		if v >= 48 && v <= 57 {
			answer += string(v)
			upcheck = 0
		} else if upcheck == 0 && v != 32 && i != 0 {
			answer += strings.ToLower(string(v))
		} else if upcheck == 1 && v != 32 {
			answer += strings.ToUpper(string(v))
			upcheck = 0
		} else if v == 32 {
			answer += string(v)
			upcheck = 1
		}
	}
	fmt.Println("answer", answer)
	return answer
}

func solution1(s string) string {
	arr := []string{}
	for _, v := range s {
		arr = append(arr, string(v))
	}

	for i, v := range arr {
		if i == 0 {
			arr[0] = strings.ToUpper(string(v))
		} else if arr[i-1] == " " {
			arr[i] = strings.ToUpper(string(v))
		} else {
			arr[i] = strings.ToLower(string(v))
		}
	}

	answer := strings.Join(arr, "")
	return answer
}

func solution2(s string) string {
	return strings.Title(strings.ToLower(s))
}

func main() {
	s := "3people unFollowed me"
	// s := "for the last week"
	// s := "  FFFor the last week"
	//    Ffffor The Last Week
	// s := "3PeoPle   UNFoLLowed   mEaAdfsaf"
	//    3people   Unfollowed   Meaadfsaf
	// 3people   Unfollowed   Me
	// 3PeoPle   unFoLLowed   mE

	solution1(s)
}
