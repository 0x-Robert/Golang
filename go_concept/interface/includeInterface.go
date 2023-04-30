package main

// Read()와 Close() 메서드를 포함한 Reader 인터페이스
type Reader interface {
	Read() (n int, err error)
	Close() error
}

// Write() 메서드와 Close() 메서드를 포함한 Writer 인터페이스
type Writer interface {
	Write() (n int, err error)
	Close() error
}

type ReadWriter interface {
	Reader //Reader의 메서드 집합을 포함한다.
	Writer //Writer의 메서드 집합을 포함한다.
}

/*
아래와 같은 타입들이 있을 때 어떤 인터페이스로 사용될 수 있는지 살펴보기
Read(), Write(), Close() 메서드를 포함한 타입 >> Reader, WRiter, ReadWriter 모두 사용가능

Read(), Close() 메서드를 포함하는 타입 >> Reader만 사용가능

Write(), Close() 메서드를 포함하는 타입 >> Writer만 사용가능

Read(), Write() 메서드를 포함한 타입 >> Close()메서드가 없기 때문에 Reader, Writer, ReadWriter 모두 사용 불가능

*/
