package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// type stringHeader struct {
// 	Data uintptr //Data필드는 uintptr 타입으로 문자열의 데이터가 있는 메모리 주소를 나타내는 일종의 포인터
// 	Len  int     //Len은 int 타입으로 문자열의 길이를 나타냄
// }

func main() {

	//실제 Data와 Len값이 복사됐는지 확인
	str1 := "Hello World!"
	str2 := str1 //str1의 모든 필드 값이 str2에 복사되서 같은 메모리 데이터를 가리킴

	//같은 구조체라서 데이터만 복사
	//go는 String타입에서 *reflect.StringHeader 타입으로 변환을 막고 있기 때문에 string 타입 변수를 *reflect.StringHeader 타입으로 강제변환을 하려고
	//unsafe.Pointer(&str1)을 사용해서 unsafe.Pointer타입으로 변환한다음에 *reflect.StringHeader 타입으로 변환함
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) //Data 값 추출, 내부필드값을 접근하고자 강제로 타입 변환한다고 생각하면 된다.
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) //Data 값 추출

	fmt.Println(stringHeader1) //두 string 변수의 data 값이 같다.
	fmt.Println(stringHeader2) //string 변수가 가리키는 문자열이 아무리 길어도 string 변수끼리 대입연산에서는 16바이트만 복사될 뿐 문자열 데이터는 복사되지 않는다. 

}
