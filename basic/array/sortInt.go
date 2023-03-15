package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceInt := []int{1, 2, 5, 6, 3, 4}

	//오름차순으로 int형 정렬이 되어있는지 체크함
	fmt.Println(sort.IntsAreSorted(sliceInt))
	//Ints형 정렬 수행
	sort.Ints(sliceInt)
	//sliceInt 수행
	fmt.Println(sliceInt)
}
