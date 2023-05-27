// . append 함수는 입력 인자가 가변 길이로 여러 개의 값을 인자로 전달하여 여러 개의 값을 한꺼번에 저장할 수도 있어요. 다음은 이를 확인하는 간략한 코드다.

package main

import "fmt"


func main(){
	var datas []int
	fmt.Println(datas)

	datas = append(datas, 9, 8, 7, 6 ,5, 4 )
	fmt.Println(datas)
}