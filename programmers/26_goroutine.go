// 문제 설명
// goroutine은 Go runtime가 담당하는 경량(lightweight) 쓰레드 입니다.

// go f(x, y, z)
// 는 다음의 새로운 goroutine을 실행하며, f, x, y, z의 값을 구하는 것은 현재 goroutine에서 진행되고, f를 실행하는 건 새로운 goroutine에서 진행됩니다.

// f(x, y, z)
// 주의 goroutine은 같은 주소 공간을 쓰기 때문에, shared memory에 접근할 때는 동기화해줘야 합니다.

package main

import (
	"fmt"
	"time"
)

func say(s string){
	for i := 0; i<5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go say("2 다른 루틴")
	say("1 루틴")
}