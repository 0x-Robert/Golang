package main

import "fmt"

type stringHeader struct {
	Data uintptr //Data필드는 uintptr 타입으로 문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
	Len  int     //Len은 int 타입으로 문자열의 길이를 나타냄
}

func main() {

	str1 := "안녕하세요 한글 문자열입니다."
	str2 := str1

	//str1의 Data와 Len 값만 str2에 복사한다. 
	fmt.Printf(str1)
	fmt.Printf("\n")
	fmt.Printf(str2)
}
