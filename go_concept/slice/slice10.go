package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)

	//copy(복사한 결과를 저장하는 슬라이스 변수, 복사 대상이 되는 슬라이스 변수)
	//반환값은 실제로 복사한 요소 개수
	//slice2 := make([]int, len(slice1))
	//copy(slice2, slice1)
	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
}
