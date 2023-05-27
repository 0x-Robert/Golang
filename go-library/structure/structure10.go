package main

import (
	"fmt"
	"unsafe"
)

// 구조체의 패딩을 최대한 줄이고 구조체 크기를 적기
type PaddingA struct {
	A int8    //1바이트
	B int     //8바이트
	C float64 //8바이트
	D uint16  //2바이트
	E int     //8바이트
	F float32 //4바이트
	G int8    //1바이트
}

// 8바이트보다 작은 메모리단위를 몰아주자. 따라서 다음과 같이 구조체를 재정의했다.
type PaddingB struct {
	A int8
	G int8
	D uint16
	F float32
	B int
	C float64
	E int
}

func main() {
	paddingA := PaddingA{1, 2, 1.1, 2, 3, 3.3, 4}
	fmt.Println(unsafe.Sizeof(paddingA)) //구조체 A의 크기
	fmt.Println(unsafe.Sizeof(paddingA.A))
	fmt.Println(unsafe.Sizeof(paddingA.B))
	fmt.Println(unsafe.Sizeof(paddingA.C))
	fmt.Println(unsafe.Sizeof(paddingA.D))
	fmt.Println(unsafe.Sizeof(paddingA.E))
	fmt.Println(unsafe.Sizeof(paddingA.F))
	fmt.Println(unsafe.Sizeof(paddingA.G))

	paddingB := PaddingB{1, 4, 2, 3.3, 2, 1.1, 3}
	fmt.Println(unsafe.Sizeof(paddingB)) //구조체 B의 크기
	fmt.Println(unsafe.Sizeof(paddingB.A))
	fmt.Println(unsafe.Sizeof(paddingB.G))
	fmt.Println(unsafe.Sizeof(paddingB.D))
	fmt.Println(unsafe.Sizeof(paddingB.F))
	fmt.Println(unsafe.Sizeof(paddingB.B))
	fmt.Println(unsafe.Sizeof(paddingB.C))
	fmt.Println(unsafe.Sizeof(paddingB.E))

}
