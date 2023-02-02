// 인터페이스는 메소드의 집합으로, 인터페이스 타입 값은 메소드를 구현하는 값을 담을 수 있으며, 타입이 인터페이스에 메소드를 구현하면 자동으로 그 인터페이스도 구현한 게 됩니다. (다른 언어와는 달리 implements등의 키워드가 필요 없습니다)

// 인터페이스 값은 value와 구체적인(concrete) 타입으로 구성된 tuple이라고 볼 수 있습니다.

// type myinterface interface {
//     myfunction() int
// }

// type MyInt int
// func (rcv MyInt) myfunction() int {
//     return 0
// }

// var a myinterface = MyInt(3)

package main

import (
	"fmt"
	"math"
)

type I interface {
    M()
}

type T struct {
    S string
}

//① 별도의 키워드를 쓰지 않아도 T가 인터페이스 I를 구현하게 됩니다.
func (t *T) M() {
    fmt.Println(t.S)
}

type F float64
//② 별도의 키워드를 쓰지 않아도 F가 인터페이스 I를 구현하게 됩니다.
func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I
    
    fmt.Println("① i = &T{\"Hello\"}에 대해")
    i = &T{"Hello"}
    describe(i)
    i.M()
    
    fmt.Println("② i = F(math.Pi)에 대해")
    i = F(math.Pi)
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("인터페이스의 (값, 타입) : (%v, %T)\n", i, i)
}