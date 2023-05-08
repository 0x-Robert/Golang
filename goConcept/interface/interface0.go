package main

import "fmt"

//인터페이스 선언
type Stringer interface {
	String() string
}

//구조체 선언 
type Student struct {
	Name string
	Age  int
}

//메서드를 선언, Student의 String() 메서드
func (s Student) String() string {
	return fmt.Sprintf("안녕 나는 %d살 %s라고 해", s.Age, s.Name)
	//문자열 만들기 

}


func main() {
	student := Student{"용가리", 30} //Student타입, 구조체  
	var stringer Stringer //Stringer 타입, 인터페이스 

	stringer = student //stringer 값으로 student 대입 

	fmt.Printf("%s\n", stringer.String()) //stringer의 String() 메서드 호출 
}
