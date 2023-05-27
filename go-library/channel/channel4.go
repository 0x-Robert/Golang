package main

import "fmt"

func main(){
	c := make(chan int)
	c <- 1 //수신루틴이 없으므로 데드락 
	fmt.Println(<-c) //코멘트해도 데드락 별도 고루틴이 없어서
}