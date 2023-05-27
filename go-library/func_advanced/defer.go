package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, err := os.Create("test.txt") //파일 생성
	fmt.Println("type check", reflect.TypeOf(f))
	if err != nil { //에러 확인
		fmt.Println("Failed to create a file") //Failed 문 출력
		return
	}

	defer fmt.Println("반드시 호출됩니다.") //지연 수행될 코드 5
	defer f.Close()                 //지연 수행될 코드 4
	defer fmt.Println("파일을 닫았습니다.") //지연 수행될 코드  3

	fmt.Println("파일에 Hello Yongari를 씁니다.") // 1
	fmt.Println(f, "Hello Yongari")        //파일에 텍스트를 쓴다. 2
}

/*

네, 맞습니다. 출력된 값 &{0xc000062180} 은 메모리 주소값입니다.
이는 파일 포인터(f)가 가리키는 파일 객체(os.File)의 메모리 주소를 나타냅니다.

출력결과
파일에 Hello Yongari를 씁니다.
&{0xc0000b2120} Hello Yongari
파일을 닫았습니다.
반드시 호출됩니다.

type check *os.File
파일에 Hello Yongari를 씁니다.
&{0xc000062180} Hello Yongari
파일을 닫았습니다.
반드시 호출됩니다.

*/
