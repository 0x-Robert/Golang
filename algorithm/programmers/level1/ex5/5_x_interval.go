package level1

import "fmt"

func solution5(x int, n int) []int64 {

	arr := make([]int64, 0)
	for i := 1; i <= n; i++ {
		arr = append(arr, int64(x*i))
	}
	fmt.Println("arr", arr)
	return arr

}

func interval() {
	x := 2
	n := 5
	solution5(x, n)
}
