package main

import "time"

func main(){
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

	EXIT: 
		for{ // for 루프문 안에서 select가 2개의 고루틴이 실행되기를 기다리고 있음 
			select {
			case <-done1:
				println("run1 완료")
			case <-done2:
				println("run2 완료")
				//브레이크 exit으로 루프문 탈출 
				break EXIT 
			}
		}
	}
	//1초간 실행되고 done1채널로부터 수신하고 case 실행 후 다시 for 루프를 돌음 
	func run1(done chan bool){
		time.Sleep(1 * time.Second)
		done <- true 
	}
	//2초간 실행되고 done2 채널로부터 수신해서 
	func run2(done chan bool){
		time.Sleep(2 * time.Second)
		done <- true 
	}