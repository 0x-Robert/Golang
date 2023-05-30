import (
	"strings"
	"strconv"
	"fmt"
)

func solution(polynomial string) string {
	arr := strings.Split(polynomial, " + ")

	s := make([]int, 2)
	for _, v := range arr {
		str := v
		if strings.Index(str, "x") == -1 {
			i, _ := strconv.Atoi(str)
			s[1] += i
		} else {
			str = strings.Replace(str, "x", "", -1)
			var i int
			if 0 >= len(str) {
				i = 1
			} else {
				i, _ = strconv.Atoi(str)
			}

			s[0] += i
		}
	}

	var str string

	if 0 < s[0] {
		if 1 == s[0] {
			str = "x"
		} else {
			str = fmt.Sprintf("%vx", s[0])
		}
	}

	if 0 < s[1] {
		if 0 < len(str) {
			str = fmt.Sprintf("%v + ", str)
		}

		str = fmt.Sprintf("%v%v", str, s[1])
	}

	return str
}
