package main

/*
각 고류틴이 서로 다른 자원에 접근하게 만드는 두 가지 방법
1. 영역을 나누는 방법
2. 역할을 나누는 방법

다음코드는 영역을 나누는 방법이다.
*/
import (
	"fmt"
	"sync"
	"time"
)

type Job interface { //각 작업을 나타내는 Job 인터페이스를 설정함 Job 인터페이스는 Do 메서드만 가지고 있다.
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
