// go 내장함수인 panic 함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수를 모두 실행한 후 즉시 리턴한다.
// 이런 패닉모드 실행방식은 상위함수에 계속 적용되고 콜스택을 타고 올라가며 적용함 결국에는 프로그램이 에러를 내고 종료함

package main

import "os"

func main(){
	openFile("2.txt")
	//openFile 안에서 panic이 실행되면 다음 코드는 실행 안됨
	println("Done")

}

func openFile(fn string){
	f, err := os.Open(fn)
	if err != nil{
		panic(err)
	}
	defer f.Close()
}

//코드를 실행하면 다음과 같은 에러 출력함
// panic: open 2.txt: no such file or directory

// goroutine 1 [running]:
// main.openFile({0x470ec5?, 0xc0000061a0?})
//         /home/robertseo/work/go_study/grammer/panic.go:18 +0x88
// main.main()
//         /home/robertseo/work/go_study/grammer/panic.go:9 +0x25
// exit status 2