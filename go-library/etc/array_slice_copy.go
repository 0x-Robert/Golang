package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	b = array //1,2,3,4,5, 8바이트가 5개라서 40바이트만큼 복사됨 , array 배열의 복사본을 가짐 

	var c []int //슬라이스 타입이다.  

	c = array[:] //1,2,3,500,5 타입이 다르기 때문에 array전체를 슬라이싱한 값으로 대입해야한다. 

	b[0] = 1000 //1000, 2,3,4,5
	c[3] = 500  //1,2,3,500,5

	fmt.Println("array:", array)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
