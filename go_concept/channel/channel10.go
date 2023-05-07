package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) //채널 생성

	wg.Add(1)          //작업 추가
	go square(&wg, ch) //고루틴 생성
	ch <- 9            //채널에 데이터 넣음
	wg.Wait()          //작업이 완료되길 기다림
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch                       //채널에 있는 데이터를 n이라는 변수로 뺌
	time.Sleep(time.Second)         //1초 대기
	fmt.Printf("Square: %d\n", n*n) //제곱 출력
	wg.Done()                       //작업을 끝냄
}
