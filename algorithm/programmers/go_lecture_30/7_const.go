package main

import "fmt"

const Pi1 float32 = 3.14 
const Pi2 = 3.14

const (
	Big_const =  1 << 100 
	Small_const = Big_const >> 99 
)

//아래와 같은 큰 수는 const로 표현해야함 var로는 표현할 수 없는 범위 
// var Big_var =  1 << 100 


func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64{
	return x * 0.1 
}

func main(){
	fmt.Println("needInt small ",needInt(Small_const))
	fmt.Println("needFloat small", needFloat(Small_const))
	fmt.Println("needFloat big", needFloat(Big_const))
}


