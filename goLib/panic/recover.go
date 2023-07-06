// Go 내장함수인 recover()함수는 panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수이다.
// recover 함수를 사용하면 panic 상태를 제거하고 openFile()의 다음 문장인 println() 을 호출하게 된다.
package main

import (
	"fmt"
	"os"
)

func main(){
	//잘못된 파일명을 넣음 
	openFile("2.txt")

	//recover에 의해 다음코드가 실행이 됨
	println("Done")
}

func openFile(fn string){
	//defer함수 panic 호출시 실행됨 
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil{
		panic(err)
	}

	defer f.Close()
}


