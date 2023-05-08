package main

import "fmt"

//go sum은 go tidy를 하면 생기는 파일
//대소문자별 특성
//대문자로 시작하면 외부사용이 가능하므로 타언어의 public과 비슷함
//소문자로 시작하면 타언어의 private와 유사함
//go언어는 카멜표기법사용을 권장함 _ 사용은 권하지 않음

var ExternalValue int //해당 변수는 전역적으로 사용가능
var internalValue int //해당 변수는 파일 내에서만 사용 가능

func ExternalFunc() int {
	//해당 함수는 전역으로 사용가능
}

func internalFunc() int {
	//해당 함수는 파일 내에서만 사용 가능
}

func main() {

	fmt.Println("hello")
}
