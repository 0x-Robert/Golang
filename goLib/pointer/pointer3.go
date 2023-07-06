package main

import "fmt"


type Data struct {
	value int 
	data [200]int //200개 원소를 가진 int형 배열 
}

//포인터 받자 , 포인터 커봐야 8바이트 용량 절약~~
func ChangeData( d *Data){
	// 아니 그럼 *d.value라고 해야하는거 아니냐? 굳이 역참조 안해줘도 go에서는 동작함
	//....
	d.value = 999;
	d.data[100] = 999; 
}

func main(){
	var data Data 

	//함수 인자로 전달되면서 함수 스코프 내의 지역변수로 값 복사 
	//struct 크기는 201 * 8 바이트만큼 할당 
	//함수호출할때마다 복사, 
	ChangeData(&data)
	fmt.Println(data.value)
	fmt.Println(data.data[100])

}