package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	//맨 뒤에 요소 추가
	//slice = [1,2,3,4,5,6,0]
	slice = append(slice, 0)
	fmt.Println("1", slice)
	idx := 2 //추가하려는 위치

	//len(slice)==7 , i:=5 i>=2 >> i= 5, 4, 3, 2
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
		//slice[6] = slice[5]
		//slice[5] = slice[4]
		//slice[4] = slice[3]
		//slice[3] = slice[2]
	}
	fmt.Println("2", slice)
	//before slice=1,2,3,4,5,6,0
	//after slice= 1,2,3,3,4,5,6

	slice[idx] = 100
	//slice = 1,2,100,3,4,5,6

	fmt.Println(slice)

}
