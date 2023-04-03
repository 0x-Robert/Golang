package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //2. 데이터를 계속 기다림
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() //4. 실행되지 않음, close가 없어서 계속 for range에서 기다리다가  모든 고루틴이 멈춤
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 //1. 데이터를 넣음
	}
	wg.Wait() //3. 작업 완료를 기다림
}
