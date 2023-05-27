package main

import "fmt"

func main() {
	var a int = 500
	var p *int //int 타입 포인터 변수 p 선언

	p = &a // a의 메모리 주소를 변수 p의 값으로 대입(복사)

	fmt.Printf("p: %p\n", p)               //memory address 출력
	fmt.Printf("p memory value: %d\n", *p) //500, p가 가리키는 메모리의 값 출력

	*p = 100                        // p가 가리키는 메모리 공간의 값을 변경한다.
	fmt.Printf("a  value: %d\n", a) //100, 당연히 a의 메모리 공간 값이 변했으므로 a는 100
}
