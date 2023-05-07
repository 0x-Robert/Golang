package main

import "fmt"

func main() {

	var a int
	//포인터 변수 선언
	//int 타입 변수를 가리키는 포인터 변수
	var p *int
	//a의 메모리 주소를 포인터 변수 p에 대입한다.
	p = &a
	fmt.Println("p1", p)
	fmt.Println("a1", a)

	*p = 20
	fmt.Println("p2", p)
	fmt.Println("a2", a)
}
