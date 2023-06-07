package main

import (
	"strconv"
	"strings"
)

func solution(s string) string {

	answer := []string{}

	slice := strings.Split(s, " ")
	min, _ := strconv.Atoi(slice[0])
	max, _ := strconv.Atoi(slice[0])
	for _, v := range slice {
		num, _ := strconv.Atoi(v)
		if max < num {
			max = num
		} else if min > num {
			min = num
		}
	}

	answer = append(answer, strconv.Itoa(min))
	answer = append(answer, strconv.Itoa(max))
	answer2 := strings.Join(answer, " ")
	//fmt.Println(answer2)
	return answer2
}

func solution0(s string) string {
	strs := strings.Fields(s)

	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	sort.Ints(ints)

	return fmt.Sprintf("%d %d", ints[0], ints[len(ints)-1])
}

func main() {
	s := "1 2 3 4"
	solution(s)
}
