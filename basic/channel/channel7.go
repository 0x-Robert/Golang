package main

func main(){
	ch := make(chan int, 2)

	//채널에 송신
	ch <- 1 
	ch <- 2

	//채널을 닫는다 해당 채널로는 더 이상 송신 불가
	//채널이 닫혀도 수신은 계속 가능함
	//<-ch는 두개의 리턴값을 갖는다. 첫째는 채널메시지, 둘째는 수신이 제대로 되었는가를 나타냄 
	close(ch)

	//채널수신 
	
	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success{
		println("더 이상 데이터 없음")
	}
	//채널을 닫았는데 데이터 수신을 시도하면 두번재 리턴값이 false 
}


