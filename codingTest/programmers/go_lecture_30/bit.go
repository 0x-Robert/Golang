package main

import "fmt"

func needFloat(x float64 ) float64 {
	return x * 0.1
}

func main(){

	a :=  1               // 00000001
	b := a << 100           // 00011100: a의 비트를 오른쪽으로 2번 이동
	fmt.Printf("b", b) // 00011100


	const (
		big_const = 1 << 100
	)
	fmt.Println("big_const",needFloat(big_const))

}