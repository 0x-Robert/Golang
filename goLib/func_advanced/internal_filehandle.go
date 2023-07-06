package main

import (
	"fmt"
	"os"
)

type Writer func(string)

// 이렇게 외부에서 로직을 주입하는 것을 의존성 주입이라고 한다.
// writeHello()함수 입장에서는 인수로 온 writer를 호출 할 때
// 파일에 쓰일지, 네트워크로 전송될지 로직을 알 수 없다.
func writeHello(writer Writer) {
	//writer 함수타입 변수 호출
}

func main() {
	f, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	writeHello(func(msg string) {
		//파일에 msg를 쓰는 함수 리터럴을 만들어서 wrtieHello()함수의 인수로
		//사용한다.
		fmt.Fprintln(f, msg) //함수 리터럴 외부 변수 f 사용
	})
}
