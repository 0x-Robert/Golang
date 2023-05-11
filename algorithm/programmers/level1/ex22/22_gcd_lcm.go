package level1

import "fmt"

func solution222(n int, m int) []int {
	gcd := gcd(n, m)
	fmt.Println("[]int{gcd, n * m / gcd}", []int{gcd, n * m / gcd})
	return []int{gcd, n * m / gcd}
}

func gcd(a, b int) int {
	fmt.Println(a, b)
	//a = 2 / b=5
	if a < b {
		a, b = b, a
	}
	//a=5, b=2
	fmt.Println(a, b)
	//반복문을 돌면서 실행
	for b != 0 {
		fmt.Println("before", a, b)
		//  a= 2 , b=1
		a, b = b, a%b
		fmt.Println("after", a, b)
	}
	fmt.Println("fin", a, b)
	//a=1
	return a
}

func gcdLcm() {
	n := 2
	m := 5
	solution222(n, m)
}
