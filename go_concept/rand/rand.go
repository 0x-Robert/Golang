package main

import (
	"fmt"
	"math/rand"

	// _ "string"
	"time"
)

func main(){
	t := time.Now() //현재시각제공 
	rand.Seed(t.UnixNano()) //시간을통해 랜덤시드를 설정 > 매번 다른 랜덤값 생성 
	for i := 0; i < 10; i++{
		fmt.Println(rand.Intn(100))
	}
}