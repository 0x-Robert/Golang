package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//컨텍스트 = 맥락을 유지하기위한 통로


func main() {
	//채널이란 고루틴을 연결해주는 통로(파이프)다. 
	//채널은 양방향이고 고루틴은 채널을 통해 데이터를 주고 받는다. 
	//채널은 동일한 유형의 데이터만 전송할 수 있음 
	
	exit := make(chan os.Signal)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	// 컨텍스트를 생성하는 방법은 context.Background() 함수를 사용하여 생성하는 것 
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("ctx",ctx)

	go func() {
		fmt.Println("Signal:", <-exit)
		cancel()
	}()

	start := time.Now()
	result, err := longFuncWithCtx(ctx)
	fmt.Printf("duration:%v result:%s\n", time.Since(start), result)
	if err != nil {
		log.Fatal(err)
	}

}

func longFuncWithCtx(ctx context.Context) (string, error) {
	done := make(chan string)

	go func() {
		done <- longFunc()
	}()

	select {
	case <-ctx.Done():
		return "Fail", ctx.Err()
	case result := <-done:
		return result, nil
	}
}

func longFunc() string {
	<-time.After(time.Second * 3)
	return "Success"
}
