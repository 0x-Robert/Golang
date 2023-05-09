package main

import "log"

func server(workChan *Work) {
	for i = 0; i <= 10; i++ {
		go safeDo(work)
	}
}

func safeDo(work *Work) {
	defer func() {
		//종료 및 실패된 work만 리커버
		if err := recover(); err != nil {
			//리커버 실패시 에러 반환
			log.Println("work failed", err)
		}
	}()
	do(work)
}
