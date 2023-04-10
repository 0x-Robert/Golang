package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	b = array //1,2,3,4,5

	var c []int

	c = array[:] //1,2,3,500,5

	b[0] = 1000 //1000, 2,3,4,5
	c[3] = 500  //1,2,3,500,5

	fmt.Println("array:", array)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
