package level1

import "strings"

// 숫자 0~9 / 48 ~ 57
// A : 65~97
// a :
// func solution(s string) bool {
// 	if len(s) == 4 || len(s) == 6 {
// 		for _, v := range s {

// 			if v < 48 || v > 57 {

// 				return false
// 			}
// 		}

// 		return true
// 	} else {

// 		return false
// 	}

// }

func solution19(s string) bool {
	return (len(s) == 4 || len(s) == 6) && strings.ToUpper(s) == s && strings.ToLower(s) == s
}

func manageStr() {
	//s := "23d2ef"
	s := "123453"
	solution19(s)
}
