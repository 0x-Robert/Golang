package main

import (
	"fmt"
)

//defer 키워드는 해당 함수를 나중에 실행하게 합니다.
//즉, defer를 호출하는 함수(하기에서는 main)가 반환되기 전에 즉각 함수 호출을 실행하도록 예약합니다.
//명시적 선언으로 인해 메모리 해제에 대한 부담감을 해소할 수 있습니다. - 함수 종료 시 메모리 해제
//지연된 함수는 LIFO(스택) 순서에 의해 실행됩니다

func one() {
	fmt.Println("1111")
}

func two() {
	fmt.Println("2222")
}

func main() {
	defer two()
	one()
}
