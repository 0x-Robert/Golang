package main

func main() {
	ch := make(chan int, 2)

	//채널에 송신
	ch <- 1
	ch <- 2
	//ch <- 3
	//ch <- 435324234
	//ch <- 3
	//채널 닫기
	close(ch)

	//fatal error: all goroutines are asleep - deadlock!
    // for {
    //     if i, success := <-ch; success {
    //         println(i)
    //     } else {
    //         break
    //     }
    // }

	for i := range ch {
        println(i)
    }

}