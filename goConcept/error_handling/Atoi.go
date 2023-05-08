package main

import "fmt"

/*
이유는 Go 언어에서 문자는 UTF-8 인코딩을 사용하며, 문자열에서 각 문자는 해당 문자의 UTF-8 코드 값으로 표현됩니다. 따라서, '3' 문자와 '4' 문자는 각각 UTF-8 코드 값이 51과 52입니다.

Atoi 함수에서 반복문이 첫 번째 실행될 때, str 문자열의 첫 번째 문자 '3'에 대응하는 UTF-8 코드 값인 51이 r 변수에 할당됩니다. 그 후, "r 51" 이라는 메시지가 출력되고, "r2 51" 과 "rst 0" 이라는 메시지가 차례로 출력됩니다. 그리고 "rst += int(r - '0')" 문장이 실행되어 rst 변수가 3으로 변경됩니다.

따라서, "r1 51" 이라는 메시지는 '3'에 대응하는 UTF-8 코드 값이 51임을 나타내며, 이후에는 '4' 문자가 처리되어 rst 변수가 34로 변경되고, 최종적으로 34가 반환되어 출력됩니다.
*/
func Atoi(str string) (int, error) {
	rst := 0
	for _, r := range str {
		fmt.Println("r1", r)
		if r >= '0' && r <= '9' {
			fmt.Println("r", r)
			rst *= 10
			fmt.Println("r2", r)
			fmt.Println("rst", rst)
			rst += int(r - '0')
		} else {
			return 0, fmt.Errorf("숫자만 입력하세요 문자:%c", r)
		}
	}
	fmt.Println("rst", rst)
	return rst, nil
}

func main() {

	//n, err := Atoi("34cd")
	n, err := Atoi("34")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
