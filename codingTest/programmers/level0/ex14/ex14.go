package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Scan(&s1)

	converted := convertString(s1)
	fmt.Println(converted)
}

func convertString(input string) string {

	converted := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 32
		} else if r >= 'A' && r <= 'Z' {
			return r + 32
		}
		return r
	}, input)
	return converted
}
