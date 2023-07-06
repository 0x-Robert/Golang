package main

import "fmt"

func solution(price int) int {

	if price >= 100000 && price < 300000 {
		price = price * 95 / 100
	} else if price >= 300000 && price < 500000 {
		price = price * 90 / 100
	} else if price >= 500000 {
		price = price * 80 / 100
	}
	fmt.Println(price)

	return price
}

func main() {
	//price := 150000
	price := 580000
	//price := 150000
	solution(price)
}
