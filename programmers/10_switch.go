package main

import (
	"fmt"
)

func no_fallthrough(score int){
	var grade string 
	switch {
	case score > 90: 
		grade = "A"
		
	case score > 70:
		grade = "B"
		
	case score > 50:
		grade = "C"
	default:
		 grade="F"
	}

	fmt.Println("fallthrough를 쓰지않으면", grade,"란다??" )
}

func yes_fallthrough(score int){
	var grade string 
	switch{
	case score > 90:
		grade = "A"
		fallthrough 
	case score > 70:
		grade = "B"
		fallthrough 
	case score > 50:
		fallthrough 
	default:
		grade = "F"
	}
	fmt.Println("fallthrough를 쓰면",grade,"란다")
}

func main(){
	fmt.Println("교수님 제성적은?")

	score := 100 
	yes_fallthrough(score)
	no_fallthrough(score)
}