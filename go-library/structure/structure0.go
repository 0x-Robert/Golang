package main

import (
	"fmt"
	"unsafe"
)

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

type User0 struct {
	Age   int32   //4
	Score float64 //8
}

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type User2 struct {
	A int8 //1
	C int8 //1
	E int8 //1
	B int  //8
	D int  //8
	//메모리 패딩을 막으려면 작은 것을 앞에 선언
}

type User3 struct {
	A int8 //1
	B int  //8
	C int8 //1
	D int  //8
	E int8 //1
	//메모리 낭비
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	var house House
	house.Address = "incheon"
	house.Size = 30
	house.Price = 100
	house.Category = "아파트"
	//%v쓰면 {} 사용가능
	//fmt.Printf("%v\n", house)

	fmt.Printf("주소:%s 사이즈:%d평 가격:%v억원 종류:%s\n",
		house.Address, house.Size, house.Price, house.Category)

	// var house House = House{"seoul", 30, 3.2, "빌라"}
	// fmt.Println(house)
	user := User2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))

	user2 := User3{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user2))

	user0 := User0{23, 77.2}
	// 12여야하는데 왜 16인가?
	// 이유는 메모리 정렬 때문에, 메모리 패딩 문제 때문이다.

	fmt.Println(unsafe.Sizeof(user0))

	//구조체의 역할
	//결합도(의존성)은 낮게 응집도는 높게
	//low coupling, high cohesion
	//객체지향 프로그래밍의 기반이 됨 >> 구조체

}
