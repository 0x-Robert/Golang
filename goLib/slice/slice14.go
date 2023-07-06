package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// 구조체 슬라이스 정렬
// []Student의 별칭 타입인 Students를 만들었다. 
//Len(),Less(),Swap()메서드를 만들어서 sort.Interface를 사용할 수 있게 해줬다. 
type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }

// 나이 비교
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18}}

	sort.Sort(Students(s)) //정렬
	fmt.Println(s)

}
