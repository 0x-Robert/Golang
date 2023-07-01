// 배열 회전시키기
// 문제 설명
// 정수가 담긴 배열 numbers와 문자열 direction가 매개변수로 주어집니다. 배열 numbers의 원소를 direction방향으로 한 칸씩 회전시킨 배열을 return하도록 solution 함수를 완성해주세요.

// 제한사항
// 3 ≤ numbers의 길이 ≤ 20
// direction은 "left" 와 "right" 둘 중 하나입니다.
// 입출력 예
// numbers	direction	result
// [1, 2, 3]	"right"	[3, 1, 2]
// [4, 455, 6, 4, -1, 45, 6]	"left"	[455, 6, 4, -1, 45, 6, 4]
// 입출력 예 설명
// 입출력 예 #1

// numbers 가 [1, 2, 3]이고 direction이 "right" 이므로 오른쪽으로 한 칸씩 회전시킨 [3, 1, 2]를 return합니다.
// 입출력 예 #2

// numbers 가 [4, 455, 6, 4, -1, 45, 6]이고 direction이 "left" 이므로 왼쪽으로 한 칸씩 회전시킨 [455, 6, 4, -1, 45, 6, 4]를 return합니다.

package main

import "fmt"

func solution(numbers []int, direction string) []int {
	// numbers = append(number, numbers...)

	if direction == "right" {
		numbers = append(numbers, numbers[0])
		fmt.Println(numbers)
		//	return numbers[(len(numbers)/2)-1 : (len(numbers)/2)+2]
		return numbers
	} else {
		// fmt.Println(numbers[1 : (len(numbers)/2)+1])
		//	return numbers[1 : (len(numbers)/2)+1]
		return []int{}
	}
}

func solution1(numbers []int, direction string) []int {
	length := len(numbers)
	result := make([]int, length)

	if direction == "left" {
		// Rotate left
		result[length-1] = numbers[0]
		copy(result[:length-1], numbers[1:])
	} else if direction == "right" {
		// Rotate right
		result[0] = numbers[length-1]
		copy(result[1:], numbers[:length-1])
	}

	return result
}

func solution2(numbers []int, direction string) []int {
	length := len(numbers)
	answer := make([]int, 0, length)
	if direction == "left" {
		answer = append(answer, numbers[1:length]...)
		answer = append(answer, numbers[0])
	} else {
		answer = append(answer, numbers[length-1])
		answer = append(answer, numbers[0:length-1]...)
	}
	return answer
}

func main() {
	numbers := []int{1, 2, 3}
	direction := "right"

	// numbers := []int{4, 455, 6, 4, -1, 45, 6}
	// direction := "left"
	solution(numbers, direction)
}
