package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	//위 코드는 같은 길이의 슬라이스를 만들고 순회를 사용해서 복사하기 때문에 복잡하다.
	//다음 코드를 이용해 한줄로 줄일 수 있다.
	//slice2 := append([]int{}, slice1...)

	//또는 내장함수 copy()를 사용하는 방법도 있다.
	
	slice1[1] = 100
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}
