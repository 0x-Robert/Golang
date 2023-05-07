package main

import (
	"fmt"
	"math/cmplx"
)

var(
	i int 
	f float64 
	MaxInt uint64 = 1<<64 -1 
	z complex128 = cmplx.Sqrt(-5 + 12i)
)


func main(){
	const format = "%T(%v)\n"
	fmt.Println(format, MaxInt, MaxInt)
	fmt.Println(format, z, z)

	f = float64(i)
}