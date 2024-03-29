package main

import "fmt"

// 함수 리터럴 외부변수를 내부상태로 가져오는 것을 캡쳐라고 한다.
// 캡쳐는 값보사가 아니라 참조형태로 가져온다.
// 333이 출력됨
func CaptureLoop() {
	f := make([]func(), 3) //함수 리터럴 3개를 가진 슬라이스
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		//i변수가 참조로 캡쳐되서 그렇다. 
		//for문이 진행할 때마다 i값이 증가하고 최종적으로 i=3에서 종료하기 때문에 
		//함수 리터럴이 호출되는 시점의 캡쳐된 i 값은 3이된다.
		f[i] = func() {
			fmt.Println(i) //캡쳐된 i값 출력, i가 3일 때 캡처
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

// 0,1,2가 출력됨
func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		v := i //v 변수에 i값 복사
		f[i] = func() {
			fmt.Println(v) //캡쳐된 v값 출력
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
