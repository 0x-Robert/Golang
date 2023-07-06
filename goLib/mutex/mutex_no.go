package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 cpu 사용 

	var data = []int{} //int형 슬라이스 생성

	go func(){
		for i:=0; i < 1000; i++{
			// 고루틴에서 1000번 반복하면서 data 슬라이스에 1을 추가
			data = append(data, 1)
			runtime.Gosched( ) //다른 고루틴이 CPU를 사용할 수 있도록 양보 
		}
	}()

	go func(){
		for i:= 0; i < 1000; i++{
			//고루틴에서  1000번 반복하면서 data 슬라이스에 1을 추가
			data = append(data, 1) 
			runtime.Gosched() //다른 고루틴이 CPU를 사용할 수 있도록 양보 
		}
	}()

	time.Sleep(2 * time.Second) //2초 대기
	fmt.Println(len(data)) // data 슬라이스의 길이 출력
}

// 실행을 해보면 대략 1800~1990 사이의 값이 나옵니다. data 슬라이스에 1을 2,000번 추가했으므로 data의 길이가 2000이 되어야 하는데 그렇지가 않습니다. 두 고루틴이 경합을 벌이면서 동시에 data에 접근했기 때문에 append 함수가 정확하게 처리되지 않은 상황입니다. 이러한 상황을 경쟁 조건(Race condition)이라고 합니다.

// runtime.Gosched 함수는 다른 고루틴이 CPU를 사용할 수 있도록 양보(yield)합니다. 지금까지 time.Sleep 함수를 사용했지만 runtime.Gosched 함수가 좀 더 명확합니다