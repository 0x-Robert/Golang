package main

func solution(num int, total int) []int {

	final := []int{}

	for i := -100; i < total; i++ {
		if total == sum(makeArr(i, i+num)) {
			final = makeArr(i, i+num)
			break
		}
	}
	if num == 1 {
		return []int{total}
	}
	return final
}
func sum(arr []int) int {
	sumCount := 0
	for _, v := range arr {
		sumCount += v
	}
	return sumCount
}

func makeArr(start int, num int) []int {
	arr := []int{}
	for i := start; i < num; i++ {
		arr = append(arr, i)
	}
	return arr
}

// 이렇게 풀수도 있네.. ㄷㄷ.... my head is stone head
func solution0(num int, total int) []int {

	startNumber := (2*total/num + 1 - num) / 2
	answerArray := []int{}

	for i := 0; i < num; i++ {
		answerArray = append(answerArray, startNumber+i)
	}

	return answerArray
}
func main() {
	// num := 3
	// total := 12
	// num := 5
	// total := 5
	// num := 5
	// total := 15
	num := 4
	total := 14
	solution(num, total)
}
