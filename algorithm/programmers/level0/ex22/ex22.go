// package main

// import "fmt"

// func main() {
// 	var s1 string
// 	fmt.Scan(&s1)
// 	for _, v := range s1 {
// 		fmt.Println(string(v))
// 	}
// }

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}
}
