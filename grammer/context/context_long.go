package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func longFunc() string{
	<-time.After(time.Second * 3)
	fmt.Println("longFunc")
	return "Success"
}

func main(){

	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		fmt.Println("cancel")
		//고루틴을 종료해야 할 상황이 되면 cancel 함수 실행 
		cancel()
	}()
//jobCount만큼 여러 개의 고루틴을 만들어 longFuncWithCtx 수행 
	var wg sync.WaitGroup
	for i := 0; i < jobCount; i++{
		wg.Add(1)

		go func(){
			defer wg.Done()
			result, err := longFuncWithCtx(ctx)
			fmt.Println("result",result)
			if err != nil{
				fmt.Println("error")
			}
		}()
	}
	fmt.Println("main")
	
	
}


func longFuncWithCtx(ctx context.Context) (string, error){
	done := make(chan string)

	go func(){
		done <- longFunc()

	}()

	select {
	case result := <-done:
		return result, nil 
	case <-ctx.Done():
		return "Fail", ctx.Err()
	}
}