package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second) //메인함수에서 3초가 대기하지 않으면 다른 고루틴 포함해서 메인 고루틴도 종료됨 바로 종료됨
}
