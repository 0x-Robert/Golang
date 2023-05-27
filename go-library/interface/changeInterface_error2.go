package main

import "fmt"

type Stringer interface {
	String() string
}

// Student 구조체 선언
type Student struct {
}

// *Student  구조체는 String() 메서드를 가지고 있으므로 Stringer 인터페이스로 사용이 가능하다.
// 덕타이핑!
func (s *Student) String() string {
	return "Student"
}

// Actor 구조체 선언
type Actor struct {
}

// Actor 타입을 리시버로 갖는 메서드 String
func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	//*Student타입은 Stringer 일터페이스로 사용은 가능하다. 그러나
	//stringer 값이 *Actor값으로 들어오기 떄문에 에러발생
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {

	//*Actor 구조체 값을 ConvertType()함수의 인수로 사용
	actor := &Actor{}
	ConvertType(actor)

}
