package main

import "fmt"

func solution(arr []int, query []int) []int {

	for _, v := range query {
		if v%2 == 0 {

			arr = arr[:v+1]

		} else {

			arr = arr[v:]

		}
	}
	fmt.Println("arr", arr)
	return arr
}
func solution2(arr []int, query []int) []int {

	/*
		위의 코드에서는 쿼리를 순회하면서 매번 arr 슬라이스를 변경하고 있다. 이러한 방식은 성능상 문제가 있을 수 있으며,
		arr 슬라이스의 원본 데이터를 보존하는 것도 어렵습니다.
		대신 새로운 슬라이스를 생성하여 원본 데이터를 보존하고, 그 위에서 쿼리에 맞게 조작하는 것이 좋습니다.
	*/

	//차이점
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 0; i < len(query); i++ {
		//차이점
		idx := query[i]

		if i%2 == 0 {
			result = result[:idx+1]
		} else {
			result = result[idx:]
		}
	}
	fmt.Println(result)
	return result
}

//0,1,2,3,4,5
//4,1,2

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	query := []int{4, 1, 2}
	solution2(arr, query)
}
