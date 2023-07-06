package main

func solution(money int) []int {

	//아이스 아메리카노는 5500원이고 가진 돈으로 아메리카노를 사먹을 수 있는 잔과 잔돈을 배열에 담아서 리턴
	a := money / 5500
	b := money % 5500
	answer := []int{}
	answer = append(answer, a)
	answer = append(answer, b)
	//fmt.Println(answer)
	return answer
}

// ....
func solution0(money int) []int {
	return []int{money / 5500, money % 5500}
}

func main() {
	// money := 5500
	money := 15000
	solution(money)
}
