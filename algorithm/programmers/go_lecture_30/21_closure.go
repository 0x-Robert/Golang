// 문제 설명
// 자기 바디 외부에 있는 변수를 참조하는 함숫값을 클로져라고 합니다. 함수가 자신이 참조하는 변수에 접근하거나 값을 변경하는 경우, 함수가 변수에 bound되었다고 합니다

// 예를 들어 이 코드에서 adder는 클로져를 리턴하며, 클로져 pos와 클로져 neg는 서로 다른 변수 sum을 가집니다.

package main

import "fmt"

func adder() func(int) int {
	sum := 0 
	return func(x int) int {
		sum += x 
		return sum 
	}
}

func main(){
	//pos, neg는 서로 다른 변수 sum을 가집니다. 
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++{
		fmt.Println(
			i, ":",
			pos(i),
			neg(-2*i), 
		)
	}
}