package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	for i, v := range a {
		a[i] = v * 2
	}
	fmt.Println("a", a)
	fmt.Println(a[2])
}
