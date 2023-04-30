package main

import "fmt"

type Student struct {
	name  string
	age   int
	group *Group
}

var student Student

func main() {
	fmt.Println("가비지 컬렉터 알고리즘에는 다음과 같은 방법이 있다.")
	fmt.Println("표시하고 지우기, 블록을 검사해서 사용하면 1 사용안하면 0, 0으로 표시된 모든 메모리 블록을 삭제함")
	fmt.Println("삼색표시, 회색은 아직 검사안한 메모리, 흰색은 아무도 사용하지 않는 방식, 검은색은 이미 검사가 끝낸 블록 >> go언어의 가비지콜렉터 방식, 고루틴으로 동작, 장점:매우짧은 멈춤시간, 단점:추가 힙 메모리 필요, 실행성능이 저하될 수 있음")
	fmt.Println("객체위치이동 삭제할 메모리를 표시한 뒤 한쪽으로 몰아서 삭제하는 방식")
	fmt.Println("세대단위수집")

}
