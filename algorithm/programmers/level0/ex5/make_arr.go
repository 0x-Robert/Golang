package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
regexp.QuoteMeta(chars)는 chars 문자열에 포함된 메타문자를 이스케이프 처리하여 정규식 패턴을 만듭니다.
예를 들어 chars에 ^ 문자가 있으면 이 문자는 정규식에서 문자열의 시작 부분을 나타내는 메타문자이므로 이스케이프 처리가 필요합니다.

"^[" + pattern + "]+$"은 chars 문자열에 포함된 문자들로 이루어진 문자열을 매칭하는 정규식 패턴을 만듭니다.

	^는 문자열의 시작을 의미하며, ]는 chars 문자열의 마지막 문자를 의미하고, +는 해당 문자열이 한 번 이상 나타나는 것을 의미합니다. $는 문자열의 끝을 의미합니다.

	regexp.MustCompile(pattern)은 위에서 만든 정규식 패턴을 컴파일하여 사용할 수 있는 정규식 객체를 생성합니다.
	이렇게 생성된 정규식 객체를 이용해 re.MatchString(d) 메소드를 호출하여 d 문자열이 패턴에 매칭되는지 확인합니다.
	매칭되면 true를 반환하고, 매칭되지 않으면 false를 반환합니다.
*/
func isDigitisMadeOf(d string, chars string) bool {
	pattern := "^["
	pattern += regexp.QuoteMeta(chars)
	pattern += "]+$"
	re := regexp.MustCompile(pattern)
	return re.MatchString(d)
}

func solution1(l int, r int) []int {
	digits := "05"

	arr := []int{}
	for i := l; i <= r; i++ {
		//fmt.Println("l", i)
		if isDigitisMadeOf(strconv.Itoa(i), digits) {
			arr = append(arr, i)
		}

	}

	if len(arr) == 0 {

		arr = append(arr, -1)
	}
	fmt.Println("arr", arr)
	return arr
}

func check(num int) bool {
	str := strconv.Itoa(num)
	for _, r := range str {
		if r != '5' && r != '0' {
			return false
		}
	}
	return true
}

func solution0(l int, r int) []int {
	result := []int{}
	for i := l / 5 * 5; i <= r; i += 5 {
		if i >= l && i <= r && check(i) {
			result = append(result, i)
		}
	}
	if len(result) == 0 {
		return []int{-1}
	}
	return result
}

func main() {
	l := 5
	r := 555
	//r := 10
	solution1(l, r)
	//solution0(l, r)
}
