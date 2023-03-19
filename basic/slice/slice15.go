package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	fmt.Println("1 array", array)
	fmt.Println("1 slice", slice) ///2,3

	//2,3,100
	slice = append(slice, 100)

	//slice는 array와 같은 배열 주소를 가진다.
	//따라서 slice의 3번째 원소로 100이 들어오듯 배열에서는 4번째 요소에 100이 들어온다.

	fmt.Println("2 array", array)
	fmt.Println("2 slice", slice)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice2[:len(slice2)-2])
}
