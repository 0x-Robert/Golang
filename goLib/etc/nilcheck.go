package main

import "fmt"

func check_a_b(){
	var a []int 
	var b []int 

	fmt.Println(a == nil) //true 
	//# command-line-arguments
	///nilcheck.go:10:14: invalid operation: a == b (slice can only be compared to nil)
	//다음 코드는 컴파일 에러가 나온다. 슬라이스는 nil하고의 비교만 가능하다 서로 다른 두 slice는 == 연산자로 비교할 수 없다.
	fmt.Println(a == b)
}

func main(){
	var a []int 
	b := make([]int, 0)
	c := []int{}
	var d []int 
	d = nil
	e := make([]int, 0, 0)
	f := make([]int, 0, 1)

	// 명시적인 초기값 없이 선언된 변수는 zero value를 갖게된다. 
	// b, c, e, f는 초기값이 있다. 
	fmt.Println(
		a == nil, //true
		b == nil, //false
		c == nil, //false 
		d == nil, //true 
		e == nil, //false 
		f == nil)
		check_a_b()
}

