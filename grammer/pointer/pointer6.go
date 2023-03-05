package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	// var u = User(name, age)
	// //이미 사라진 주소를 반환하면 에러가 나온다.
	// return &u

	return &User{name, age}
}

func main() {
	// 다음은 무효
	// 탈출 분석 >> 스택 또는 힙에 만들지 결정함
	// 탈출하는 놈은 스택에다 만들면 안됨 그래서 힙에다 만들음
	// 힙에 있는 거는 쓰임이 다하면 사라짐, 쓰임이 있으면 사라지지 않음
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
