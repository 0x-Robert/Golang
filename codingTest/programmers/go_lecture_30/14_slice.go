// 문제 설명
// 배열은 고정 길이인 반면 슬라이스는 가변 길이로, 슬라이스를 쓰면 배열을 동적인 것처럼 쓸 수 있습니다.

// []T는 타입 T원소들에 대한 slice이며, 다음 코드는 배열 a의 첫 번째 원소부터 다섯 번째 원소까지의 슬라이스를 생성합니다.

// a[0:5]
// 슬라이스는 배열의 참조와 비슷합니다. 값을 저장하지는 않지만, 슬라이스를 통해 배열의 값에 접근하거나 값을 수정할 수 있습니다. 또 슬라이스의 슬라이스를 만드는 것 또한 가능해 Go언어에서는 배열보다 슬라이스를 더 많이 씁니다.

package main

import "fmt"


func main(){
	names := [4]string{
		"John",
		"Paul",
		"Georage",
		"Ringo",
	}
	fmt.Println("배열 names:", names)
	
	fmt.Println("슬라이스선언")
	var s1 []string = names[0:3] 
	s2 := names[0:2] 

	fmt.Println("names[0:3]", s1)
	fmt.Println("names[0:2]", s2)

	fmt.Println("슬라이스로 값 변경")
	fmt.Println("s1[0]", s1[0])
	s1[0] = "XXX"
	fmt.Println("s1[0] XXX실행 후 ", s1)
	fmt.Println("s1[0] XXX실행 후 ", s2)
	fmt.Println("s1[0] XXX실행 후 ", names)

	s2 = s1[0:2]
	fmt.Println("s2 = s1[0:2]",s2) 

	


}