package main

import "fmt"
func main(){

	slice := make([]int, 3) // []int 슬라이스, len이 3
	fmt.Println(slice) // [0 0 0]

	slice2 := make([]int, 3, 5) // []int 슬라이스 len이 3, cap이 5
	fmt.Println(slice2) // [0 0 0] 이라 출력은 되는데 나머지 2개는 추가될 요소를 위해 비워둠

	slice3 := []int{1,2,3} // len 3, cap 3
	slice4 := append(slice3, 4, 5) 

	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice3[1] = 100;
	fmt.Println(slice3)
	fmt.Println(slice4)
} 