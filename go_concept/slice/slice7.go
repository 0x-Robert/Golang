package main

import "fmt"

func main() {
	//slice1에서 빈공간이 없을 때 append를 하면
	//go언어에서는 기존 슬라이스 배열의 2배 크기로 마련한 뒤 append를 수행한다.
	//이 떄 2배로 커진 슬라이스의 메모리 주소는 달라지므로
	// slice1과 slice2의 메모리 주소는 달라진다.
	//이처럼 slice 내부에는 배열을 가리키는 포인터가 있고 빈 공간이 있을 경우 
	//append를 한다면 같은 메모리 주소공간을 가지면서 빈 공간에 추가 요소를 넣을 수 있지만
	//빈 공간이 없으면 슬라이스 크기를 2배로 만들면서 전혀 다른 슬라이스를 만들어내고 
	//메모리주소도 달라진다. 따라서 빈공간에 대해서 인식해야 하며
	//빈공간은 Cap - Len = 슬라이스 전체 크기 - 슬라이스 요소 개수로 구한다. 
	slice1 := []int{1, 2, 3}
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
