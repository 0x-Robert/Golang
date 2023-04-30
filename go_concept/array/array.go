package main

import (
	"fmt"
)

func main() {
	//초기화
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	var nums [5]int
	nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	days := [3]string{"mon", "tue", "wed"}
	fmt.Println(days)

	var temps [5]float64 = [5]float64{24.3, 26.7}

	var s = [5]int{1: 10, 3: 30}
	//...세개는 길이를 뒤에있는걸로 정하겠다. 이런 뜻 그러면 ...은 길이가 3이됨
	x := [...]int{10, 20, 30}
	fmt.Println(temps, s, x)

	//배열 선언시 개수는 항상 상수
	const Y int = 3
	x1 := 5
	//a := [x]int{1, 2, 3, 4, 5} //변수를 길이로 하면 에러남

	b := [Y]int{1, 2, 3}

	var c [10]int
	fmt.Println(x1, b, c)

	nums2 := [...]int{10, 20, 30, 40, 50}
	nums2[2] = 300

	// for i := 0; i < len(nums2); i++ {
	// 	fmt.Println(nums2[i])
	// }

	var t2 [5]float64 = [5]float64{24.0, 25.9, 26.5, 27.5, 26.3}

	for i, v := range t2 {
		fmt.Println(i, v)
	}

	// 배열은 연속된 메모리
	// 컴퓨터는 인덱스와 타입 크기를 사용해서 메모리 주소를 찾는다.....
	// 배열은 요소를 찾아가는게 빠르다.
	// 배열이 리스트나 맵, 트리 중 랜덤 엑세스면에서 가장 빠르다.
	// 두서없이 인덱스 접근할 때는 배열을 선택하는 것이 좋다.

	//배열 복사
	a1 := [5]int{1, 2, 3, 4, 5}
	b1 := [5]int{500, 400, 300, 200, 100}

	for i, v := range a1 {
		fmt.Printf("a1[%d] = %d\n", i, v)
	}
	fmt.Println()
	for i, v := range b1 {
		fmt.Printf("b1[%d] = %d\n", i, v)
	}
	fmt.Println()

	//golang은 타입과 크기가 같아야한다. !

	b1 = a1
	fmt.Println()
	for i, v := range b1 {
		fmt.Printf("b1[%d] = %d\n", i, v)
	}
	fmt.Println()

	//이중배열까지는 많이 쓰임, 다중배열은 많이 안쓰임
	// var b [2][5] int
	// [2][5][100][200] int

	//이중배열 순회
	a3 := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
	}
	for _, arr := range a3 {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	

}
