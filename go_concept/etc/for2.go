package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Print(i)
	}

	for j := 5; j >= 0; j-- // for 문 끝에 세미콜론이 삽입되어 컴파일 오류 발생
	{
		fmt.Print(j)
	}
}
