package main

import ("fmt")


func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y 
}

func main() {
	fmt.Println("add1(x int, y int)의 결과 ", add1(42,13))
	fmt.Println("add2(x,y int)의 결과", add2(42,13))

}
