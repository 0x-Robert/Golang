package main

import "fmt"

func main() {
	slice := []int{0, 10, 20, 30}
	fmt.Println("slice1", slice)
	newslice := slice
	newslice[0] = 100
	fmt.Println("slice2", slice)
	fmt.Println("newslice", newslice)
	fmt.Println(newslice)
	fmt.Println(slice)

	slice2 := []int{0, 10, 20, 30}

	//새로운 슬라이스 생성
	copyslice := make([]int, len(slice))
	//copy 함수로 슬라이스 복사
	copy(copyslice, slice2)

	//값변경
	copyslice[0] = 100

	fmt.Println("copyslice", copyslice)
	fmt.Println("slice2", slice2)

}
