package main

import "fmt"


func main(){
	arr := [10]int{9,1,2,8,6,7,5,3,0,4}
	fmt.Println("init",arr)
	fmt.Println("")

	tmp := 0

	for i := 0; i < len(arr); i++{
		for j := 0; j < len(arr) -1 -i; j++{
			fmt.Println("i",i,"j",j)
			if arr[j] > arr[j+1]{
				tmp = arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp 
			}
		}
	}
	fmt.Println("sorted array is ", arr)
}