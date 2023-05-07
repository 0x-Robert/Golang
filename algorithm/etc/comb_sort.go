package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//난수 배열 생성 함수 호출 
	slice := generateSlice(20)
	//정렬 전 
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//빗질 정렬 호출  
	combsort(slice)
	//정렬 후
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

//난수 배열 생성 
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 
func combsort(items []int) {
	var (
		n = len(items)	//배열의 길이 
		gap = len(items) //배열의 길이
		shrink = 1.3 //gap을 축소해주는 계수 설정 
		swapped = true //swap 판단 값
	)
	
	for swapped {
		swapped = false
			//다음 정렬을 위한 gap 밸류 업데이트 
        	gap = int(float64(gap) / shrink)
			//갭이 1 미만이면 
        	if gap < 1 {
				//갭은 1 
			gap = 1
		}
		//반복 순회 특이점은 i에 gap을 더함  
		for i := 0; i+gap < n; i++ {
			//아이템 배열에서 i 인덱스가 i + gap 인덱스보다 크면 인덱스 위치 교환
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				//스왑=true 
				swapped = true
			}	
		}
	}
}