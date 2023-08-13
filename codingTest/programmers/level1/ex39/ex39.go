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
