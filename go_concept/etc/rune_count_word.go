// ch15/ex15.5/ex15.5.go
package main

import "fmt"

func main() {
	str := "Hello World"

	//[]rune 타입은 상호 타입 변환이 가능하다.
	//문자열은 각  문자의 코드값의 배열인 rune 배열로 나타낼 수 있다.
	// ❶ ‘H’, ‘e’, ‘l’, ‘l’, ‘o’, ‘ ‘, ‘W’, ‘o’, ‘r’, ‘l’, ‘d’ 문자코드 배열
	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 107, 104, 65, 65}

	fmt.Println(str)
	fmt.Println(string(runes))
}
