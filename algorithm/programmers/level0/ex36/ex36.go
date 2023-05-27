package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func solution(quiz []string) []string {
// 	answer := []string{}

// 	for _, v := range quiz {
// 		v := strings.ReplaceAll(v, " ", "")
// 		split := regexp.MustCompile("=").Split(v, -1)
// 		num := regexp.MustCompile("[0-9]+")
// 		operatorPattern := `[+-]`
// 		operatorReg := regexp.MustCompile(operatorPattern)
// 		leftOp := operatorReg.FindAllString(split[0], -1)
// 		leftV := num.FindAllString(split[0], -1)
// 		rightOp := operatorReg.FindAllString(split[1], -1)
// 		rightV := num.FindAllString(split[1], -1)

// 		// fmt.Println(leftOp[0])     //왼쪽 연산자
// 		// fmt.Println(split[0])      //왼쪽 항들
// 		// fmt.Println("0", leftV[0]) //왼쪽 값 첫 번째
// 		// fmt.Println("1", leftV[1]) //왼쪽 값 두 번째
// 		// fmt.Println(split[1])      //오른쪽 항
// 		// fmt.Println(len(split))    //배열 길이
// 		// fmt.Println("split", split)
// 		// fmt.Println("a,b, leftOp,leftV[0],leftV[1], rightOp, rightV", a, b, leftOp[0], leftV[0], leftV[1], rightOp, rightV)
// 		a, _ := strconv.Atoi(leftV[0])
// 		b, _ := strconv.Atoi(leftV[1])

// 		switch len(rightOp) {
// 		case 1:
// 			rightV, _ := strconv.Atoi(rightV[0])
// 			rightV = -rightV
// 			switch leftOp[0] {
// 			case "+":
// 				if a+b == rightV {
// 					answer = append(answer, "O")
// 				} else {
// 					answer = append(answer, "X")
// 				}
// 			case "-":
// 				if a-b == rightV {
// 					answer = append(answer, "O")
// 				} else {
// 					answer = append(answer, "X")
// 				}
// 			}
// 		case 0:
// 			rightV, _ := strconv.Atoi(rightV[0])
// 			rightV = +rightV

// 			switch leftOp[0] {
// 			case "+":
// 				if a+b == rightV {
// 					answer = append(answer, "O")
// 				} else {
// 					answer = append(answer, "X")
// 				}
// 			case "-":
// 				if a-b == rightV {
// 					answer = append(answer, "O")
// 				} else {
// 					answer = append(answer, "X")
// 				}
// 			}
// 		}

// 	}
// 	//fmt.Println("answer", answer)

//		// =을 기준으로 문자열을 나눈다.
//		//나눈 문자열에서 숫자와 기호를 뽑아내서 계산한다.
//		//우항과 좌항의 비교를 통해 O or X를 배열에 넣는다.
//		//배열을 리턴한다.
//		return answer
//	}
func solution(quiz []string) []string {
	answer := make([]string, len(quiz))

	for i, q := range quiz {
		equation := strings.Split(q, " ")
		x, _ := strconv.Atoi(equation[0])
		y, _ := strconv.Atoi(equation[2])
		z, _ := strconv.Atoi(equation[4])

		operator := equation[1]
		result := 0

		if operator == "+" {
			result = x + y
		} else if operator == "-" {
			result = x - y
		}

		if result == z {
			answer[i] = "O"
		} else {
			answer[i] = "X"
		}
	}
	fmt.Println(answer)
	return answer
}

func main() {
	//quiz := []string{"3 - 4 = -3", "5 + 6 = 11"}
	quiz := []string{"19 - 6 = 13", "5 + 66 = 71", "5 - 15 = 63", "3 - 1 = 2"}
	solution(quiz)
}
