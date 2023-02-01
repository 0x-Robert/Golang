// 문제 설명
// 타입 [n]T는 타입 T 값을 n개 저장하는 배열이며, 배열 크기는 한 번 설정하면 바꿀 수 없습니다.

// var a [10]int
// 위 코드는 a를 정수 10개를 저장하는 배열로 선언합니다.

package main

import "fmt"

func main(){
	var a [2]string 
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("a[0], a[1]", a[0], a[1])
	fmt.Println("a:", a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println("primes", primes)
}