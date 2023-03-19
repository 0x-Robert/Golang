package main

import "fmt"

func main() {
	//len이 3 cap이 5인 스랄이스를 만든다.
	//배열 요소는 3개고 실제 저장 가능한 공간은 5개라는 뜻으로
	//빈 공간은 2개이다.
	//따라서 빈 공간에 있는 slice1의 공간에 4,5를 넣은 slice2는 slice1과 같은 배열을 가리킨다.
	//왜냐하면 슬라이스는 다음의 구조를 가진다.

	/*
		type SliceHeader struct{
			Data uintptr //실제 배열을 가리키는 포인터
			Len int // 요소 개수
			Cap int //실제 배열길이
		}

	*/
	slice1 := make([]int, 3, 5)

	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100
	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)
	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
