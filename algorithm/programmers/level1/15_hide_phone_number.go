package main

import "strings"

func solution(phone_number string) string {

	var result string
	starNum := len(phone_number[:len(phone_number)-4])
	for i := 0; i < starNum; i++ {

		result += "*"
	}
	result += phone_number[len(phone_number)-4:]
	return result
}

func solution2(phone_number string) string {
	return strings.Repeat("*", len(phone_number)-4) + phone_number[len(phone_number)-4:]
}

func main() {
	phone_number := "01033334444"
	solution(phone_number)
}
