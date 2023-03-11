package main

import "fmt"

func main() {
	//배열은 연속된 메모리다.
	//컴퓨터는 인덱스와 타입 크기를 사용해서 메모리 주소를 찾는다.
	var a0 [10]int32
	a0[3] = 300
	fmt.Println("a0", a0)

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	//배열 복사시 타입크기와 같아야 복사 가능 
	//golang에서 대입 연산자는 우변의 값을 좌변의 메모리 공간에 복사한다. 
	b = a
	

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

}
