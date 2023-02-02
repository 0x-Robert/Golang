package main

import "fmt"

func main() {
    //① 버퍼크기가 1인 채널을 만듭니다
    ch := make(chan int, 5)
    //② 채널에 1을 전달합니다. 버퍼가 꽉 찼습니다.
    ch <- 1
    //③ 버퍼가 꽉 찬 상태에서 2를 또 전달합니다 - 버퍼를 비워줄 루틴이 없어 데드락이 발생합니다.
    ch <- 2
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}