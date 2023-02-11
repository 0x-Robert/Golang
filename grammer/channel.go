package main

import "fmt"

func main(){
	myChan := make(chan string)

	go func(){

		myChan <- "Welcome"
		myChan <- "Yongari"
		myChan <- "Blockchain"

		close(myChan)
	}()

	for msg:=range myChan{
		fmt.Println(msg)
	}
}

// 채널은 기본적으로 양방향 채널이지만 단방향 채널로도 사용할 수 있습니다.