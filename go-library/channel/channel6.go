package main

import "fmt"

func main(){
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

//송신 파라미터는 (p chan<-int)와 같이 
//chan<-을 사용 
func sendChan(ch chan <- string){
	ch <- "Data"
	//x := <-ch //에러발생 
	fmt.Println("ch", ch)
	//fmt.Println("x", x)
}

//수신 파라미터는 p <-chan int
//<-chan 
func receiveChan(ch <-chan string){
	data := <-ch 
	fmt.Println("data",data)
}