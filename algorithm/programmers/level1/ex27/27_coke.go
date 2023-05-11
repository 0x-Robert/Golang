package level1

import "fmt"

func solution27(a int, b int, n int) int {
	var result int
	//20>2
	for n > a {
		if n%a == 0 {
			result += n / a
			n = n / a
			fmt.Println(result)
			fmt.Println(n)
		} else {
			result += n/a + 1
			n = n/a + 1
			fmt.Println(result)
			fmt.Println(n)
		}
	}

	fmt.Println("result", result)
	return 0
}

func coke() {
	a := 2
	b := 1
	n := 20
	solution27(a, b, n)
}
