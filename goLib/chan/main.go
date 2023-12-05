package main

import "fmt"

func main() {
	// 채널 생성
	myChannel := make(chan string)

	go func() {
		// 값 전달
		myChannel <- "Wellcome"
		myChannel <- "to"
		myChannel <- "yoongrammer"

		// 채널 닫음
		close(myChannel)
	}()

	// 값 수신
	for msg := range myChannel {
		fmt.Println(msg)
	}
}
