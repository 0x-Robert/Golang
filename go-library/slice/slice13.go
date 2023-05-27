package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s) //정렬함수
	fmt.Println(s)

}
