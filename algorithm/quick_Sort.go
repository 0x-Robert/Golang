package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    //배열 무작위 생성 
	slice := generateSlice(20)
    //정렬 전 배열 
	fmt.Println("\n--- Unsorted --- \n\n", slice) 
    //퀵정렬 함수 호출 
	quicksort(slice)
    //정렬 후 배열 
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
    //slice 초기화 
	slice := make([]int, size, size)
    //유닉스시간 나노초 
	rand.Seed(time.Now().UnixNano())
	
    //1부터 999 난수 - 1부터 999난수를 원소로하는 배열 생성 
    for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 
func quicksort(a []int) []int {

    //배열 길이가 2 미만이면 a를 리턴 
    if len(a) < 2 {
        return a
    }
     
    //left는 0, right는 배열 길이 - 1 
    left, right := 0, len(a)-1
     
    //pivot은 랜덤난수를 a 길이로 나눈 나머지 
    pivot := rand.Int() % len(a)
     
    //right와 pivot 위치 바꾸기 
    a[pivot], a[right] = a[right], a[pivot]
     
    //a의 길이만큼 반복 
    for i, _ := range a {
        if a[i] < a[right] {
            //i요소가 right보다 작으면 
            //a left와 a i의 위치를 교환 
            a[left], a[i] = a[i], a[left]
            //left ++ 
            left++
        }
    }
     
    //left와 right 위치 교환 
    a[left], a[right] = a[right], a[left]
     
    //재귀호출 
    quicksort(a[:left])
    quicksort(a[left+1:])
     
    return a
}