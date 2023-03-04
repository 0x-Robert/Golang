package main

import (
	"fmt"
	"sort"
)

// Go에서 배열과 슬라이스는  다릅니다.

// 배열: 고정 크기의 메모리 블록입니다. 정적으로 할당되며 크기는 초기화할 때 결정됩니다. 배열의 크기를 바꿀 수 없습니다.
// 슬라이스: 배열의 일부를 참조하는 동적 크기의 구조입니다. 슬라이스는 make() 함수로 생성되거나 슬라이스 리터럴로 생성됩니다. 슬라이스는 길이와 용량이 있으며 용량은 내부 배열의 길이입니다. 슬라이스의 크기는 필요에 따라 동적으로 늘어날 수 있습니다.
// 따라서 슬라이스는 동적으로 크기를 조정할 수 있으므로 좀 더 유연하고 배열은 고정 크기를 가지므로 메모리 사용량이 더 효율적입니다. 또한 배열은 값 타입이고 슬라이스는 참조 타입입니다.

func medianOfMedians(sliceList []int, k, r int) int {

	//배열 크기 지정
	num := len(sliceList)
	//num이 10 미만일 때 조건문 실행
	if num < 10 {
		//Int형으로 정렬함
		sort.Ints(sliceList)
		//10미만이면 sliceList[k-1]을 리턴
		return sliceList[k-1]
	}

	//med 값 설정
	med := (num + r - 1) / r
	fmt.Println("med1", med)

	//슬라이스 생성
	medians := make([]int, med)
	fmt.Println("med2", med)
	fmt.Println("medians", medians)

	for i := 0; i < med; i++ {

		v := (i * r) + r

		var arr []int

		if v >= num {
			//슬라이스 생성
			arr = make([]int, len(sliceList[(i*r):]))
			//arr 변수로 sliceList[(i*r)]을 복사한다.
			copy(arr, sliceList[(i*r):])
		} else {
			//슬라이스 생성
			arr = make([]int, r)
			//arr 변수로 sliceList[(i*r):v]을 복사한다.
			copy(arr, sliceList[(i*r):v])
		}
		//Int형으로 오름차순 정렬
		sort.Ints(arr)
		//arr길이의 1/2 인덱스의 원소를 medians[i]로 지정
		medians[i] = arr[len(arr)/2]
	}

	pivot := medianOfMedians(medians, (len(medians)+1)/2, r)
	fmt.Println("pivot", pivot)
	var leftSide, rightSide []int

	for i := range sliceList {
		if sliceList[i] < pivot {
			// append 함수는 슬라이스에 값을 추가할 때 사용하는 함수입니다.
			// leftSide 변수는 빈 슬라이스로 선언되었기 때문에 append 함수를 사용하여 sliceList[i] 값을 추가하면 leftSide 슬라이스에 그 값이 저장됩니다.
			leftSide = append(leftSide, sliceList[i])
		} else if sliceList[i] > pivot {
			rightSide = append(rightSide, sliceList[i])
		}
	}
	fmt.Println("leftSide", leftSide)
	fmt.Println("rightSide", rightSide)

	switch {

	//k == len(leftSide) + 1 길이와 같을 때
	case k == (len(leftSide) + 1):
		return pivot

		//k == len(leftSide)길이 이하일 대
	case k <= len(leftSide):
		return medianOfMedians(leftSide, k, r)
		//디폴트로 설정
	default:
		return medianOfMedians(rightSide, k-len(leftSide)-1, r)
	}
}

// main 함수에서는 medianOfMedians 함수를 사용하여 정수 슬라이스에서 k번째로 작은 요소를 찾습니다.
func main() {

	intSlice := []int{5, 9, 77, 62, 71, 11, 22, 46, 36, 18, 19, 33, 75, 17, 39, 41, 73, 50, 217, 79, 120}
	fmt.Println("intSlice 1", intSlice)
	sort.Ints(intSlice)
	fmt.Println("intSlice 2", intSlice)

	for _, j := range []int{5, 10, 15, 20} {

		i := medianOfMedians(intSlice, j, 5)
		//j번째 엘리먼트는 i
		fmt.Println(j, "smallest element = ", i)

		v := intSlice[j-1]
		fmt.Println("arr[", j-1, "]=", v)

		if i != v {
			fmt.Println("oops Algorithm is wrong")
		}

	}
}
