package main

import "fmt"

func Solution(arr []int) []int {

	stk := []int{}

	for i := 0; i < len(arr); i++ {
		if len(stk) == 0 {
			stk = append(stk, arr[i])
		} else if len(stk) >= 1 && stk[len(stk)-1] == arr[i] {
			stk = stk[:len(stk)-1]
		} else if len(stk) >= 1 && stk[len(stk)-1] != arr[i] {
			stk = append(stk, arr[i])
		}
	}
	if len(stk) == 0 {
		stk = append(stk, -1)
	}
	fmt.Println(stk)
	return stk
}

