package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}
//함수 호출시 입력하는 값은 argument 
//함수가 외부로부터 입력받는 변수는 parameter 

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}
