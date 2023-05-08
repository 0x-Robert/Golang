// ch15/ex15.17/ex15.17.go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//문자열 합산을 하면 기존 문자열 메모리 공간을 건드리지 않고 새로운 메모리 공간을 만들어서 두 문자열을 합치기 때문에
	//string 합 연산 이후 주솟값이 변경된다.
	//따라서 문자열 합산이 많을 경우 string 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있다.
	var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // ❶
	addr1 := stringheader.Data                                    // ❷

	str += " World"            // ❸
	addr2 := stringheader.Data // ➍

	str += " Welcome!"         // ➎
	addr3 := stringheader.Data // ➏

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}
