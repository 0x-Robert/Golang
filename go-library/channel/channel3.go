package main

import "fmt"

//unbuffered channel
func main(){
	done := make(chan bool)
	go func(){
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true 
	}()

	//위의 go 루틴이 끝날때까지 대기 
	println("before done")
	<-done 
	println("after done")
}