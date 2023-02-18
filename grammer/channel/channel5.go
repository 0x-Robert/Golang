package main

import "fmt"

//다음과 같이 channel4.go와 다르게 channel5.go에서 버퍼채널을 사용하면 에러가 발생하지 않는다.

func main(){
	ch := make(chan int, 1)
	ch <- 123 
	fmt.Println(<-ch)
}