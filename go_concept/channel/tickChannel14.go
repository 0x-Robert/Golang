package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            //1초 간격 시그널
	terminate := time.After(10 * time.Second) //10초 간격 시그널

	for {
		select { //select문을 이용해서 tick, terminate,ch순서대로 채널에서 데이터읽기를 시도함, terminate에서 신호를 못 읽으면 ch에서 읽어오게 된다. tick은 1초 간격으로 신호를 보내고 
			//10초 이후에는 terminate 신호가 오므로 함수가 종료된다. 
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
