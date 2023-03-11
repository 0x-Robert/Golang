package main

import "fmt"

func ChangeArray(arr [5]int) {
	arr[3] = 3000
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	ChangeArray(a)
	//arr과  a는 다른 메모리주소를 가진 배열이라서 arr[3] 값을 바꿔도 a[3]의 값은 바뀌지 않음
	fmt.Println(a[3])
}
