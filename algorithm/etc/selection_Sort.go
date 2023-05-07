package main

import (
	"fmt"       //print함수를 위한 패키지
	"math/rand" //난수 생성을 위한 패키지
	"time"      //시간활용을 위한 패키지
)



func main() {
	//난수 배열 생성이후 slice 변수에 저장 
	slice := generateSlice(20)
	//정렬 전 배열 
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//선택정렬 함수 호출 
	selectionsort(slice)
	//정렬 된 배열 
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	//make를 통해 슬라이스 초기화 
	slice := make([]int, size, size)

	// 현재 시간으로 Seed 값 설정, time 패키지를 이용한 것이고 현재 시간을 설정하는 함수로 보임 
	rand.Seed(time.Now().UnixNano())
	//size만큼 반복문 돌고 slice 배열에 난수배열 생성 
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
//선택정렬 함수 
func selectionsort(items []int) {
	//n은 입력받은 배열의 크기 
    var n = len(items)
	//n의 크기만큼 반복 순회 
    for i := 0; i < n; i++ {
		//minIdx 값 설정 
        var minIdx = i
		//j는 i부터 시작하고 n 미만일 때 j에 +1씩 진행함 
        for j := i; j < n; j++ {
			//items[j]가 items[minIdx] 미만일 때 minIdx에 j 값 설정 
            if items[j] < items[minIdx] {
                minIdx = j
            }
			//fmt.Println("minIdx",minIdx)
			fmt.Println("items",items)
        }
		//i, minIdx 위치 교환 
        items[i], items[minIdx] = items[minIdx], items[i]
    }
}