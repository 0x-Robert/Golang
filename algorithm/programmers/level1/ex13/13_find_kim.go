package main

import (
	"strconv"
)

func solution13(seoul []string) string {

	var answer string
	for i, v := range seoul {
		if v == "Kim" {
			answer += "김서방은 " + strconv.Itoa(i) + "에 있다"
			//return fmt.Sprintf("김서방은 %d에 있다", i)
		}
	}

	return answer
}

func main() {
	arr := []string{"Jane", "Kim"}
	solution13(arr)
}
