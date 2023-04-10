package main

import "fmt"

func main() {
	fmt.Println("문자열 타입")
	//문자열 타입
	var str = "hello 월드"
	for i, c := range str {
		//i는 인덱스, c는 문자값
		fmt.Println(i, c)
	}
	fmt.Println("str", str)

	fmt.Println("슬라이스 타입")
	//슬라이스 타입
	var slice = []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Println("slice", slice)

	//맵 타입
	fmt.Println("맵타입")
	var m = map[string]int{"aaa": 1, "bbb": 2, "ccc": 3}
	for k, v := range m {
		//k는 키, v는 키에 해당하는 값
		fmt.Println(k, v)
	}
	fmt.Println("m", m)

	fmt.Println("채널 타입")
	//채널타입
	var ch := make(chan int)
	for v:= range ch {
		//계속 채널에 값이 들어올때까지 들어온 값을 반환합니다. 
	}
}
