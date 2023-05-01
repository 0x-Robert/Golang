package main

import "fmt"

//변수를 하나 선언
var num1 int 

//같은 타입을 가지는 변수를 여러 개 선언
var num2, num3 int 

//여러 변수에 한 번에 값을 초기화 : 선언과 동시에 
var num4, num5, str1 = 4, 5, "example"

//함수 밖에서는 :=를 쓸 수 없음
//errorvar := str 


//다른 타입을 가지는 변수를 여러 개 선언 
var (
	i int 
	b bool 
	s string 
)


func main(){
	fmt.Println("1",num1 )
	fmt.Println("2",num2)
	fmt.Println("3",num3)
	fmt.Println("4",num4, num5, str1)

	//함수 안에서는 :=을 쓰면 var와 타입을 지정하지 않고 변수를 선언과 동시에 초기화할 수 있다. 
	num6 := 6 
	fmt.Println("5",num6)
	fmt.Println("6",i,b,s)
}