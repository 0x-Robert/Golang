package main

import (
	"sort"
)

func solution(array []int, commands [][]int) []int {

	answer := []int{}
	array2 := make([]int, len(array))
	temp := []int{}

	for _, v := range commands {
		if v[0] != v[1] {
			copy(array2[:], array[:])
			temp = array2[v[0]-1 : v[1]]
			sort.Sort(sort.IntSlice(temp))
			answer = append(answer, temp[v[2]-1])

		} else if v[0] == v[1] {
			answer = append(answer, array[v[0]-1])
		}
	}

	return answer
}

func solution0(array []int, commands [][]int) []int {
	var ret []int
	for _, cmd := range commands {
		slice := append([]int{}, array[cmd[0]-1:cmd[1]]...) //이게 핵심!

		sort.Ints(slice)
		ret = append(ret, slice[cmd[2]-1])
	}
	return ret
}

func main() {
	//arr := []int{1, 5, 2, 6, 3, 7, 4}
	arr := []int{1, 5, 2, 60, 3, 7, 4}
	commands := [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}
	solution(arr, commands)
}
