package main

import "fmt"

type account struct {
	balance   int
	firstNAme string
	lastName  string
}

// 포인터 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {

	var mainA *account = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)  //포인터 메서드 호출
	fmt.Println(mainA.balance) //70출력

	mainA.withdrawValue(20)    //값 타입 메서드 호출
	fmt.Println(mainA.balance) //여전히 70출력

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) //50출력

	mainB.withdrawPointer(30)  //포인터 메서드 호출
	fmt.Println(mainB.balance) //20 출력

}
