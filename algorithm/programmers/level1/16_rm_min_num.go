package main

import "fmt"

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func solution(arr []int) []int {
	if len(arr) == 1 {
		arr2 := []int{}
		arr2 = append(arr2, -1)
		return arr2
	}
	min := arr[0]
	var finNum int
	for i := 0; i < len(arr)-1; i++ {
		if min > arr[i+1] {
			min = arr[i+1]
			finNum = i + 1
		}
	}
	fmt.Println("min", min)
	fmt.Println("finNum", finNum)
	arr = remove(arr, finNum)
	fmt.Println("arr", arr)
	return arr
}

func solution2(arr []int) []int {
	idx := 0
	min := arr[0]

	for i, val := range arr {
		if min > val {
			min = val
			idx = i
		}
	}

	if len(arr) > 1 {
		arr = append(arr[:idx], arr[idx+1:]...)
	} else {
		arr[0] = -1
	}

	return arr
}
func main() {
	//arr := []int{1, 2, 3, 4}
	arr := []int{4, 3, 2, 1}
	//arr := []int{10}
	solution(arr)
}
