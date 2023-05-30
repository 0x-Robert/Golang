package main

import (
	"fmt"
)

// 조건문자열
func solution(ineq string, eq string, n int, m int) int {
	switch ineq {
	case "<":
		if eq == "=" {
			if n <= m {
				return 1
			}
		} else if eq == "!" {
			if n < m {
				return 1
			}
		}
	case ">":
		if eq == "=" {
			if n >= m {
				return 1
			}
		} else if eq == "!" {
			if n > m {
				return 1
			}
		}
	}
	return 0
}

func main() {
	result := solution("<", "=", 20, 50)
	fmt.Println(result) // Output: 1
	result = solution(">", "!", 41, 78)
	fmt.Println(result) // Output: 0
}
