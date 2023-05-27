package main

import (
	"fmt"
	"log"
)

func helloOne(n int)(string, error){
	if n == 1{ //1일때 
		s := fmt.Sprintln("Hello", n) //hello 문자열을 리턴 
		return s, nil  //정상 동작 에러는 nil 
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) //1이 아닐때 에러 리턴 
}

func main(){
	s, err := helloOne(1) //매개변수에 1을 넣어서 정상
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(s) //Hello 1 

	s, err = helloOne(2) //매개변수에 2를 넣어서 에러발생
	if err != nil {
		log.Fatal(err) //에러 문자열이 출력되고 프로그램 종료 
	}

	//에러가발생하고 프로그램 종료되서 다음 코드는 실행되지 않음
	fmt.Println(s)
	fmt.Println("Hello World!")
}
//output
// Hello 1

// 2023/02/05 18:35:22 2는 1이 아닙니다.
// exit status 1