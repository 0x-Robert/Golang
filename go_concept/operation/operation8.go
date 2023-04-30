// ch6/ex6.9/ex6.9.go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Float).SetString("0.1") // ❶
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b) // ❷ 
	fmt.Println(a, b, c, d) 
	fmt.Println(c.Cmp(d)) // ❸ 반환값 -1은 x가 작은 경우, 반환값 1은 x가 큰 경우, 0은 두 값이 같은 경우
}
