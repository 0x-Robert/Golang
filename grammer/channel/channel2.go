package main

//unbuffered channel
// 이 채널에서는 하나의 수신자가 데이터를 받을 때까지
// 송신자가 데이터를 보내는 채널에 묶여 있게 된다.

func main() {
	//정수형 채널을 생성한다.
	ch := make(chan int)

	go func() {
		ch <- 123 //채널에 123을 보낸다.
	}()
	// for msg := range ch {
	// 	println(msg)
	// }

	var i int
	i = <-ch //채널로부터 123을 받는다.
	println(i)
}