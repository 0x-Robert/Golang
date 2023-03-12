// ch13/ex13.8/ex13.8.go
package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1바이트
	C int8 // 1바이트
	E int8 // 1바이트
	B int  // 8바이트
	D int  // 8바이트
}

func main() {
	//메모리 낭비와 패딩을 줄이기 위해서는
	//8바이트보다 작은 필드는 8바이트 크기(단위)를 고려해서 몰아서 배치하자.
	//프로그래밍 역사는 객체간 결합도는 낮추고 연관있는데이터간 응집도를 올리는 방향으로 흘러왔다. 
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))

}
