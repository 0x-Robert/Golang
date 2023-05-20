package main

import (
	"math/big"
)

// func solution(a string, b string) string {

// 	if a == "0" && b == "0" {
// 		return "0"
// 	}
// 	var answer int64
// 	a1, _ := strconv.ParseInt(a, 10, 64)
// 	b1, _ := strconv.ParseInt(b, 10, 64)
// 	answer = a1 + b1
// 	fmt.Println("strconv.Itoa(answer)", strconv.FormatInt(answer, 10))
// 	return strconv.FormatInt(answer, 10)
// }

func solution(a string, b string) string {

	if a == "0" && b == "0" {
		return "0"
	}

	a1 := new(big.Int)
	b1 := new(big.Int)
	a1.SetString(a, 10)
	b1.SetString(b, 10)

	answer := new(big.Int)
	answer.Add(a1, b1)

	//fmt.Println("answer.String()", answer.String())
	return answer.String()
}

func main() {
	// a := "582"
	// b := "734"
	a := "18446744073709551615"
	b := "287346502836570928366"
	solution(a, b)
}
