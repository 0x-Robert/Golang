// ch15/ex15.12/ex15.12.go
package main

import "fmt"

func main() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	//B의 UTF8값이 66번, a의 UTF8값이 97번 따라서 B보다 a가 더 큰 값 
	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)   // ❶
	//BB까지는 같지만 B는 UTF값이 66번이고 A는 65번이라서 BBB가 BBAD보다 크다. 
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)   // ❷
	//B는 66번, Z는 90번이라서 BBB <= ZZZ는 true 이다. 
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4) // ❸
}
