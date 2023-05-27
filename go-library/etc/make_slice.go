package main

import (
	"fmt"
)

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice1 []int = array[1:5]    //배열 슬라이싱
	var slice2 []int = slice1[1:8:9] //슬라이스 슬라이싱 slice1ㅣ 가리키는 배열의 일부를 다시 가리킴
	//slice1은 rray의 두 번째 요소부터 가리키고 있으므로 slice2는 slice1의 두 번재 요소인 rray 세 번째 요소를 가진다.
	//그리고 Len은 7(8-1), Cap은 8(9-1)이 된다.
	//슬라이싱 기준은 슬라이스가 가리키는 원래 배열이된다.

	var slice3 []int = make([]int, 5)       //make() 길이 5짜리 배열을 만들고 slice3이 가리킨다
	var slice4 []int = make([]int, 0)       //길이 0인 슬라이스 메모리를 차지하지않고 데이터저장도 할 수 없으나 슬라이스 길이를 특정할 수 없을 때 사용, 슬라이스는 nil이 아니고 길이가 0인 배열을 가리킴
	var slice5 []int = []int{1, 2, 3, 4, 5} //초기화, 슬라이스를 특정 값을 가진 배열로 초기화
	var slice6 []int                        //초기화하지 않고 슬라이스를 만들면 nil인 슬라이스를 만드나.

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2, cap(slice2))
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)

	if slice4 != nil {
		fmt.Println("slice4 is not nil")
	}
	fmt.Println("slice5", slice5)
	fmt.Println("slice6", slice6)
	if slice6 == nil {
		fmt.Println("slice6 is nil")
	}
}
