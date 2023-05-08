package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 컨텍스트는 작업을 위한설정, 취소를 할 수 있는 작업명세서 같은 것이다.
func main() {
	wg.Add(1)                                               //작업단위 추가
	ctx, cancel := context.WithCancel(context.Background()) // 취소 가능한 컨텍스트 생성, 상위 컨텍스트가 없으면 가장 기본적인 context.Background()를 넣어준다.
	//이 함수는 값을 두 개를 반환하는데 첫 번째가 컨텍스트 객체고 두 번째가 취소함수다.
	go PrintEverySecond(ctx)    //고루틴 생성
	time.Sleep(5 * time.Second) //5초 이후
	cancel()                    //취소 함수 호출 >> 작업취소
	wg.Wait()                   //작업이 완료될때가지 대기
}

func PrintEverySecond(ctx context.Context) { //루틴에서 인수로 받은 컨텍스트의 Done()채널의 시그널이 있는지 검사한다.
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done() //컨텍스트의 Done()채널의 시그널을 보내 작업자가 작업취소를 할 수 있도록 알림 , 컨텍스트가 완료될때 Done
			//채널에 시그널을 넣기 때문에 여기서 메시지를 수신하면 고루틴 종료
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
