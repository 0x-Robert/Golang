package level1

import (
	"strconv"
)

func solution4(n int) (answer int) {

	str_num := strconv.Itoa(n) //숫자를 문자열로 변환

	for _, v := range str_num {

		num, _ := strconv.Atoi(string(v)) //문자열 변환 후 정수로 파싱
		answer += num
	}
	//fmt.Println("answer", answer)
	return answer
}

// 즉, n /= 10은 n 변수의 일의 자리 숫자를 없애는 효과가 있습니다.
// 이와 같은 코드는 주어진 수를 일의 자리부터 한 자리씩 떼어내어 계산할 때 자주 사용됩니다.
// func solution1(n int) int {
// 	answer := 0

// 	for {
// 		fmt.Println("before answer", answer)
// 		fmt.Println("n%10", n%10)
// 		answer += n % 10 // 10으로 나눈 나머지
// 		fmt.Println("answer1", answer)
// 		fmt.Println("before  n", n)
// 		n /= 10 //나눗셈의 몫만 취함
// 		fmt.Println("n", n)
// 		if n == 0 {
// 			break
// 		}
// 	}
// 	fmt.Println("answer", answer)
// 	return answer
// }

func sumSolution() {
	N := 123
	solution4(N)
}
