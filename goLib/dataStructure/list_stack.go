package main

import (
	"container/list"
	"fmt"
)

//큐는 리스트로 스택은 배열로? 구현을 많이한다.
//스택은 배열로 만들어도 기능손상이 없는 이유는 스택은 요소 맨 뒤에서 삽입과 삭제를 하기 떄문

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}

}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
		fmt.Println("&stack", &stack)
		fmt.Println("stack", stack)
		fmt.Println("*stack", *stack)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		fmt.Println("stack", stack)
		val = stack.Pop()
	}

}
