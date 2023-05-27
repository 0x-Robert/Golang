// ch13/ex13.1/ex13.1.go
package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	//구조체 선언
	//구조체를 초기화할 경우 값은 다음과 같다.
	// string ""
	// int 0
	// float: 0.0
	// Type : 0.0
	var house House
	//주소설정
	house.Address = "인천시 서구"
	house.Size = 28
	house.Price = 2.8
	house.Type = "아파트"

	//다음과 같은 방법도 가능함
	//var house House = House{"인천시", 28, 9.80, "아파트"}

	//다음과 같은 방법도 가능함
	/*
		var house House = House{
			"인천시",
			28,
			8.80,
			"아파트", 여러 줄로 초기화할 때 마지막 값에 쉼표가 필요함
		}

	*/

	fmt.Println("주소:", house.Address)
	fmt.Printf("크기: %d평\n", house.Size)
	fmt.Printf("가격: %.2f억원\n", house.Price)
	fmt.Println("타입:", house.Type)

}
