package main

import "fmt"

func main() {
	//str[i]와 같이 문자열을 인덱스를 이용해 출력할 경우 바이트 단위로 출력되므로
	//[]rune으로 타입 변환 후 실행

	str := "Hello 월드!" //한 영이 섞인 문자열
	arr := []rune(str) //[]rune으로 타입 변환가능 그러나 변환되는 과정에서 별도의 배열을 할당하므로 불필요한 메모리를 사용함, 따라서 range를 사용하면 이런 불필요함을 방지할 수 있음

	for i := 0; i < len(arr); i++ {
		//데이터타입, UTF8코드, 문자값
		//byte와 uint8은 이름만 다르고 같은 별칭
		//rune과 int32타입과 같음
		fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
	}

}
