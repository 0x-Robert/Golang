package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU()) //모든 cpu 사용 

	var data = []int{} //int형 슬라이스 생성
	var mutex = new(sync.Mutex) //mutex 변수 생성 

	go func(){
		for i:=0; i < 1000; i++{
			mutex.Lock() //뮤텍스 잠금, data 슬라이스 보호 시작
			// 고루틴에서 1000번 반복하면서 data 슬라이스에 1을 추가
			data = append(data, 1)
			mutex.Unlock() //뮤텍스 잠금 해제, data 슬라이스 보호 종료
			runtime.Gosched( ) //다른 고루틴이 CPU를 사용할 수 있도록 양보 
		}
	}()

	go func(){
		for i:= 0; i < 1000; i++{
			mutex.Lock() //뮤텍스 잠금, data 슬라이스 보호 시작
			//고루틴에서  1000번 반복하면서 data 슬라이스에 1을 추가
			data = append(data, 1) 
			mutex.Unlock() //뮤텍스 잠금해제, data 슬라이스 보호 종료 
			runtime.Gosched() //다른 고루틴이 CPU를 사용할 수 있도록 양보 
		}
	}()

	time.Sleep(2 * time.Second) //2초 대기
	fmt.Println(len(data)) // data 슬라이스의 길이 출력
// }
// 뮤텍스는 sync.Mutex를 할당한 뒤에 고루틴에서 Lock, Unlock 함수로 사용합니다. 보호를 시작할 부분에서 Lock 함수를 사용하고, 보호를 끝낼 부분에서 Unlock 함수를 사용합니다. Lock, Unlock 함수는 반드시 짝을 맞추어주어야 하며 짝이 맞지 않으면 데드락(deadlock, 교착 상태)이 발생하므로 주의합니다.

// 여기서는 data 슬라이스를 보호할 것이므로 두 고루틴 모두 data = append(data, 1) 부분 위 아래로 Lock, Unlock 함수를 사용합니다. 이제 실행을 해보면 정확히 2000이 출력됩니다.
//실제로 출력결과가 2000임

