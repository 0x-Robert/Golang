package main

import (
	"fmt"
)

//크기가 고정된 배열과 달리 슬라이스는 자유롭게 요소를 추가할 수 있다.
//make(슬라이스 타입, 슬라이스 길이, 슬라이스의 용량)
func uglyNumbers(n int) int{
	//1은 첫 번째 ugly number이므로 배열에 미리 추가한다.
	uglyArr := []int{1}

	idx2, idx3, idx5 := 0,0,0 

	for len(uglyArr) < n {
		// 2, 3, 5를 곱한 수 중 가장 작은 수를 선택하기
		nextUgly := min(uglyArr[idx2]*2, uglyArr[idx3]*3, uglyArr[idx5]*5, )
		fmt.Println("nextUgly",nextUgly)

		//선택된 ugly number를 배열에 추가한다.
		//append(배열, 추가할 원소 )
		uglyArr = append(uglyArr, nextUgly)
		fmt.Println("uglyArr",uglyArr)

		if nextUgly == uglyArr[idx2]*2{
			idx2 += 1
		}
		if nextUgly == uglyArr[idx3]*3{
			idx3 += 1 
		}
		if nextUgly == uglyArr[idx5]*5{
			idx5 += 1
		}

	}
	return uglyArr[n-1]
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:]{
		if num < res {
			res = num 
		}
	}
	return res 
}

func main(){
	//:= 선언, 양도 및 재선언을위한 것이고, 변수의 유형도 자동으로 추론
	//(참고로 함수내에서만 := 을 사용할 수 있다.)
	//주의할 점은 :=을 이용해 이미 존재하던 변수에 재선언을 하는 경우 variables scope를 주의해야한다.
	result := uglyNumbers(1)
	fmt.Println(result) //1 

	//= 연산자는 변수에 값을 할당할 때 사용합니다. 예를 들어 a = 10 은 변수 a에 10을 할당하는 코드입니다.
	// := 연산자는 변수에 값을 할당할 때 사용하며, 초기화와 선언을 동시에 수행합니다. 즉, := 연산자를 사용하면 변수를 선언하고 값을 할당하는 것이 가능합니다. 예를 들어 a := 10은 변수 a를 선언하면서 10을 할당하는 코드입니다.
	// 따라서, Go 언어에서 변수를 선언하고 초기화하려면 := 연산자를 사용해야 합니다. 만약 이미 선언된 변수에 값을 할당하려면 = 연산자를 사용합니다.
	result = uglyNumbers(3)
	fmt.Println(result) //3
}

