// 외계행성의 나이
// 문제 설명
// 우주여행을 하던 머쓱이는 엔진 고장으로 PROGRAMMERS-962 행성에 불시착하게 됐습니다. 입국심사에서 나이를 말해야 하는데, PROGRAMMERS-962 행성에서는 나이를 알파벳으로 말하고 있습니다. a는 0, b는 1, c는 2, ..., j는 9입니다. 예를 들어 23살은 cd, 51살은 fb로 표현합니다. 나이 age가 매개변수로 주어질 때 PROGRAMMER-962식 나이를 return하도록 solution 함수를 완성해주세요.

// 제한사항
// age는 자연수입니다.
// age ≤ 1,000
// PROGRAMMERS-962 행성은 알파벳 소문자만 사용합니다.
// 입출력 예
// age	result
// 23	"cd"
// 51	"fb"
// 100	"baa"
// 입출력 예 설명
// 입출력 예 #1

// age가 23이므로 "cd"를 return합니다.
// 입출력 예 #2

// age가 51이므로 "fb"를 return합니다.
// 입출력 예 #3

// age가 100이므로 "baa"를 return합니다.
// solution.go
// 1
// func solution(age int) string {
// 2
//
//	return ""
//
// 3
// }
// 실행 결과
// 실행 결과가 여기에 표시됩니다.

package main

import (
	"fmt"
	"strconv"
)

func solution(age int) string {
	// a=0, b=1, c=2, d=3, e=4, f=5, g=6, h=7, i=8, j=9
	// "1" = 49, "2" = 50, "3" = 51
	// ageMap := make(map[string]string)
	ageMap := map[string]string{
		"0": "a",
		"1": "b",
		"2": "c",
		"3": "d",
		"4": "e",
		"5": "f",
		"6": "g",
		"7": "h",
		"8": "i",
		"9": "j",
	}
	answer := ""
	age2 := strconv.Itoa(age)
	for _, v := range age2 {
		answer += ageMap[string(v)]
	}
	// fmt.Println("answer", answer)
	return answer
}

func solution0(age int) string {
	res := ""
	fmt.Println("strconv.Itoa(age)", strconv.Itoa(age))

	for _, v := range strconv.Itoa(age) {
		res += string(v + 49)
		fmt.Println("v", v)
		fmt.Println("v+49", v+49)
		fmt.Println("string(v + 49)", string(v+49))

	}
	return res
}

func main() {
	// age := 123
	// age := 51
	age := 100
	solution0(age)
}
