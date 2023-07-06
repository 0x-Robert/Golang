package main

import (
	"fmt"
	"time"
)

func runLoopSend(n int, ch chan int) {
	for i := 0; i < n; i++ {
		//채널 입력
		ch <- i
	}
	//채널 닫기
	close(ch)
}

func runLoopReceive(ch chan int) {
	for {
		//채널 출력
		//채널이 닫히면 ok는 false, 열려있으면 ok는 true
		i, ok := <-ch
		if !ok {
			break
		}
		//Println 출력
		fmt.Println("Received value:", i)
	}
}

func main() {
	myChannel := make(chan int)
	go runLoopSend(10, myChannel)

	go runLoopReceive(myChannel)
	//다음과 같이 주석하면 동작하지 않음 메인에서 끝남
	//time.Sleep(2 * time.Second)
	time.Sleep(2 * time.Second)
}
