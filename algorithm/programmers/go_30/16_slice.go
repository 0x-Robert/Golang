// 문제 설명
// 슬라이스 s에 대해 슬라이스의 length와 capacity는 다음과 같이 표현할 수 있습니다.

// length : len(s) - 슬라이스가 포함하고 있는 원소의 개수
// capacity : cap(s) - 슬라이스가 가리키는 배열에서 슬라이스의 첫 번째 원소부터 센 원소 개수
// 슬라이스의 zero value는 nil이고, 가리키는 배열이 없으며 length와 capacity 모두 0입니다.

// 다음 코드에서 capacity가 어떻게 증가하고 있는지 확인해보세요

// make함수로 0으로 초기화된 배열을 생성하고, 이를 가리키는 슬라이스를 리턴 받아 가변 길이 배열을 만들어낼 수 있습니다. 새 원소를 추가할 때는 func append(s []T, vs ...T) []T를 씁니다. append함수의 첫 번째 인자에 타입 T 슬라이스 s를 넣고, 그다음으로 추가할 T값(들)을 전해주면 값이 추가된 슬라이스를 리턴합니다.

package main

import "fmt"

func main() {
    // make()로 가변 길이 배열 만들기
    a := make([]int, 5)
    fmt.Printf("a := make([]int, 5)의\t")
    printSlice(a)
    
    b := make([]int, 0, 5)
    fmt.Printf("b := make([]int, 0, 5)의\t")
    printSlice(b)
    
    c := b[:2]
    fmt.Printf("c := b[:2]의\t\t")
    printSlice(c)
    
    d := c[2:5]
    fmt.Printf("d := c[2:5]의\t\t")
    printSlice(d)
    
    // 한번에 여러개 원소를 추가할 수 있습니다.
    d = append(d, 1,2,3)
    fmt.Printf("d = append(d, 1,2,3)후\t")
    printSlice(d)
    
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}