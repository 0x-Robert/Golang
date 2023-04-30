package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 //1바이트
	B int  //8바이트
	C int8 //1바이트
	D int  //8바이트
	E int8 //1바이트
}

func main() {

	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
	//구조체크기가 19바이트가 될 것 같지만 실제로는 메모리 패딩때문에 구조체 크기는 40바이트가 된다.

}
