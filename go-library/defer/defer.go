//지연 실행 defer
//Go 언어의 defer 키워드는 특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하게 한다. 일반적으로 defer는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용된다

// go 내장함수인 panic 함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수를 모두 실행한 후 즉시 리턴한다.
// 이런 패닉모드 실행방식은 상위함수에 계속 적용되고 콜스택을 타고 올라가며 적용함 결국에는 프로그램이 에러를 내고 종료함
package main

import "os"
func main(){
	f, err := os.Open("1.txt")
	if err != nil{
		panic(err)
	}
	//main 마지막에 파일 close 실행
	defer f.Close()

	//파일 읽기
	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}


