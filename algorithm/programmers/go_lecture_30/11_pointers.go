// 포인터는 변수의 메모리 주소를 저장합니다.

// *T 타입은 T value의 포인터이며, zero value는 nil로 설정되어 있습니다.
// &연산자는 피연산자의 포인터를 생성하며, *연산자는 포인터가 가리키는 값을 참조합니다

// 또한 C와 달리 포인터 연산은 불가능합니다.

package main

import "fmt"

func main(){
	i, j := 42, 2701 

	p := &i 
	fmt.Println(*p)
	*p = 21 
	fmt.Println(i)
	

	p=&j //*p = j 
	*p = *p / 37 
	fmt.Println(j)
}