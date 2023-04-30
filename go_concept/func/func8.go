package main

import "fmt"

func AAA() {
	fmt.Println("start AAA()")
	BBB()
	fmt.Println("end AAA()")

}

func BBB() {
	fmt.Println("BBB()")
}

func main() {
	AAA()
}
