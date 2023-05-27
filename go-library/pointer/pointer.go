package main

import "fmt"

//Pointer를 언제 써야 하나? : 값복사가 발생하여 메모리를 낭비하고 싶지 않을 때

func main(){
	var d *int //포인터 변수 
	e := 3 
	d = &e 
	a := 2
	b := a 
	a = 10 
	c := &a 
	fmt.Println(&a, &b, c, *c, d, *d) // &를 붙여 메모리 주소 확인 가능 


	var numPtr *int = new(int) //자료형에 맞는 메모리 공간(int = 4byte)을 할당함
	*numPtr = 10 
	fmt.Println(numPtr)
	fmt.Println(*numPtr)

	var a1 int = 10
	var b2 int = 20 

	var p1 *int = &a1
	var p2 *int = &b2
	var p3 *int  = &a1 

	fmt.Printf("%v\n", p1 == p2)
	fmt.Printf("%v\n", p1 == p3)


	a3 := 2 
	b3 := &a3 
	fmt.Println(&a3, b3)

	*b3 = 5 
	fmt.Println(a3)


	x := 0
	foo(&x)
	fmt.Println("test",x)
}

func foo(x *int){
	fmt.Println("before", *x)
	*x = 40 
	fmt.Println("after", *x)
}