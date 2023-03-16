package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//문자열은 불변이다.
//문자열과 슬라이스는 다른 메모리값을 가지고 있다.
//go언어는 슬라이스로 타입 변환을 할 때 문자열을 복사해서 새로운 메모리 공간을 만들어 슬라이스가 가리키도록 한다.

// reflect 패키지와 unsafe 패키지는 Go 언어에서 메모리에 직접 접근할 수 있는 기능을 제공합니다. 이를 통해 우리는 문자열과 슬라이스가 저장되는 메모리 주소를 확인할 수 있습니다.

// 첫 번째 줄에서는 문자열 변수 str과 slice 변수를 선언합니다. 그리고 문자열 str을 byte 슬라이스로 변환하여 slice 변수에 할당합니다.

// 두 번째 줄에서는 문자열과 슬라이스의 메모리 주소를 확인하기 위해 reflect 패키지에서 제공하는 StringHeader와 SliceHeader를 사용합니다.

// StringHeader는 문자열의 내부 구조체로, 문자열이 저장된 메모리 주소와 길이 정보를 담고 있습니다.

// SliceHeader는 슬라이스의 내부 구조체로, 슬라이스가 가리키는 배열의 메모리 주소와 길이 정보를 담고 있습니다.

// unsafe 패키지에서 제공하는 Pointer 함수는 인터페이스 값을 원시 포인터 값으로 변환할 수 있습니다. 이 함수를 사용하여 문자열과 슬라이스의 메모리 주소를 가져옵니다.

// ❷에서는 문자열과 슬라이스의 메모리 주소를 출력합니다. 문자열은 불변이므로 메모리 주소는 변하지 않지만, 슬라이스는 가리키는 배열이 변경될 수 있으므로 메모리 주소도 변할 수 있습니다.
func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // ❶
	sliceheader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringheader.Data) // ❷
	fmt.Printf("slice:\t%x\n", sliceheader.Data)
}
