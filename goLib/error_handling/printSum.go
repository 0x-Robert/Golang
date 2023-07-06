package main

import "fmt"

func sum(nums ...int) int {
	if len(nums) == 0 {
		panic("nums should be not empty") //3. defer에서 recover가 실행되서 패닉에러를 가져옴 에러메시지 출력
	}
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func PrintSum() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(sum(1, 2, 3, 4, 5)) //1. 실행
	fmt.Println(sum())              //2. 실행 - 패닉발생, 함수가 종료됨
	fmt.Println(sum(1, 2, 3))       //PrintSum 함수가 패닉으로 인해 종료됨
}

func main() {
	PrintSum()
	fmt.Println("end of main") // panic이 복구됐기 때문에 해당 에러 출력됨
}
