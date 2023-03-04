package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceString := []string{"a", "d", "c", "b"}
	// string형으로 오름차순 정렬 되있는지 체크
	fmt.Println(sort.StringsAreSorted(sliceString))
	// string으로 오름차순 정렬
	sort.Strings(sliceString)
	//문자열 배열 출력
	fmt.Println(sliceString)
}
