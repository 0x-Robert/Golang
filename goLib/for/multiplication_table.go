package main

import "fmt"

func main() {
	k := 0
	for i := 1; i < 10; i++ {

		for j := 2 + k; j < 6+k; j++ {
			if i%4 == 0 {
				continue
			} else {
				fmt.Printf("%d * %d = %d\t", j, i, i*j)
			}
		}
		if i%4 == 0 {
			continue
		} else {
			fmt.Println()
		}

		if i == 9 && k == 0 {
			i = 0
			k += 4
			fmt.Println()
			continue
		}
	}
}
