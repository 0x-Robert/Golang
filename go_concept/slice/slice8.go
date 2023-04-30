package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	//array를 슬라이싱 할 때 cap 길이는 array의 인덱스 1부터 배열의 마지막 인덱스까지 
	//길이를 가진다.
	//slice는 array가 요소가 5개이고 인덱스가 1부터 마지막 배열까지 사용할 수 있으므로
	//cap은 4가된다. 따라서 slice는 append를 호출해도 새로운 배열을 만들지 않고
	//기존 array와 같은 주소를 공유한다. 
	slice := array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:, ", array)
	fmt.Println("slice", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("array:, ", array)
	fmt.Println("slice", slice, len(slice), cap(slice))
}
