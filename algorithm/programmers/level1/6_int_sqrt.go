package main 

import(
	"fmt"
)

func solution(n int64) int64 {
    
 	var a int64 
	for a = 1; a <= n; a++ {
		if (n == a * a ){
			fmt.Println("a+1",a+1)
			return (a+1) * (a+1) 
		}
	}
	return -1 
}


func main (){

	n:= int64(3)
	// n:= int64(121)
	solution(n)
}