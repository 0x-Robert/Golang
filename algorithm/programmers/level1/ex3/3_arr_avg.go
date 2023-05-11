package main

import "fmt"

func solution3(arr []int) (answer float64) {

	sum := 0.0 // float64 타입으로 초기화
	for _, v := range arr {
		sum += float64(v) // int를 float64로 변환하여 누적 합산
	}
	fmt.Println(sum)
	fmt.Println(len(arr))
	answer = sum / float64(len(arr)) // int를 float64로 변환하여 나누기
	fmt.Println("answer", answer)
	return
}

func main() {
	test := [4]int{1, 2, 3, 4}
	//배열 > 슬라이스 test[:]
	solution3(test[:])
}
