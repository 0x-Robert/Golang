package main

import "fmt"

type SliceHeader struct {
	Data uintptr //실제 배열을 가리키는 포인터
	Len  int     //요소 개수
	Cap  int     //실제 배열 길이
}

func main() {

	var array [10]int
	var slice []int = array[1:3] //slice의 data 필드는 array의 두 번째 요소 메모리 주소를 가진다. 
	//Len 필드값은 끝 인덱스 - 처음 인덱스 
	//slice의 데이터는 rray의 두 번째 요소주소이고, Len은 2, Cap은 9다. 
	fmt.Println("array", array)
	fmt.Println("slice", slice)
}
