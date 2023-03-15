// ch15/ex15.15/ex15.15.go
package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str) // ❶ 슬라이스로 타입 변환, 문자열은 byte 배열이므로 []byte로 타입변환하려면 같은 메모리 공간을 가리킬 것이라고 생각할 수 있으나
	//str이 가리키는 메모리공간과 slice가 가리키는 메모리 공간은 다르다.

	//str[2] ='a'
	//str과 slice가 다른 메모리 공간을 가리키기 때문에 slice에서 3번째 요소를 a로 바꿔도 str은 그대로인 것 을 알 수 있다.
	slice[2] = 'a' // ❷ 3번째 문자 변경

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
