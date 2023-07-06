package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)

	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) //채널을 닫음으로서 모든 고루틴이 정상종료됨,
	//이렇게 채널을 제 때 닫아주지 않아서 고루틴에서 데이터를 기다리며 무한 대기하는 경우를 좀비루틴, 고루틴 릭(leak)이라고 부른다.
	wg.Wait()
}
