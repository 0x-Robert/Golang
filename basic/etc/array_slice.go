package main

import "fmt"

type SliceHeader struct {
	Data uintptr //실제 배열을 가리키는 포인터
	Len  int     //요소 개수
	Cap  int     //실제 배열 길이
}

func main() {

	var array [10]int
	var slice []int = array[1:3]
	fmt.Println("array", array)
	fmt.Println("slice", slice)
}
