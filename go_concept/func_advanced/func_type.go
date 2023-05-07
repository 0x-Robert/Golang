package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}
func sub(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

//다음과 같이 함수도 별칭 타입으로 만들 수 있다.
type opFunc func(int, int) int

func getOperator2() opFunc

func getOperator(op string) func(int, int) int { //op에 따른 함수 타입 반환
	if op == "+" {
		return add
	} else if op == "*" { //+나 *이 아니면 nil 반환
		return mul

	} else if op == "-" {
		return sub
	} else {
		fmt.Println("error")
		return nil
	}
}

func main() {
	//int타입 인수 2개를 받아서 int타입을 반환하는 함수 타입 변수
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)

	var operator2 func(int, int) int
	operator2 = getOperator("+")
	var result2 = operator2(2, 3)
	fmt.Println(result2)

	var operator3 func(int, int) int
	operator3 = getOperator("-")
	var result3 = operator3(10, 3)
	fmt.Println(result3)

	var operator4 func(int, int) int
	operator4 = getOperator("!@#!@$")
	var result4 = operator4(10, 3)
	fmt.Println(result4)

}
