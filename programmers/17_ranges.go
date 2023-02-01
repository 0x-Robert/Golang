// 문제 설명
// 반복문에서 range문을 쓰면 슬라이스나 맵을 이터레이트할 수 있습니다.
// 슬라이스에서 range를 쓸 경우, 인덱스, 원소 값이 매 순회마다 리턴됩니다. 이때 인덱스나 원소값이 필요 없다면 변수를 지정하지 않고 _를 쓸 수 있습니다.

package main

import "fmt"


var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main(){
	//일반적인 range
	fmt.Println("일반적인 range")
	for i, v := range pow{
		fmt.Printf("2**%d = %d\n", i, v)

	}

	//인덱스가 필요 없는 경우 _로 비워둘 수 있음
	fmt.Println("2인덱스가 필요없는 경우")
	for _, v := range pow {
		fmt.Println(v)
	}
}