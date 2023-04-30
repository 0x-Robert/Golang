package main

import "fmt"

func Print(args ...interface{}) string { //모든 타입을 받는 가변 인수

	for _, arg := range args { //모든 인수 순회
		switch f := arg.(type) { //인수의 타입에 따른 동작

		case bool:
			val := arg.(bool) //인터페이스 변환
			fmt.Println(val)
		case float64:
			val := arg.(float64)
			fmt.Println(val)
		case int:
			val := arg.(int)
			fmt.Println(val)
		default:
			fmt.Println(f)
		}

	}
	return "0"
}

func main() {
	Print(2, "hello", 3.14)
}
