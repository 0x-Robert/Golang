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
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	//쉘정렬 
	shellsort(slice)
	//정렬 후
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Generates a slice of size, size filled with random numbers
// 랜덤난수로 배열 생성
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 

//쉘정렬 
func shellsort(items []int) {
	//n은 배열의 크기 설정, gaps 배열 설정 
    var (
        n = len(items)
        gaps = []int{1}
        k = 1
     
    )
     
    for {
        //k를 이용해서 간격 배열 계산 
        gap := element(2, k) + 1
        if gap > n-1 {
            break
        }
		
        gaps = append([]int{gap}, gaps...)
        k++
    }
     
    for _, gap := range gaps {
        for i := gap; i < n; i += gap {
            j := i
			//for문 j가 양수일 때 
            for j > 0 {
                if items[j-gap] > items[j] {
					//j-gap 인덱스가 j보다 크면 인덱스 위치 변경 
                    items[j-gap], items[j] = items[j], items[j-gap]
                }
				//j 값에 j - gap 대입 
                j = j - gap
            }
        }
    }
}
 
func element(a, b int) int {
    e := 1
    for b > 0 {
        if b&1 != 0 {
            e *= a
        }
        b >>= 1
        a *= a
    }
    return e
}
