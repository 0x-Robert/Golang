package main

import "fmt"

func main() {
	str := "Hello 월드"    //한글과 영문자가 섞인 문자열
	runes := []rune(str) //[]rune 타입으로 타입 변환
	//영문자, 공백 1바이트, 한글 3바이트
	// "Hello "= 6바이트 / 월드=6바이트 >> 12바이트
	fmt.Printf("len(str) = %d\n", len(str)) //str 타입 길이
	//배열 요소 8개
	fmt.Printf("len(runes) = %d\n", len(runes)) //rune 타입 길이

	//문자열 요소 그대로, string타입은 연속된 바이트 메모리 
	fmt.Println(str) 
	//문자열을 코드로 변경했음 []rune 타입은 글자들의 배열로 이루어져있다. 
	fmt.Println(runes)

	//네트워크로 데이터를 전송하는 경우 io.Writer 인터페이스를 사용하고 io.Writer 인터페이스는 []byte 인수를 받기 때문에 []byte 인수로 반환해야한다. 
	//그래서 문자열을 쉽게 전송하고자 string에서 []byte타입으로 변환을 지원한다. 
	//string과 []byte 타입간 상호변환이 가능하다. 

	
}
