package main

import (
	"sort"
	"strconv"
)

func solution(n int64) int64 {

	var answer []int
	temp := n

	for temp > 0 {
		answer = append(answer, int(temp%10)) //10으로 나눈 나머지를 계속해서 answer 배열에 추가
		temp /= 10                            //10으로 나눈 몫을 temp 변수에 저장
		sort.Sort(sort.Reverse(sort.IntSlice(answer)))
	}

	var final string
	for _, v := range answer {
		final += strconv.Itoa(v)
	}
	final2, _ := strconv.Atoi(final)

	return int64(final2)
}

func main() {
	n := 118372
	solution(int64(n))
}
