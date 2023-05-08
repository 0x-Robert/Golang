package main

import (
	"context"
	"fmt"
	"sync"
)
//동시성 프로그래밍 패턴을 위한 방법으로 구글에 golang concurrency patterns를 검색하면 활용법이 많다.
var wg sync.WaitGroup

func main() {

	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	//컨텍스트에 값을 추가한다.
	go square(ctx)
	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		//컨텍스트에서 값을
		//Value 메서드의 반환타입은 빈 인터페이스입니다.
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}
