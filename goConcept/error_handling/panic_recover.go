package main

import "fmt"

func f() {
	fmt.Println("f()함수 시작")
	defer func() { //5. recover 함수로 panic 복구
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()                     //2. g() 호출
	fmt.Println("f() 함수 끝") //6. 함수 끝
}

func g() {
	//3. g() >>> h() 함수 호출
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.") //4. b가 0인 panic 발생
	}
	return a / b
}

func main() {
	f() //1. f() 함수 호출
	fmt.Println("프로그램이 계속 실행됨")
}
