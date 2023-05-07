// 메소드 선언하기
// 문제 설명
// Go에는 클래스가 없는 대신 리시버 인자를 갖는 함수로 메소드를 정의할 수 있습니다. 리시버는 func 키워드와 메소드 이름 사이에 인자로 들어갑니다. Go언어도 다른 언어와 같이 타입 뒤에 점을 찍어 메소드에 접근합니다.

// func (리시버 인자) 함수이름 리턴타입
// 메소드 호출로 자기 자신의 값을 바꾸려면 인자로 포인터 리시버를 받아야 합니다. 46, 49번째 줄에서 메소드 power10과 power100이 각각 어떻게 다른지 확인해보세요.

// 16, 42번째 줄은 Myfloat과 같은 기본형 타입에 Abs 메소드를 선언해 접근하고 있습니다.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
    X, Y float64
}

//① Abs 메소드는 리시버인자로 v Vertex를 받습니다.
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//② 기본형 타입(여기는 float64)도 메소드를 만들수 있습니다.
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

//③ MyFloat이 포인터가 아닌 리시버 인자입니다
func (f MyFloat) power10() {
    f = f * MyFloat(10)
}

//④ MyFloat이 포인터 리시버 인자입니다.
func (f *MyFloat) power100() {
    *f = *f * MyFloat(100)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println("① 점을 찍어 메소드에 접근합니다")
    fmt.Println("v.Abs():", v.Abs())
    
    f := MyFloat(-math.Sqrt2)
    fmt.Println("②numeric type도 메소드 정의가 가능합니다")
    fmt.Println("f.Abs():", f.Abs())
    
    fmt.Println("③포인터 리시버를 쓰면 메소드 내부에서 값을 바꿀 수 있습니다")
    fmt.Println("기존의 f\t\t\t\t", f)
    f.power10()
    fmt.Println("일반 리시버를 써서 10을 곱한 경우\t", f)
    
    f.power100()
    fmt.Println("포인터 리시버를 써서 100을 곱한 경우\t", f)
}