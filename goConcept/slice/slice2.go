// ch18/ex18.1/ex18.1.go
package main

import "fmt"

func main() {
	var slice []int
	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)

	}
	// slice[1] = 10
	// fmt.Println(slice)

	var slice2 = []int{1, 2, 3}
	for i := 0; i < len(slice2); i++ {
		slice2[i] += 10
	}
	for i, v := range slice2 {
		slice2[i] = v * 2
	}

	fmt.Println(slice2)
}
