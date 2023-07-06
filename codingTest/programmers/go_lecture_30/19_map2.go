// 문제 설명
// 맵을 조작하는 방법에 대해 알아봅시다.

// 1. 맵에 원소를 추가하려면
// m[key] = elem

// 2. 맵에서 특정 키 값을 가져오려면
// elem = m[key]

// 3. 맵에 원소를 제거하려면
// delete(m, key)

// 4. 맵에 키가 존재하는지 확인하려면
// elem, ok = m[key]

// m에 key가 있으면 ok가 true이고, 없으면 false입니다.

package main

import "fmt"

func main(){

	m := make(map[string]int)

	m["Answer"] = 42 
	fmt.Println("m[\"Answer\"]값은:", m["Answer"]) 

	delete(m, "Answer")
	fmt.Println("m[\"Answer\"]값은:", m["Answer"])

	v, ok  := m["Answer"]
	fmt.Println("m[\"Answer\"]값은:", v, "존재하나요?", ok)

	
}