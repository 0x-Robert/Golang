package main

import "fmt"

func solution(num int) (answer string) {

	if num%2 == 0 {
		//fmt.Println("Even")
		return "Even"
	} else {
		//fmt.Println("Odd")
		return "Odd"
	}
}

func main() {
	num1 := 4
	//num2:=5
	solution(num1)
}
