package main

import "fmt"

type Stringer interface {
	//인터페이스
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string { //Student 타입의 String() 메서드
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15} //*Student 타입 변수 s 선언 및 초기화
	PrintAge(s)       //변수 s를 인터페이스 인수로 PrintAge() 함수 호출
}
