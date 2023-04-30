// 문제 설명
// 슬라이스 리터럴은 길이가 없는 배열 리터럴과 같습니다.

// 배열 리터럴은 다음과 같이 표현할 수 있는데요,

// [3]bool{true, true, false}
// 아래 코드는 위와 같은 배열을 만든 후, 그 배열을 참조하는 슬라이스를 만듭니다.

// []bool{true, true, false}
// 슬라이스의 상한과 하한을 지정하지 않을 경우 기본 값으로 설정됩니다. 배열 var a [10]int에 대해 다음 슬라이스는 모두 같은 의미입니다.

// a[0:10]
// a[:10]
// a[0:]
// a[:]
package main

import "fmt"

func main(){
	fmt.Println("슬라이스 리터럴 선언")
	//기본형 슬라이스 리터럴 
	q := []int{2,3,5,7,11,13}
	//구조체 슬라이스 리터럴 
	fmt.Println("기본형 슬라이스 리터럴",q)

	s := []struct{
		i int 
		b bool 
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("구조체 슬라이스 리터럴:", s)
    
    fmt.Println("② 슬라이스를 슬라이스")
    q = q[:2]
    fmt.Println("q[:2]:", q)
    
    q = q[1:]
    fmt.Println("q[1:]:", q)

}