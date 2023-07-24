// 옹알이 (1)
// 문제 설명
// 머쓱이는 태어난 지 6개월 된 조카를 돌보고 있습니다. 조카는 아직 "aya", "ye", "woo", "ma" 네 가지 발음을 최대 한 번씩 사용해 조합한(이어 붙인) 발음밖에 하지 못합니다. 문자열 배열 babbling이 매개변수로 주어질 때, 머쓱이의 조카가 발음할 수 있는 단어의 개수를 return하도록 solution 함수를 완성해주세요.

// 제한사항
// 1 ≤ babbling의 길이 ≤ 100
// 1 ≤ babbling[i]의 길이 ≤ 15
// babbling의 각 문자열에서 "aya", "ye", "woo", "ma"는 각각 최대 한 번씩만 등장합니다.
// 즉, 각 문자열의 가능한 모든 부분 문자열 중에서 "aya", "ye", "woo", "ma"가 한 번씩만 등장합니다.
// 문자열은 알파벳 소문자로만 이루어져 있습니다.
// 입출력 예
// babbling	result
// ["aya", "yee", "u", "maa", "wyeoo"]	1
// ["ayaye", "uuuma", "ye", "yemawoo", "ayaa"]	3
// 입출력 예 설명
// 입출력 예 #1

// ["aya", "yee", "u", "maa", "wyeoo"]에서 발음할 수 있는 것은 "aya"뿐입니다. 따라서 1을 return합니다.
// 입출력 예 #2

// ["ayaye", "uuuma", "ye", "yemawoo", "ayaa"]에서 발음할 수 있는 것은 "aya" + "ye" = "ayaye", "ye", "ye" + "ma" + "woo" = "yemawoo"로 3개입니다. 따라서 3을 return합니다.
// 유의사항
// 네 가지를 붙여 만들 수 있는 발음 이외에는 어떤 발음도 할 수 없는 것으로 규정합니다. 예를 들어 "woowo"는 "woo"는 발음할 수 있지만 "wo"를 발음할 수 없기 때문에 할 수 없는 발음입니다.
// ※ 공지 - 2022년 10월 27일 문제 지문이 리뉴얼되었습니다. 기존에 제출한 코드가 통과하지 못할 수도 있습니다.

package main

import (
	"regexp"
	"strings"
)

func solution(babbling []string) int {
	a := "aya"
	b := "ye"
	c := "woo"
	d := "ma"
	result := ""
	answer := 0

	for _, v := range babbling {

		result = v
		if strings.Contains(result, a) == true {
			result = strings.Replace(result, a, "1", -1)
		}
		if strings.Contains(result, b) == true {
			result = strings.Replace(result, b, "1", -1)
		}
		if strings.Contains(result, c) == true {
			result = strings.Replace(result, c, "1", -1)
		}

		if strings.Contains(result, d) == true {
			result = strings.Replace(result, d, "1", -1)
		}

		// 처음에 각각 a,b,c,d 문자열에 해당하는 값이 있다면 공백으로 치환한 뒤 문자열 길이가 0일 때만 answer에 1을 더하는 생각만 했었다.
		// 그러나 두 번째 테스트케이스는 통과하지만 첫 번째 테스트케이스가 통과하지 않았다. 왜냐하면 wyeoo 문자열은 내가 생각한 대로 하면 전부 공백이 되지만 문제 조건에 따르면 공백이 되면 안되기 때문이다.
		// 이후 숫자로 변환해볼까 생각은 해봤으나 그다음 단계를 생각하지 못했다.
		// 이후 구글링을 통해 1로 바꾼 걸 전부 공백으로 바꾸면 된다는 생각을 보고 다음 줄을 넣으면 해결이 되겠다는 생각을 했다.
		result = strings.Replace(result, "1", "", -1)

		if len(result) == 0 {
			answer += 1
		}

	}
	return answer
}

// 정규표현식 풀이... 간단히 푸는 방법이 있었다..
func solution1(babbling []string) (result int) {
	regex := regexp.MustCompile("aya|ye|woo|ma")
	for _, word := range babbling {
		if len(regex.ReplaceAllString(word, "")) == 0 {
			result++
		}
	}
	return
}

func main() {
	babbling := []string{"aya", "yee", "u", "maa", "wyeoo"}
	// babbling := []string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"}
	solution(babbling)
}
