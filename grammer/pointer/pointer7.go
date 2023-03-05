package main

import "fmt"

func add(p1, p2, p3 *int) {
	fmt.Println("1", p1, p2, p3)
	//실제 주소값을 넣어서 계산이 완료됨
	*p3 = *p1 + *p2
	fmt.Println("2", p1, p2, p3)
}

func add1(p1, p2, p3 *int) {
	*p3 = *p1 + *p2
}

func main() {
	a := 3
	b := 5
	c := 0

	a1 := 3
	b1 := 5
	c1 := 0
	//주소값을 인자에 넣는다.
	add(&a, &b, &c)
	add(&a1, &b1, &c1)
	fmt.Println(c)
	fmt.Println(c1)
}
