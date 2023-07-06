package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//난수로 배열생성
	slice := generateSlice(20)
	//정렬 전 
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//삽입정렬함수 호출 
	insertionsort(slice)
	//정렬 후 
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {
	//slice 변수 초기화 
	slice := make([]int, size, size)
	//현재시간 맞추기 
	rand.Seed(time.Now().UnixNano())

	//랜덤 난수 생성해서 배열에 저장 
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	//배열 반환 
	return slice
}
 
//삽입정렬 
func insertionsort(items []int) {
	//n은 배열의 크기 
    var n = len(items)
	//1부터 n까지 반복 순회 
    for i := 1; i < n; i++ {
		//j에 i 대입 
        j := i
		//j가 0보다 클 때 
        for j > 0 {
			fmt.Println("before items",items)
			//만약 j기준으로 배열의 앞 요소가 뒤 요소보다 크면
            if items[j-1] > items[j] {
				//j-1 인덱스와 j인덱스 위치 교환 
                items[j-1], items[j] = items[j], items[j-1]
            }
			fmt.Println("after items",items)
			//j에 j-1 값 대입 
            j = j - 1
        }
    }
}