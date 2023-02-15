package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//난수 배열 생성 
	slice := generateSlice(20)
	//정렬 전 
	fmt.Println("--- Unsorted ---")
	//slice 출력 
	fmt.Println(slice)
	//정렬 후 
	fmt.Println("--- Sorted ---")
	//병합정렬 결과 
	fmt.Println(mergeSort(slice))
}

func generateSlice(size int) []int {
	//slice 초기화 
	slice := make([]int, size)
	//시간 동기화 
	rand.Seed(time.Now().UnixNano())
	//난수생성으로 배열 저장 
	for i := 0; i < size; i++ {
		slice[i] = rand.Int() % 1000 - 500
	}
	return slice
}

//병합정렬 
func mergeSort(items []int) []int {
	//2미만은 그대로 리턴 
	if len(items) < 2 {
		return items
	}

	//middle은 중간 값, middle을 기준으로 left, right 값 설정 
	middle := len(items) / 2
	//재귀 호출 
	left := mergeSort(items[:middle])
	right := mergeSort(items[middle:])
	fmt.Println("left",left, "right",right)
	//merge 반환 
	return merge(left, right)
}

//병합 함수
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	//left와 right 배열 크기가 0일 때 반복문 종료 
	for len(left) > 0 && len(right) > 0 {
		//left[0]가 작으면 result에 left[0]대입
		//left[1:]이후부터 모든 원소를 left 대입
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		//right[0]이 더 클 경우 
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	//result에 left, right 삽입 
	result = append(result, left...)
	result = append(result, right...)

	//result 반환
	return result
}