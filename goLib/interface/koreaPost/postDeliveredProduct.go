package main

import (
	"github.com/tuckersGo/musthaveGo/ch20/fedex"
	"github.com/tuckersGo/musthaveGo/ch20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{} //*koreaPost.PostSender 타입 
	SendBook("어린 왕자", sender) 	   //타입이 맞지 않습니다. 
	SendBook("그리스인 조르바", sender) 
}
