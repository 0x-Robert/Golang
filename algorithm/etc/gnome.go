// 난쟁이 정렬
// stupid sort
package main

import "fmt"
func main (){
	arr := [10]int{9,1,2,8,6,7,5,3,0,4}
	fmt.Println("init", arr)
	fmt.Println("")


	i:=1
	tmp := 0
	for ; i < len(arr) ; {
		if arr[i] >= arr[i-1]{
			i++
		}else{
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[ i - 1] = tmp 
		
			if i > 1 {
				i--
			}
		}
	}

}