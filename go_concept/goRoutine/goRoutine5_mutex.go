package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex //패키지 전역 변수 뮤텍스, 동시성 문제를 해결할 뮤텍스

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         //뮤텍스 획득
	defer mutex.Unlock() //defer를 사용한 Unlock() mutex를 락하면 반드시 Unlock을 해줘야한다.
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
