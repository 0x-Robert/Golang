package main

import (
	"fmt"
	"strings"
)

// 작은 따옴표: 작은 따옴표는 문자(character) 리터럴을 나타냅니다. 즉, 하나의 문자를 표현하는 데 사용됩니다. 예를 들어, 'a', 'b', 'c'와 같은 문자 리터럴은 모두 작은 따옴표로 둘러싸여 있습니다. 작은 따옴표는 유니코드 코드 포인트를 나타내는 rune 타입의 값으로 해석됩니다.

// 큰 따옴표: 큰 따옴표는 문자열(string) 리터럴을 나타냅니다. 즉, 하나 이상의 문자를 포함하는 문자열을 표현하는 데 사용됩니다. 예를 들어, "Hello, world!"와 같은 문자열 리터럴은 큰 따옴표로 둘러싸여 있습니다. 큰 따옴표는 내부적으로 바이트(byte) 슬라이스를 나타내는 []byte 타입의 값으로 해석됩니다.

// 따라서, 작은 따옴표는 하나의 문자를, 큰 따옴표는 하나 이상의 문자열을 나타냅니다. Go 언어에서는 문자열과 문자를 구분하여 사용하며, 작은 따옴표를 사용하는 경우에는 반드시 유니코드 코드 포인트를 하나만 포함해야 합니다.
// Go 언어의 strings.Builder 타입의 WriteRune 메서드는 rune 타입의 문자를 Builder에 추가하는 함수입니다. 이 메서드는 Writer 인터페이스를 구현하는 메서드 중 하나이며, io.WriteString() 함수와 함께 문자열을 처리하는 데 사용됩니다.

// strings.Builder의 WriteRune 메서드는 UTF-8 인코딩을 사용하며, rune 타입의 인자로 유효한 Unicode 코드 포인트를 받습니다. 이 메서드는 주어진 rune 값을 UTF-8으로 인코딩하여 Builder 버퍼에 추가합니다.

// WriteRune 메서드는 단일 문자만 처리하므로, 여러 문자를 처리하려면 반복문을 사용하거나 strings.Builder의 다른 메서드를 사용해야 합니다.
func ToUpper(str string) string {
	var builder strings.Builder

	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			fmt.Println("if")
			builder.WriteRune('A' + (v - 'a'))
			//fmt.Println(builder.WriteRune('A' + (v - 'a')))
		} else {

			fmt.Println("else")
			builder.WriteRune(v)
		}
	}
	// 	위 코드에서 builder.WriteRune('A' + (v - 'a')) 는 입력 문자열에서 소문자인 경우 해당 문자를 대문자로 변경하기 위한 코드입니다.

	// 여기서 'A' + (v - 'a')는 아스키 코드에서 대문자 'A'와 해당 소문자 문자 간의 차이를 더해주는 것으로, 소문자 'a'부터 'z'까지 순회할 때마다 해당 소문자에 대응되는 대문자를 계산합니다. 예를 들어, 'a'에 대응되는 대문자는 'A'이므로, 'A' + ('a' - 'a')는 'A'와 같습니다. 'b'에 대응되는 대문자는 'B'이므로, 'A' + ('b' - 'a')는 'B'와 같습니다.

	// 이러한 방법을 사용하여 대소문자를 구분하지 않고 문자열을 처리할 수 있습니다. 따라서 위 코드의 결과는 "HELLO YONGARI 22"와 같이 모든 소문자를 대문자로 변경한 문자열이 됩니다.
	return builder.String()
}

func main() {
	str := "hello yongari 22"
	fmt.Println(ToUpper(str))
	var builder strings.Builder
	fmt.Println(builder.WriteRune('A' + ('b' - 'a')))
}
