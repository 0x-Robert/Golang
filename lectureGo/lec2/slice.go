package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a1 []int
	var a2 [5]int

	fmt.Println(reflect.ValueOf(a1).Kind())
	fmt.Println(reflect.TypeOf(a1))

	fmt.Println(reflect.ValueOf(a2).Kind())
	fmt.Println(reflect.TypeOf(a2))

	var intSlice = []int{0, 1, 2, 3}
	fmt.Printf("%#v\n", intSlice)
	fmt.Printf("%v\n", intSlice)
	fmt.Printf(" :2 %v\n", intSlice[:2])
	fmt.Printf("1:3 %v\n", intSlice[1:3])
	fmt.Printf("2: %v\n", intSlice[2:])
	fmt.Printf("0:3 %v\n", intSlice[0:3])

}
