package main

import "fmt"

func longestPalindrome(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLength := 1
	start := 0
	fmt.Println(start)
	for i := 1; i < len(s); i++ {
		// 팰린드롬 중심이 한 글자일 때
		left := i - 1
		right := i + 1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length := right - left + 1
			if length > maxLength {
				maxLength = length
				start = left
			}
			left--
			right++
		}

		// 팰린드롬 중심이 두 글자일 때
		left = i - 1
		right = i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			length := right - left + 1
			if length > maxLength {
				maxLength = length
				start = left
			}
			left--
			right++
		}
	}

	return maxLength
}
func main() {
	str := "My dad is a racecar athlete"
	result := longestPalindrome(str)
	fmt.Println("result", result)

	str2 := " dad "
	result2 := longestPalindrome(str2)
	fmt.Println("result", result2)
}
