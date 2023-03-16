package main

import (
	"fmt"
	"strings"
)

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
