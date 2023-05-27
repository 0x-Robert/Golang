package main

import "fmt"

func CallbyCopy(n int, b [5]int, s []int) {
	n = 3000
	b[0] = 1000
	s[3] = 500
}

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5} //배열
	var c []int                              //슬라이스 타입
	c = array[:]                             //배열전체를 슬라이싱, 1,2,3,4,5
	CallbyCopy(100, array, c)                //array = 1,2,3,500,5 / c = 1,2,3,500,5

	fmt.Println("array:", array)
	fmt.Println("c:", c)
}
