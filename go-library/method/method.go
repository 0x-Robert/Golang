package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {

	a := &account{100}
	println("a", a, a.balance)
	withdrawFunc(a, 30)
	println("a", a, a.balance)
	a.withdrawMethod(30)
	println("a", a, a.balance)
	//a는 40?
	fmt.Printf("%d\n", a.balance)
}
