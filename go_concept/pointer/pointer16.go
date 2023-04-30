package main

import "fmt"

func add(p1, p2, p3 *int) {
	*p3 = *p1 + *p2
}

func main() {
	a := 3
	b := 5
	c := 0

	add(&a, &b, &c)
	fmt.Println(c)
}
