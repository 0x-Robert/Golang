package main

// 코드 처리하기
// mode는 0과 1이 있다.
// idx를 0부터 code-1길이까지 1씩 순회하면서 동작한다.
func solution(code string) string {
	mode := 0
	ret := ""

	for idx := range code {
		// fmt.Println("v", stringWW(v))

		if mode == 0 {
			// fmt.Println(string(code[idx]))
			if string(code[idx]) != "1" {
				if idx%2 == 0 {
					ret += string(code[idx])
				}
			} else if string(code[idx]) == "1" {
				mode = 1
			}
		} else if mode == 1 {
			// fmt.Println(string(code[idx]))
			if string(code[idx]) != "1" {
				if idx%2 != 0 {
					ret += string(code[idx])
				}
			} else if string(code[idx]) == "1" {
				mode = 0
			}
		}
	}
	// fmt.Println("ret", ret)
	if len(ret) == 0 {
		return "EMPTY"
	}
	return ret
}

func main() {
	code := "abc1abc1abc"
	// result := "acbac"
	solution(code)
}
