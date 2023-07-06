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

		인터페이스는 한국어로 상호작용면 

		추상화계층을 이용해 서로 결합을 끊는 것을 디커플링(decoupling)이라 함 
		결합도는 낮출수록 좋다. 

		go언어에서 어떤 타입이 인터페이스를 포함하고 있는지 여부를 결정할 때 덕 타이핑 방식을 사용한다. 
		덕 타이핑 방식이란 타입 선언시 인터페이스 구현 여부를 명시적으로 나타낼 필요 없이 인터페이스에 정의한 메서드 포함 여부만으로 결정하는 방식이다. 
		
		덕타이핑 유례 윗콤릴리가 말한 글귀
		만약 어떤 새를 봤는데 그 새가 오리처럼 걷고 오리처럼 날고 오리처럼 소리내면 나는 그 새를 오리처럼 부르겠다. 
	*/
}
