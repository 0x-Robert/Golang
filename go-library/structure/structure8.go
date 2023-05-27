package main

import "fmt"

type Product struct {
	Name        string
	Price       int
	ReviewScore float64
}

func main() {

	product := Product{"yongari", 10, 3.3}
	fmt.Println("product", product)
}
