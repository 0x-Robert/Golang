package main

import (
	"fmt"
	"sort"
	"strconv"
)

func cmp(a, b string) bool {
	if a+b > b+a {
		// fmt.Println("a+b",a+b, "b+a",b+a)
		return true
	} else {
		return false
	}
}

func solution(numbers []int) string {
	//슬라이스 선언
	strArr := make([]string, len(numbers))


	for i, num := range numbers {
		//string으로 변경
		strArr[i] = strconv.Itoa(num)
	}


	sort.Slice(strArr, func(i, j int) bool {
		return cmp(strArr[i], strArr[j])
		
	})

	answer := ""
	for _, str := range strArr {
		answer += str
	}
	if answer[0] == '0' {
		return "0"
	}
	fmt.Println(answer)
	return answer
}


func solution1(numbers []int) string {
    sort.Slice(numbers, func(i, j int) bool {
        s1 := strconv.Itoa(numbers[i])
        s2 := strconv.Itoa(numbers[j])
        return s1+s2 > s2+s1
    })

    if numbers[0] == 0 {
        return "0"
    }
    answer := ""
    for i := range numbers {
        answer += strconv.Itoa(numbers[i])
    }
    return answer
}


func main() {
	//numbers := []int{6, 10, 2}
	numbers2 := []int{3, 30, 34, 5, 9}
	solution(numbers2)
}