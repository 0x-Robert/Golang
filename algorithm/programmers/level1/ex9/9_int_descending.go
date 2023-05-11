package level1

import (
	"fmt"
	"sort"
)

// func solution(n int64) int64 {

// 	var answer []int
// 	temp := n

// 	for temp > 0 {
// 		answer = append(answer, int(temp%10)) //10으로 나눈 나머지를 계속해서 answer 배열에 추가
// 		temp /= 10                            //10으로 나눈 몫을 temp 변수에 저장
// 		sort.Sort(sort.Reverse(sort.IntSlice(answer)))
// 	}

// 	var final string
// 	for _, v := range answer {
// 		final += strconv.Itoa(v)
// 	}
// 	final2, _ := strconv.Atoi(final)

// 	return int64(final2)
// }

// func solution2(n int64) int64 {
// 	str := strconv.Itoa(int(n)) //숫자를 스트링으로 변경
// 	fmt.Println("str", str)
// 	result_arr := strings.Split(str, "") //스트링을 전부 분할
// 	fmt.Println("result", result_arr)

// 	sort.Sort(sort.Reverse(sort.StringSlice(result_arr))) //내림차순 정렬?
// 	fmt.Println("result_arr", result_arr)

// 	result, _ := strconv.Atoi(strings.Join(result_arr, "")) //각 흩어진 원소를 하나로 조인
// 	fmt.Println("Result", result)

// 	return int64(result)
// }

func solution9(n int64) int64 {
	arr := make([]int, 0, 16)

	for n != 0 {
		arr = append(arr, int(n%10))
		n /= 10
	}
	fmt.Println("arr", arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("arr2", arr)
	var ret int64

	for _, v := range arr {
		fmt.Println("ret1", ret)
		ret *= 10
		fmt.Println("ret2", ret)
		ret += int64(v)
		fmt.Println("ret3", ret)
	}

	return ret
}
func intDescending() {
	n := 118372
	solution9(int64(n))
}
