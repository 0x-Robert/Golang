package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("output.txt") //os 패키지의 Create함수를 이용해서 파일 생성, 같은 이름의 파일이 있으면 제거하고 새로생성, 성공하면 *File 객체 반환
	//*File 타입은 io.Writer를 구현하고 있기 때문에 Fprint() 함수의 인수로 사용할 수 있다.
	if err != nil { //에러가 true면 에러 실행
		fmt.Errorf("Create: %v\n", err)
	}
	fmt.Println("err1", err) //err은 nil

	defer f.Close() //함수 끝나기전에 파일 닫기

	const name, age = "Kim", 22
	n, err := fmt.Fprint(f, name, " is ", age, " years old.\n")
	//파일에 기록함 Kim is 22 years old.

	if err != nil {
		fmt.Errorf("Fprint: %v\n", err)
	}
	fmt.Println("err2", err)
	//-rw-rw-r-- 1 robertseo robertseo 21  4월  9 23:40 output.txt
	//21바이트 크기
	fmt.Print(n, " bytes written. \n")
}
