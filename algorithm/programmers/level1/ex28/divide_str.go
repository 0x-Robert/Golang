package main

import "fmt"

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
	// fmt.Println("count", count)
	return count
}

func main() {

	//s := "banana"
	//s := "abracadabra"
	s := "aaabbaccccabba"
	solution(s)
}
