package main

import "fmt"

//다음 코드에서 인스턴스는 총 몇개?

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

func main() {
	//다음 코드에서 인스턴스 한개가 생성됨
	//나머지는 모두 *User 타입 포인터 변수로 만들어진 인스턴스를 가리킬 뿐
	newUser := NewUser("AAA", 23)
	var p *User = newUser
	p.Age += 10

	fmt.Println(p.Age)
}
