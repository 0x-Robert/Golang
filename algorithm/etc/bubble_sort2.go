package main

import (
	"fmt"       //프린트 패키지
	"math/rand" //난수 패키지
	"time"      //시간 패키지
)

func main(){
	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubblesort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano()) //유닉스시간의 나노초 
	fmt.Println(time.Now().UnixNano())

	// rand.Intn(100) > 1부터 100까지 정수 형태의 난수가 출력됨 
	fmt.Println("rand.Intn(100)",rand.Intn(100))
	for i := 0; i < size; i++{
		//1부터 999까지의 난수 - 1부터 999까지의 난수 
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}

	fmt.Println("slice",slice)
	return slice 
}

func bubblesort(items []int){
	var (
		n = len(items) //배열의 길이 
		sorted = false //sorted에 bool값 선언 
	)
	for !sorted{

		swapped := false 
		for i := 0; i < n - 1; i++ {
			//앞의 원소가 뒤 원소보다 크면 swapped  
			if items[i] > items[i+1]{
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true 
			}
		}
		
		//마지막 출력을 위해 선정한 것으로 보임 
		if !swapped{
			sorted = true 
		}
		//n에는 n - 1 값 대입
		n = n - 1 
		
	}
}