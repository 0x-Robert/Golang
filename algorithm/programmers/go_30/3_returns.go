package main 

import "fmt"

//return 뒤에 리턴 타입을 적어주는 방법
func divide1(dividend, divisor int)(int, int){
	var quotient = (int)(dividend/divisor)
	var remainder = dividend%divisor
	return quotient, remainder
}

//return 뒤에 리턴할 변수를 선언하는 방법, 위와 달리
//내부에서 var를 선언하지 않고 바로 쓰기
func divide2(dividend, divisor int)(quotient, remainder int){
	quotient = (int)(dividend/divisor)
	remainder = dividend%divisor 
	return 
	//return이라고만 적으면 미리 return 값으로 정해놓은 quotient와 remainder를
	//return 합니다. 
}

func main(){
	//한 번에 여러개의 결과를 return하는 부분
	var quotient, remainder int 
	quotient, remainder = divide1(10,3)
	fmt.Println("1 결과", quotient, remainder)

	//한 번에 여러개의 결과를 return 받는 부분 
	quotient, remainder = divide2(10,3)
	fmt.Println("2 결과", quotient, remainder)
}

