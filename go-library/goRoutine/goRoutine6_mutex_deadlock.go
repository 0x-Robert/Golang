package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(2)
	fork := &sync.Mutex{} //포크와 수저 뮤텍스
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저") //포크먼저
	go diningProblem("B", spoon, fork, "수저", "포크") //수저 먼저
	wg.Wait()                                      //작업이 완료될때까지 대기
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다. \n", name)
		first.Lock() //첫 번째 뮤텍스를 획득 시도
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock() //두 번째 뮤텍스를 획득 시도
		fmt.Printf("%s 밥을 먹습니다. \n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock() //뮤텍스 반납
		first.Unlock()
	}
	wg.Done() //작업이 끝날때까지 대기
	//위 코드는 일부러 데드락을 발생시킨 예제

}
