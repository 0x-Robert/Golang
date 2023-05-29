package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(polynomial string) string {

	//공백으로 구분한다.

	//슬라이스로 만든다.

	//순회를 돌면서 x변수 앞의 상수를 계산한다. 앞의 상수가 없으면 1을 더한다.
	//상수계산한 변수를 따로 선언 후 누적한다.
	//순회가 끝나면 x앞 상수계산한 변수 + 정수 변수를 더한 식을리턴한다.

	slice := strings.Split(polynomial, " ")
	count := 0
	cons := 0

	for _, v := range slice {
		mulx, _ := strconv.Atoi(string(v[:len(v)-1]))
		if string(v[len(v)-1]) == "x" {
			if v[:len(v)-1] != "" {
				count = count + mulx

			} else {
				count += 1
			}
		}
		if string(v[len(v)-1]) != "x" && string(v[len(v)-1]) != "+" {
			a, _ := strconv.Atoi(v)
			cons += a
		}

	}
	lcons := strconv.Itoa(cons)
	answer := strconv.Itoa(count)
	if len(lcons) >= 1 {
		answer += "x" + " " + "+" + " " + lcons
	} else {
		answer += "x"
	}
	fmt.Println(answer)
	return answer
}

func solution1(polynomial string) string {
	//공백을 제거
	polynomial = strings.ReplaceAll(polynomial, " ", "")
	fmt.Println(polynomial)
	terms := strings.Split(polynomial, "+")
	//+를 기준으로 구분해서 슬라이스화
	fmt.Println(terms)

	var a, b int
	for _, term := range terms {
		if strings.Contains(term, "x") {
			if len(term) > 1 {
				coefficient, _ := strconv.Atoi(term[:len(term)-1])
				a += coefficient
			} else {
				a += 1
			}
		} else {
			coefficient, _ := strconv.Atoi(term)
			b += coefficient
		}
	}

	if a == 0 {
		return fmt.Sprintf("%d", b)
	} else if a == 1 {
		if b == 0 {
			return "x"
		} else {
			return fmt.Sprintf("x + %d", b)
		}
	} else if a > 1 {
		if b == 0 {
			return fmt.Sprintf("%dx", a)
		} else {
			return fmt.Sprintf("%dx + %d", a, b)
		}
	}

	return ""
}

func main() {

	polynomial := "3x + 7 + 16x + 1334"
	//polynomial := "x + x + x"
	solution1(polynomial)
}
