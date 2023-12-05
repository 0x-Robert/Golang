// 숫자 짝꿍
// 문제 설명
// 두 정수 X, Y의 임의의 자리에서 공통으로 나타나는 정수 k(0 ≤ k ≤ 9)들을 이용하여 만들 수 있는 가장 큰 정수를 두 수의 짝꿍이라 합니다(단, 공통으로 나타나는 정수 중 서로 짝지을 수 있는 숫자만 사용합니다). X, Y의 짝꿍이 존재하지 않으면, 짝꿍은 -1입니다. X, Y의 짝꿍이 0으로만 구성되어 있다면, 짝꿍은 0입니다.

// 예를 들어, X = 3403이고 Y = 13203이라면, X와 Y의 짝꿍은 X와 Y에서 공통으로 나타나는 3, 0, 3으로 만들 수 있는 가장 큰 정수인 330입니다. 다른 예시로 X = 5525이고 Y = 1255이면 X와 Y의 짝꿍은 X와 Y에서 공통으로 나타나는 2, 5, 5로 만들 수 있는 가장 큰 정수인 552입니다(X에는 5가 3개, Y에는 5가 2개 나타나므로 남는 5 한 개는 짝 지을 수 없습니다.)
// 두 정수 X, Y가 주어졌을 때, X, Y의 짝꿍을 return하는 solution 함수를 완성해주세요.

// 제한사항
// 3 ≤ X, Y의 길이(자릿수) ≤ 3,000,000입니다.
// X, Y는 0으로 시작하지 않습니다.
// X, Y의 짝꿍은 상당히 큰 정수일 수 있으므로, 문자열로 반환합니다.
// 입출력 예
// X	Y	result
// "100"	"2345"	"-1"
// "100"	"203045"	"0"
// "100"	"123450"	"10"
// "12321"	"42531"	"321"
// "5525"	"1255"	"552"
// 입출력 예 설명
// 입출력 예 #1

// X, Y의 짝꿍은 존재하지 않습니다. 따라서 "-1"을 return합니다.
// 입출력 예 #2

// X, Y의 공통된 숫자는 0으로만 구성되어 있기 때문에, 두 수의 짝꿍은 정수 0입니다. 따라서 "0"을 return합니다.
// 입출력 예 #3

// X, Y의 짝꿍은 10이므로, "10"을 return합니다.
// 입출력 예 #4

// X, Y의 짝꿍은 321입니다. 따라서 "321"을 return합니다.
// 입출력 예 #5

// 지문에 설명된 예시와 같습니다.

package main

import (
	"fmt"
	"strings"
)

// 두 문자열이 일치하는지 체크
// 문자열이 일치하는 것 중 0으로만 구성되면 답은 0, 아니면 그냥 숫자로 슬라이스에 저장
// 이 중 내림차순 순서대로 정렬후 다시 전체 값을 문자열로 통합 후 리턴
// 일치하는게 없으면 -1 리턴, 아니면 가장 큰 수 리턴
// func solution(X string, Y string) string {
// 	arr := []int{}
// 	answer := ""

// 	for _, v := range X {
// 		if strings.Contains(Y, string(v)) {
// 			v2, _ := strconv.Atoi(string(v))
// 			//&& allkey[v] != 1
// 			//allkey[v] += 1
// 			arr = append(arr, v2)

// 		}
// 	}
// 	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
// 	fmt.Println("arr", arr)

// 	for _, v := range arr {
// 		answer += strconv.Itoa(v)
// 	}
// 	fmt.Println("answer", answer)

// 	if len(answer) == 0 {
// 		return "-1"
// 	} else {
// 		return answer
// 	}

// 	return ""
// }

func solution(X string, Y string) string {
	var sb strings.Builder

	xArr := make([]int, 10) // X 숫자별 개수 파악용 배열
	yArr := make([]int, 10) // Y 숫자별 개수 파악용 배열

	fmt.Println(xArr, yArr)
	// 숫자별로 배열 설정
	for _, x := range X {
		fmt.Println("x", x, x-'0', int(x-'0'))
		// 얼마나 쓰였는지 중복체크
		xArr[int(x-'0')]++
		fmt.Println(xArr)
	}
	for _, y := range Y {
		// 얼마나 쓰였는지 중복체크
		yArr[int(y-'0')]++
		fmt.Println(yArr)
	}
	fmt.Println("sb1", sb)
	// 마지막(9)부터 반복하면서 둘 다 있는 숫자면 sb에 추가한다.
	for i := 9; i >= 0; i-- {
		for xArr[i] > 0 && yArr[i] > 0 {
			// sb에 값을 축적
			sb.WriteRune(rune(i + '0'))

			xArr[i]--
			yArr[i]--
		}
	}
	fmt.Println("sb2", sb)
	result := sb.String()
	fmt.Println("res", result)
	if result == "" {
		return "-1"
	}
	// 9부터 역순으로 조회했으니 첫번재 문자열 값이 0이면 전체도 0
	if result[0] == '0' {
		return "0"
	}

	return result
}

func main() {
	// X := "100"
	// X := "12321"
	X := "5525"
	// Y := "2345"
	// Y := "203045"
	//	Y := "123450"
	// Y := "42531"
	Y := "1255"
	solution(X, Y)
}
