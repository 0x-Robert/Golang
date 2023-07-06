package main

import (
	"bufio" //bufio는 입력스트림으로부터 한 줄을 읽는 Reader 객체를 제공합니다. io를 담당하는 패키지
	"fmt"
	"os" //표준 입출력등을 가지고 있는 패키지
)

func main() {

	//인수로 입력되는 입력스트림을 가지고 Reader 객체를 생성해줍니다.
	//표준 입력을 읽는 객체
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil { //에러 발생시
		fmt.Println(err)       //에러 출력
		stdin.ReadString('\n') //표준 입력 스트림 지우기
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b) // 다시 입력 받기
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}
