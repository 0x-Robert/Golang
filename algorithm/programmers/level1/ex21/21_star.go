package level1

import "fmt"

func Star() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Println()
	}
}
