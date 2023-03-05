package main

import "fmt"

func main() {
	var a int
	var b int

	//다음 코드가 동작하지 않는 이유
	//실제 메모리주소에 접근하는 주소값을 적어야 적용되므로
	//Scanln의 인수는 변수의 메모리 주소를 넘겨야함
	fmt.Scanln(a, b)
	fmt.Println(a, b)

	fmt.Scanln(&a, &b)
	fmt.Println(&a, &b)
}
