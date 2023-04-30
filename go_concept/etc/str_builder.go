package main

import (
	"fmt"
	"strings"
)

// package main: Go에서 이 파일이 실행 가능한 프로그램임을 나타내기 위한 키워드입니다.

// import "fmt": "fmt" 패키지를 가져옵니다. "fmt"는 표준 입출력 함수를 제공합니다.

// import "strings": "strings" 패키지를 가져옵니다. "strings"는 문자열 조작 함수를 제공합니다.

// func ToUpper1(str string) string {...}: 문자열을 인수로 받아 대문자로 변환된 새로운 문자열을 반환하는 함수입니다. 이 함수는 문자열의 모든 문자를 하나씩 검사하며,
//소문자인 경우 대문자로 변환합니다. 새로운 문자열을 만들기 위해 문자열 연산자인 +를 사용합니다.

// func ToUpper2(str string) string {...}: ToUpper1과 같은 기능을 수행하지만, 문자열을 조작하기 위해 strings 패키지의 strings.Builder 타입을 사용합니다.
//이 타입은 문자열 연산이 많은 상황에서 성능을 개선하기 위해 설계되었습니다. ToUpper1과 다르게, 문자열을 직접 만드는 대신, Builder 타입의 객체를 만들고, WriteRune 메서드를 사용하여 문자를 추가합니다.

// func main() {...}: 프로그램의 진입점입니다. main 함수에서는 문자열을 만들고, ToUpper1 함수와 ToUpper2 함수를 호출하여
// 문자열의 모든 소문자를 대문자로 변환한 결과를 출력합니다. fmt.Println 함수를 사용하여 결과를 표시합니다.
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			//strings.Builder는 내부에 슬라이스를 가지고 있어서
			//WriteRune을 통해 문자를 더할 때 매번 메모리를 새로 생성하지 않고 기존 메모리 공간에 빈자리가 있으면 그냥 더한다.
			//그래서 메모리 공간 낭비를 없앨 수 있다.
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello Yongari"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
