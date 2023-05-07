package main

import (
	"unsafe"
)

type User struct {
	FirstName string
	LastName  string //string 구조체는 포인터필드와 길이 필드 2개를 가지고 있어서 16바이트를 갖습니다. 따라서 string 2개와 8바이트 1개로 총 40바이트 크기를 갖습니다.
	Age       int    //8바이트
}

func main() {

	userSize := unsafe.Sizeof(User{})
	println(userSize)
}

// 위 코드에서 unsafe.Sizeof(User{})는 User 구조체의 인스턴스를 생성하고,
// 해당 인스턴스의 크기를 unsafe.Sizeof 함수로 계산하여 출력합니다.
// 이때 출력값은 48이 됩니다. 이는 string 타입의 LastName 필드가 16바이트를 차지하고,
// int 타입의 Age 필드가 8바이트를 차지하기 때문입니다.
// 따라서 FirstName 필드가 16바이트를 차지하므로,
// User 구조체 전체의 크기는 16 + 16 + 8 = 40바이트가 됩니다.
// 하지만 구조체 필드들은 구조체 내부에서 정렬되어 저장되므로, 이에 따라 추가적인 패딩이 발생하여 최종적으로 48바이트의 크기를 가지게 됩니다.
