package main

import (
	"fmt"
	"sync"
	"time"
)

// select 문은 여러 채널을 동시에 기다린다. 하나의 채널이라도 데이터를 읽어오면 해당 구문을 읽고 select문은 종료됨
func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { //ch와 quit 양쪽 모두 기다림
		case n := <-ch: //ch값을 n으로 뺀다.
			fmt.Printf("Square: %d\n", n*n) //n 제곱 출력
			time.Sleep(time.Second)         //1초 대기
		case <-quit:
			wg.Done()
			return
		}

	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool) //종료 채널

	wg.Add(1)                //작업 추가
	go square(&wg, ch, quit) //제곱근 함수 고루틴 생성

	for i := 0; i < 10; i++ {
		ch <- i * 2 //ch 채널에 데이터 삽입
	}

	quit <- true //quit에 true값 입력
	wg.Wait()
}
