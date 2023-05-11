package level1

import (
	"fmt"
)

/*
Go 언어에서 rune은 UTF-8 문자열에서 유니코드 코드 포인트를 나타내는 타입입니다. rune은 32비트 정수형(int32)과 같은 크기를 갖고 있습니다.

Go 언어에서 문자열은 기본적으로 UTF-8 인코딩으로 되어 있기 때문에, rune 타입을 사용하여 문자열에서 개별 문자를 처리할 수 있습니다. 문자열에서 []rune 타입으로 변환하면, 개별 문자를 인덱스로 접근할 수 있습니다.
예를 들어, 다음 코드는 string 타입인 str 변수에서 개별 문자를 추출하여 []rune으로 변환한 후, 루프를 돌면서 개별 문자를 출력합니다.
*/

// func solution(n int64) []int {

// 	arr := []int{}
// 	str := strconv.Itoa(int(n))
// 	runes := []rune(str)

// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	reverseStr2 := string(runes)
// 	for _, r := range reverseStr2 {
// 		num, _ := strconv.Atoi(string(r))
// 		arr = append(arr, num)
// 	}

// 	return arr
// }

func solution7(n int64) []int {

	var answer []int

	temp := n
	for temp > 0 {
		answer = append(answer, int(temp%10))
		fmt.Println(answer)
		temp /= 10 //남은 자릿수만큼 저장
		fmt.Println(temp)
	}
	return answer
}

func reverseArr() {
	num := int64(12345)
	solution7(num)
}
