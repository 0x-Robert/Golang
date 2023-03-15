package main

import "fmt"

func main() {

	//rune타입은 4바이트 정수타입인 int32타입의 별칭 타입
	//문자 하나를 표현하는데 rune 타입을 사용한다. 
	//문자 한개를 표현하는데 작은 따옴표를 사용한다. 
	//UTF-8은 한 글자가 1~3바이트 크기 
	//go 언어 기본 타입에서 3바이트 정수타입은 제공되지 않아서 rune 타입을 사용함 
	
	var char rune = '한'

	//char의 변수 타입 프린트
	fmt.Printf("%T\n", char)
	//char 값 출력
	fmt.Println(char)
	//char 문자 출력
	fmt.Printf("%c\n", char)

}
