package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {

}

func ReadFile(reader Reader) {
	//Reader 인터페이스 변수를 Closer인터페이스로 타입 변환합니다.
	//런 타임 에러가 발생한다.
	c, ok := reader.(Closer)
	if ok {
		c.Close()

	}
	fmt.Println("ReadFile")

}

func main() {
	file := &File{}
	ReadFile(file)
}

/*
인터페이스는 메서드 집합체
인터페이스에서 정의한 메서드 집합을 가진 모든 타입은 인터페이스로 쓰일 수 있다.
덕 타이핑이란 인터페이스 구현 여부를 명시적으로 선언하는게 아닌 인터페이스에서 정의한 메서드 포함 여부로 판단합니다.
인터페이스를 사용해 추상화 계층을 만들고 관계를 통한 상호작용을 정의합니다.
모든 타입이 빈 인터페이스 변숫값으로 쓰일 수 있다.
인터페이스 변환을 사용하면 인터페이스 변수를 구체화된 타입이나 다른 인터페이스로 변경할 수 있다.

*/
