package main

import "fmt"

func main() {
	a := 1
	b := 2

	_ = b // 사용하지 않는 변수로 인한 컴파일 에러 방지

	fmt.Println(a)

	var a1 int = 10
	fmt.Println("a1", a1)
	a1 = 20
	fmt.Println("a1", a1)

	var test string = "hello"
	fmt.Println("test", test)
	test = "test test"
	fmt.Println("test", test)

	var a2 int = 3
	var b2 int
	var c2 = 4
	d := 5
	f := 3.14
	var e = "hello"
	fmt.Println(a2, b2, c2, d, e, f)

	// a3 := 3
	// var b3 float64 = 8.5
	// d3 := a3 * b3
	// fmt.Println(d3)

	a3 := 3
	var b3 float64 = 8.5
	var c3 = int(b3)
	d3 := a3 * c3
	fmt.Println(d3)

	//큰 그릇에서 작은 그릇으로 데이터를 옮기면 값이 버려질 수 있음
	//타입변환할 때 조심해야함 큰 타입에서 작은 타입으로 변환하면 값이 변할 수 있다.

	var a5 int16 = 3456    //int16타입 2바이트 정수
	var b5 int8 = int8(a5) // int16 > int8
	fmt.Println(a5, b5)

	var a6 float32 = 1234.523
	var b6 float32 = 3456.523
	var c6 float32 = a6 * b6
	var d6 float32 = c6 * 3
	fmt.Println(a6, b6, c6, d6)

}
