package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	//우체국 전송객체, fedex 전송 객체 모두 SendBook 인수로 사용가능
	//우체국 전송객체를 만든다.
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	//Fedex 전송 객체를 만든다.
	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)

	/*

		내부 동작을 감춰서 서비스를 제공하는 쪽과 사용하는 쪽 모두에게 자유를 주는 방식을 추상화라고 한다.
		인터페이스는 추상화를 제공하는 추상화계층이다.

	*/
}
