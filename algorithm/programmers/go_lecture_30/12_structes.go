// 구조체는 필드(들)의 집합으로, 다음과 같이 선언할 수 있습니다.

// type name struct {}
// 위 선언에서, 구조체 type 이름은 name으로 쓸 수 있으며, 구조체의 필드는 .로 접근할 수 있습니다

// 구조체 인스턴스를 생성할 때, 특정 필드만 초기화하고 싶으면 초기화할 필드: 초기화할 값을 지정합니다
package main

import "fmt"

type Vertex struct {
	X int 
	Y int 
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}

)

func main(){
	fmt.Println("v1.X", v1.X)
	v1.X = 4 
	fmt.Println("v1.X=4로 바꾼 v1.X", v1.X)


	var p = &v1 
	p.X = 10 
	fmt.Println("포인터로 바꾼 v1.x값", v1.X)
}