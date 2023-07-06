package main

import "strings"

func solution1(s string) int {
	characters := strings.Split(s, "")
	target := ""
	count, same_count, not_same_count := 0, 0, 0

	for _, char := range characters {
		if target == "" {
			target = char
		}

		if char == target {
			same_count += 1
		} else {
			not_same_count += 1
		}

		if same_count == not_same_count {
			count += 1
			target = ""
			same_count, not_same_count = 0, 0
		}
	}
	if same_count != not_same_count {
		count += 1
	}
	return count
}

func solution(s string) int {

	x := int32(s[0])
	isX := []string{}
	isNotX := []string{}
	count := 0

	for _, v := range s {
		if len(isX) == len(isNotX) && len(isX) != 0 && len(isNotX) != 0 {
			count += 1
			isX = []string{}
			isNotX = []string{}
			isX = append(isX, string(v))
			x = v
		} else if x == v {
			isX = append(isX, string(v))
		} else {
			isNotX = append(isNotX, string(v))
		}
	}

	if len(isX) > 0 {
		count += 1
	}

	return count
}

func main() {

	//s := "banana"
	//s := "abracadabra"
	s := "aaabbaccccabba"
	solution(s)
}
