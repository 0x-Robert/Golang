package main

import (
	"fmt"
	"unicode"
)

func solution(s string) string {

	slice := []rune(s)
	index := 0
	for i, v := range slice {
		if unicode.IsSpace(v) {
			index = i + 1
		} else if (i-index)%2 == 0 {
			slice[i] = unicode.ToUpper(slice[i])
			fmt.Println("unicode.ToUpper(slice[i])", unicode.ToUpper(slice[i]))
			fmt.Println("unicode.ToUpper(slice[i])", string(unicode.ToUpper(slice[i])))

		} else {
			slice[i] = unicode.ToLower(slice[i])
		}
	}

	return string(slice)
}

func main() {
	s := " try   hello   world "
	solution(s)
}
