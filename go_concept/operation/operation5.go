// ch6/ex6.6/ex6.6.go
package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c) // ❶ 0.3이 아니다 값은 2번 출력값과 같다.
	fmt.Println(a + b)                                    // ❷ 0.30000~4가 나온다. float 표현은 이런 오차를 가지고 있다.
}
